package WorkerPool

import (
	"NeuroNet/Core/Brain"
	"NeuroNet/Core/Brain/Connection"
	"NeuroNet/Core/Brain/Layer/Layout"
	"NeuroNet/Core/Logging"
	"NeuroNet/Core/WorkerPool/Worker"
	"strconv"
	"time"
)

// Struct for worker pool
type WorkerPool struct {
	mode          string
	targetColumns [2]int
	learningRate  float64
	lambda        float64

	preProcessing string
	networkLayout [][2]string

	parallel int

	dataChannel   chan [][]float64
	outputChannel chan int
}

// Initialize worker pool
func (workerPool *WorkerPool) Initialize(mode string, parallel int, targetColumns [2]int, learningRate, lambda float64, networkLayout [][2]string, preProcessing string, dataChannel chan [][]float64, outputChannel chan int) {
	workerPool.mode = mode
	workerPool.parallel = parallel
	workerPool.targetColumns = targetColumns
	workerPool.learningRate = learningRate
	workerPool.lambda = lambda
	workerPool.networkLayout = networkLayout
	workerPool.preProcessing = preProcessing
	workerPool.dataChannel = dataChannel
	workerPool.outputChannel = outputChannel
}

// Create worker pool
func (workerPool *WorkerPool) CreatePool() ([]*Brain.Network, []*Connection.Connection, float64, error) {
	start := time.Now()

	layerLayouts := make([]*Layout.LayerLayout, len(workerPool.networkLayout))
	last := len(workerPool.networkLayout) - 1

	for index, element := range workerPool.networkLayout {
		element1, err := strconv.Atoi(element[1])
		if err != nil {
			Logging.Panic("Could not convert neuron count to int: " + err.Error())
		}

		layout := new(Layout.LayerLayout)

		err = layout.Initialize(element1, index != last, element[0])
		if err != nil {
			Logging.Fatal("Could not initialize layer layout: " + err.Error())
		}

		layerLayouts[index] = layout
	}

	connections := workerPool.createConnections(layerLayouts)

	networks, err := workerPool.createNetworks(layerLayouts)
	if err != nil {
		return nil, nil, 0, err
	}

	setupTime := time.Since(start).Seconds()

	return networks, connections, setupTime, nil
}

// Create connections
func (workerPool *WorkerPool) createConnections(layerLayouts []*Layout.LayerLayout) []*Connection.Connection {
	connectionCount := len(layerLayouts) - 1
	connections := make([]*Connection.Connection, connectionCount)

	for i := 0; i < connectionCount; i++ {
		connections[i] = new(Connection.Connection)
		connections[i].Initialize(layerLayouts[i].GetRealSize(), layerLayouts[i+1].GetSize())
	}

	return connections
}

// Create workers and networks
func (workerPool *WorkerPool) createNetworks(layerLayouts []*Layout.LayerLayout) ([]*Brain.Network, error) {
	networks := make([]*Brain.Network, workerPool.parallel)
	for i := 0; i < workerPool.parallel; i++ {
		networks[i] = new(Brain.Network)
		err := networks[i].Initialize(layerLayouts, workerPool.preProcessing, workerPool.targetColumns, workerPool.learningRate, workerPool.lambda)
		if err != nil {
			return nil, err
		}

		go Worker.Worker(i, workerPool.mode, networks[i], workerPool.dataChannel, workerPool.outputChannel)
	}
	return networks, nil
}
