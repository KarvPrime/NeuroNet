package Core

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCore_Start(t *testing.T) {
	core := new(Core)
	core.Start("./Batchfile/TestBatch")
}

func TestCore_runMode(t *testing.T) {

}

func TestCore_loadBatch(t *testing.T) {

}

func TestCore_runBatch(t *testing.T) {

}

func TestCore_loadPersistence(t *testing.T) {

}

func TestCore_pathFinder(t *testing.T) {

}

func TestCore_run(t *testing.T) {

}

func TestCore_shuffleBatch(t *testing.T) {
	var dataBatch [][]float64
	dataBatch = append(dataBatch, []float64{1.0, 2.0, 3.0})
	dataBatch = append(dataBatch, []float64{1.1, 2.1, 3.1})
	dataBatch = append(dataBatch, []float64{1.2, 2.2, 3.2})
	dataBatch = append(dataBatch, []float64{1.3, 2.3, 3.3})
	dataBatch = append(dataBatch, []float64{1.4, 2.4, 3.4})
	dataBatch = append(dataBatch, []float64{1.5, 2.5, 3.5})

	origBatch := make([][]float64, len(dataBatch))
	for index := range dataBatch {
		origBatch[index] = make([]float64, len(dataBatch[0]))
		for i, elem := range dataBatch[index] {
			origBatch[index][i] = elem
		}
	}

	for index, element := range dataBatch {
		for i, elem := range element {
			if origBatch[index][i] != elem {
				t.Errorf("dataBatch and origBatch mismatch.")
			}
		}
	}

	core := new(Core)
	core.shuffleBatch(dataBatch)

	shuffleError := true
	for index, element := range dataBatch {
		for i, elem := range element {
			if origBatch[index][i] != elem {
				shuffleError = false
				break
			}
		}
	}

	if shuffleError {
		t.Errorf("Batch hasn't been shuffled.")
	}

	wrongShuffle := false
	for _, dElement := range dataBatch {
		for _, oElement := range origBatch {
			if dElement[0] == oElement[0] {
				for dIn, dEl := range dElement {
					if oElement[dIn] != dEl {
						wrongShuffle = true
						break
					}
				}
			}
		}
	}

	if wrongShuffle {
		t.Errorf("Batches have been shuffled inside.")
	}
}

func TestShuffle(t *testing.T) {
	test := []float64{0, 1, 2, 3, 5, 8, 12}
	test2 := []float64{0, 1, 2, 3, 5, 8, 12}

	pseudoRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(test), func(i, j int) {
		test[i], test[j] = test[j], test[i]
	})
	pseudoRand.Shuffle(len(test2), func(i, j int) {
		test[i], test[j] = test[j], test[i]
	})

	fmt.Println(test)
	fmt.Println(rand.Float64())
}

/*func TestCore_findBestParams(t *testing.T) {
	core := new(Core)

	weightRanges := []float64{0.1}
	learningRates := []float64{0.5, 0.25, 0.1}
	lambdas := []float64{1.0, 0.1, 0.01, 0.001}

	var correct float64
	var incorrect float64
	var correctMax float64
	var incorrectMax float64
	var accuracy float64
	var accuracyOld float64
	var maxWeightRange [2]float64
	var maxLearningRate float64
	var maxLambda float64

	for _, weightRangeElement := range weightRanges {
		for _, learningRateElement := range learningRates {
			for _, lambdaElement := range lambdas {
				core.Start("./Batchfile/TestBatch")

				core.weightRange = [2]float64{-weightRangeElement, weightRangeElement}
				core.learningRate = learningRateElement
				core.lambda = lambdaElement

				core.runMode("train")
				core.result.SaveLine()
				core.runMode("test")
				core.result.SaveLine()

				var params [][]byte
				readWriter := new(ReadWriter.ReadWriter)
				readWriter.Initialize(core.pathFinder("./Data/Result/Test"), true, true, true)

				for {
					byteLine := readWriter.ReadLine()
					if byteLine == nil {
						break
					}
					if len(byteLine) != 0 {
						params = bytes.Split(byteLine, []byte(";"))
					}
				}

				value, err := strconv.ParseInt(string(params[10]), 10, 64)
				if err == nil {
					correct = float64(value)
					accuracy = correct / 10000
				}

				value, err = strconv.ParseInt(string(params[11]), 10, 64)
				if err == nil {
					incorrect = float64(value)
				}

				if accuracy > accuracyOld {
					maxWeightRange = core.weightRange
					maxLearningRate = core.learningRate
					maxLambda = core.lambda
					correctMax = correct
					incorrectMax = incorrect
					accuracyOld = accuracy
				}
				os.Remove(core.pathFinder("./Data/Persistence/Test/Test"))
				os.Remove(core.pathFinder("./Data/Result/Test"))
			}
		}
	}

	fmt.Println("Best values:")
	fmt.Println("WeightRange: ", maxWeightRange)
	fmt.Println("LearningRate: ", maxLearningRate)
	fmt.Println("Lambda: ", maxLambda)
	fmt.Println("Correct: ", correctMax)
	fmt.Println("Incorrect: ", incorrectMax)
	fmt.Println("Accuracy: ", accuracyOld)
}*/
