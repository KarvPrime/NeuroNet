package Core

import (
	"NeuroNet/Core/Brain"
	"NeuroNet/Core/Brain/Connection"
	"NeuroNet/Core/FileHandler"
	"NeuroNet/Core/Logging"
	"NeuroNet/Core/Persistence"
	"NeuroNet/Core/WorkerPool"
	"fmt"
	"math/rand"
	"path"
	"runtime"
	"strconv"
	"time"
)

// Struct for the core parameters
type Core struct {
	workOrder     [][2]string
	networkLayout [][2]string

	preProcessing string

	root          string
	trainFile     string
	testFile      string
	parallel      int
	workerBatch   int
	weightRange   [2]float64
	targetColumns [2]int

	learningRate float64
	lambda       float64

	data  [][]float64
	lines int

	dataChannel   chan [][]float64
	outputChannel chan int

	connections []*Connection.Connection
	networks    []*Brain.Network
	workerPool  *WorkerPool.WorkerPool
	state       []*Connection.Connection

	persistence  *Persistence.Persistence
	networkBatch *FileHandler.FileHandler
	dataReader   *FileHandler.FileHandler
	result       *FileHandler.FileHandler
}

// Start the neural network
func (core *Core) Start(batchFile string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		Logging.Fatal("Could not get path to logfile.")
	}
	core.root = path.Join(path.Dir(filename), "../")

	Logging.Initialize(core.root)

	Logging.Info("Program started.")
	Logging.Info("Go Max Procs " + strconv.Itoa(runtime.GOMAXPROCS(runtime.NumCPU())) + " / " + strconv.Itoa(runtime.NumCPU()) + " CPUs")
	fmt.Println()

	core.loadBatch(core.pathFinder(batchFile))

	core.dataReader = new(FileHandler.FileHandler)
	core.result = new(FileHandler.FileHandler)
	core.persistence = new(Persistence.Persistence)
	core.networkBatch = new(FileHandler.FileHandler)
	core.workerPool = new(WorkerPool.WorkerPool)

	// Standard values
	core.weightRange = [2]float64{-1, 1}
	core.targetColumns = [2]int{0, 1}
	core.parallel = 1
	core.workerBatch = 100
	core.learningRate = 1
	core.lambda = 1
	core.preProcessing = "None"

	core.runBatch()
}

// Prepare the neural network
func (core *Core) runMode(mode string) {
	core.result.AddValue("Mode", mode)
	core.result.AddValue("CPUs", strconv.Itoa(runtime.NumCPU()))
	core.result.AddValue("Parallel", strconv.Itoa(core.parallel))
	core.result.AddValue("WorkerBatch", strconv.Itoa(core.workerBatch))
	core.result.AddValue("LearningRate", strconv.FormatFloat(core.learningRate, 'f', 10, 64))

	fmt.Println()

	// Set up channels for the workers
	core.dataChannel = make(chan [][]float64, core.parallel)
	core.outputChannel = make(chan int, core.parallel)

	// Create the workers
	core.workerPool.Initialize(mode, core.parallel, core.targetColumns, core.learningRate, core.lambda, core.networkLayout, core.preProcessing, core.dataChannel, core.outputChannel)
	networks, connections, setupTime, err := core.workerPool.CreatePool()
	if err != nil {
		Logging.Fatal("Error while setting up worker pool networks: " + err.Error())
	}

	fmt.Printf("%d worker(s) ready after %f seconds. ", core.parallel, setupTime)
	Logging.Info("Number of Goroutines: " + strconv.Itoa(runtime.NumGoroutine()))
	fmt.Println()

	// Create the networks
	core.networks = networks
	core.connections = connections
	core.state = make([]*Connection.Connection, len(core.connections))
	for index, element := range core.connections {
		core.state[index] = new(Connection.Connection)
		core.state[index].Initialize(element.GetSize()[0], element.GetSize()[1])
	}

	core.result.AddValue("WorkerSetupTime", strconv.FormatFloat(setupTime, 'f', 10, 64))

	var filePath string
	if mode == "train" {
		filePath = core.trainFile
	} else {
		filePath = core.testFile
	}

	// Read the data
	core.dataReader.Initialize(filePath, true, false, false, false)
	core.data = core.dataReader.ReadData()
	core.lines = len(core.data)

	// Shuffle data
	core.shuffleBatch(core.data)

	totalLineNumber, totalTimeElapsed := core.loadPersistence(mode)
	epochLineNumber, epochTimeElapsed, correct, incorrect, errorRate := core.run(mode, totalLineNumber, totalTimeElapsed)

	fmt.Println()

	// Show progress in console
	fmt.Printf("Time: \033[01;33m%f\033[00;00m seconds. Line: \033[00;33m%d\033[00;00m/%d\n", epochTimeElapsed, epochLineNumber, core.lines)

	core.result.AddValue("EpochElapsedTime", strconv.FormatFloat(epochTimeElapsed, 'f', 10, 64))
	core.result.AddValue("EpochLines", strconv.Itoa(epochLineNumber))

	core.result.AddValue("TotalElapsedTime", strconv.FormatFloat(totalTimeElapsed, 'f', 10, 64))
	core.result.AddValue("TotalLines", strconv.Itoa(totalLineNumber))

	// Show test results in console
	if mode == "test" {
		correctPercent := float64(correct) / float64(core.lines) * 100
		incorrectPercent := float64(incorrect) / float64(core.lines) * 100
		fmt.Printf("\033[00;32mCorrect %d (%f%%)\033[00;00m : \033[00;91m(%f%%) %d Incorrect\033[00;00m\n", correct, correctPercent, incorrectPercent, incorrect)

		fmt.Printf("Error Rate: %f%%\n", errorRate)

		core.result.AddValue("Correct", strconv.Itoa(correct))
		core.result.AddValue("Incorrect", strconv.Itoa(incorrect))
		core.result.AddValue("CorrectPercent", strconv.FormatFloat(correctPercent, 'f', 10, 64))
		core.result.AddValue("IncorrectPercent", strconv.FormatFloat(incorrectPercent, 'f', 10, 64))
		core.result.AddValue("ErrorRate", strconv.FormatFloat(errorRate, 'f', 10, 64))
	} else {
		totalLineNumber += epochLineNumber
		totalTimeElapsed += epochTimeElapsed

		// Remember after training
		core.persistence.Write(totalLineNumber, totalTimeElapsed, core.connections)

		fmt.Printf("Total Time: \033[01;33m%f\033[00;00m seconds. Total Line: \033[00;33m%d\033[00;00m/%d\n", totalTimeElapsed, totalLineNumber, core.lines)

		core.result.AddValue("Correct", "")
		core.result.AddValue("Incorrect", "")
		core.result.AddValue("CorrectPercent", "")
		core.result.AddValue("IncorrectPercent", "")
		core.result.AddValue("ErrorRate", "")
	}

	fmt.Println()

	close(core.dataChannel)
	close(core.outputChannel)
}

// Load batch file
func (core *Core) loadBatch(batchFile string) {
	batch := new(FileHandler.FileHandler)
	batch.Initialize(batchFile, true, false, false, false)
	core.workOrder = batch.ReadBatch(": ")
}

// Run through the batch
func (core *Core) runBatch() {
	for _, element := range core.workOrder {
		fmt.Printf("\033[00;36m%s\033[00;00m: %s\n", element[0], element[1])
		switch element[0] {
		case "NetworkFile":
			core.networkBatch.Initialize(core.pathFinder(element[1]), true, false, false, false)
			core.networkLayout = core.networkBatch.ReadBatch(", ")
			break
		case "TrainFile":
			core.trainFile = core.pathFinder(element[1])
			break
		case "TestFile":
			core.testFile = core.pathFinder(element[1])
			break
		case "PersistenceFile":
			core.persistence.Initialize(core.pathFinder(element[1]))
			break
		case "ResultFile":
			core.result.Initialize(core.pathFinder(element[1]), true, true, true, true)
			break
		case "Parallel":
			value, err := strconv.ParseInt(element[1], 10, 64)
			if err != nil {
				Logging.Panic("Batch Parallel is NaN: " + err.Error())
			}
			core.parallel = int(value)
			break
		case "WorkerBatch":
			value, err := strconv.ParseInt(element[1], 10, 64)
			if err != nil {
				Logging.Panic("Batch WorkerBatch is NaN: " + err.Error())
			}
			core.workerBatch = int(value)
			break
		case "LearningRate":
			value, err := strconv.ParseFloat(element[1], 64)
			if err != nil {
				Logging.Panic("Batch LearningRate NaN: " + err.Error())
			}
			core.learningRate = value
			break
		case "Lambda":
			value, err := strconv.ParseFloat(element[1], 64)
			if err != nil {
				Logging.Panic("Batch Lambda NaN: " + err.Error())
			}
			core.lambda = value
			break
		case "MinWeight":
			value, err := strconv.ParseFloat(element[1], 64)
			if err != nil {
				Logging.Panic("Batch MinWeight is NaN: " + err.Error())
			}
			core.weightRange[0] = value
			break
		case "MaxWeight":
			value, err := strconv.ParseFloat(element[1], 64)
			if err != nil {
				Logging.Panic("Batch MaxWeight is NaN: " + err.Error())
			}
			core.weightRange[1] = value
			break
		case "TargetColumnStart":
			value, err := strconv.ParseInt(element[1], 10, 64)
			if err != nil {
				Logging.Panic("Batch TargetColumnStart is NaN: " + err.Error())
			}
			core.targetColumns[0] = int(value)
			break
		case "TargetColumnEnd":
			value, err := strconv.ParseInt(element[1], 10, 64)
			if err != nil {
				Logging.Panic("Batch TargetColumnEnd is NaN: " + err.Error())
			}
			core.targetColumns[1] = int(value)
			break
		case "PreProcessing":
			core.preProcessing = element[1]
			break
		case "Train":
			value, err := strconv.ParseInt(element[1], 10, 64)
			if err != nil {
				Logging.Panic("Batch train runs NaN: " + err.Error())
			}
			for i := 0; i < int(value); i++ {
				core.runMode("train")
				core.result.SaveLine()
			}
			break
		case "Test":
			value, err := strconv.ParseInt(element[1], 10, 64)
			if err != nil {
				Logging.Panic("Batch test runs NaN: " + err.Error())
			}
			for i := 0; i < int(value); i++ {
				core.runMode("test")
				core.result.SaveLine()
			}
			break
		default:
			break
		}
	}

	Logging.Info("Finished work order.")
}

// Load persistence file
func (core *Core) loadPersistence(mode string) (int, float64) {
	lineNumber, timeElapsed, connections := core.persistence.Read()

	if lineNumber != 0 {
		for index, element := range connections {
			core.connections[index] = element
		}
		fmt.Printf("Training status: \033[00;33m%d\033[00;00m lines, elapsed time: \033[01;33m%f\033[00;00m seconds.\n", lineNumber, timeElapsed)
		if mode == "train" {
			Logging.Info("Training (\033[01;34mEpoch " + strconv.Itoa(lineNumber/core.lines+1) + "\033[00;00m) started.")
		} else {
			Logging.Info("New test started.")
		}
	} else if mode == "test" {
		Logging.Panic("No persistence for testing found.")
	} else {
		weightRange := core.weightRange[1] - core.weightRange[0]

		// Deterministic rand
		deterministicRand := rand.New(rand.NewSource(42))
		for index := range core.connections {
			for x := 0; x < core.connections[index].GetSize()[0]; x++ {
				for y := 0; y < core.connections[index].GetSize()[1]; y++ {
					core.connections[index].Weights[x][y] = deterministicRand.Float64()*-weightRange + core.weightRange[1]
				}
			}
		}

		Logging.Info("New training (\033[01;34mEpoch 1\033[00;00m) started.")
	}
	fmt.Println()

	return lineNumber, timeElapsed
}

// Find the path to the program
func (core *Core) pathFinder(pathname string) string {
	if pathname[0] == '.' {
		return path.Join(core.root, pathname)
	}
	return pathname
}

// Run the neural network
func (core *Core) run(mode string, totalLineNumber int, totalTimeElapsed float64) (int, float64, int, int, float64) {
	currentLine, activeWorkers, correct, incorrect, length := 0, 0, 0, 0, 0
	timeElapsed, errorRate := 0.0, 0.0

	running := true

	for running {
		startTime := time.Now()

		// Set the states for the networks
		for index := range core.networks {
			core.networks[index].SetState(core.connections)
		}

		// Give date to workers
		for ; activeWorkers < core.parallel; activeWorkers++ {
			if currentLine < core.lines {
				if currentLine+core.workerBatch < core.lines {
					length = core.workerBatch
				} else {
					length = core.lines - currentLine
				}
				core.dataChannel <- core.data[currentLine : currentLine+length]
				currentLine += length
			} else {
				running = false
				break
			}
		}

		// If there were lines to send wait for and process output
		if activeWorkers > 0 {
			if mode == "train" {
				for index, element := range core.connections {
					core.state[index].SetState(element.GetState())
				}
			}

			// Get results from workers
			for ; activeWorkers > 0; activeWorkers-- {
				id := <-core.outputChannel
				if mode == "train" {
					for index, element := range core.networks[id].GetState() {
						core.connections[index].AddWeighedDiff(core.state[index].GetState(), element.GetState(), activeWorkers)
					}
				}
			}

			timeElapsed += time.Since(startTime).Seconds()

			// FIXME: refresh line even when linebreak occurs in smaller terminals
			fmt.Printf("\033[1K\rTime: \033[01;33m%f\033[00;00m seconds. Line: \033[00;33m%d\033[00;00m/%d.", timeElapsed, currentLine, core.lines)

			// When testing get and show the results
			if mode == "test" {
				correct, incorrect = 0, 0

				for _, element := range core.networks {
					truthValues := element.GetTruthValues()
					correct += truthValues[true]
					incorrect += truthValues[false]
					errorRate += element.GetResetError()
				}

				fmt.Printf(" \033[00;32mCorrect %d\033[00;00m : \033[00;91m%d Incorrect\033[00;00m", correct, incorrect)
			}
		}
	}

	if mode == "test" {
		errorRate = 100 * errorRate / float64(core.lines)
	}

	fmt.Print("\r\033[K")
	Logging.Info("Process finished.")

	return currentLine, timeElapsed, correct, incorrect, errorRate
}

// Shuffle the data lines within the batch randomly to prevent learning repetitive patterns
func (core *Core) shuffleBatch(dataBatch [][]float64) {
	pseudoRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	pseudoRand.Shuffle(len(dataBatch), func(i, j int) {
		dataBatch[i], dataBatch[j] = dataBatch[j], dataBatch[i]
	})
}
