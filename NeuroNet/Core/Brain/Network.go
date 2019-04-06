// Package Network contains the functionality of and sets up the neural network
package Brain

import (
	"NeuroNet/Core/Brain/Connection"
	"NeuroNet/Core/Brain/Layer"
	"NeuroNet/Core/Brain/Layer/Layout"
	"NeuroNet/Core/Brain/PreProcessing"
	"NeuroNet/Core/Brain/PreProcessing/Interface"
	"errors"
	"math"
)

// Struct for handling the input data
type Network struct {
	preProcessing                   Interface.PreProcessingInterface
	learningRate, lambda, errorRate float64
	targetColumns                   [2]int
	targetInterpretation            string

	targets, inputs []float64

	truthValues map[bool]int

	layers      []*Layer.Layer
	connections []*Connection.Connection
}

// Initialize the network to handle the data
func (network *Network) Initialize(layerLayout []*Layout.LayerLayout, preProcessing string, targetColumns [2]int, learningRate, lambda float64) error {
	network.targetColumns = targetColumns

	network.learningRate = learningRate
	network.lambda = lambda

	// Preprocessing of data
	switch preProcessing {
	case "None":
		network.preProcessing = &PreProcessing.None{}
		break
	case "Proportional":
		network.preProcessing = &PreProcessing.Proportional{}
		break
	case "MeanSubtraction":
		network.preProcessing = &PreProcessing.MeanSubtraction{}
		break
	default:
		return errors.New("incorrect pre processing: " + preProcessing + " not found")
		break
	}

	// Create the network
	network.create(layerLayout)

	// Find out which output pattern is needed
	outputs := len(network.layers[len(network.layers)-1].GetNeurons())
	targets := targetColumns[1] - targetColumns[0]
	if targets == 1 && outputs > 1 {
		network.targets = make([]float64, outputs)
		network.targetInterpretation = "index"
	} else if targets == outputs {
		network.targets = make([]float64, targets)
		network.targetInterpretation = "direct"
	} else {
		return errors.New("can't handle target interpretation")
	}

	// Create a map for the test results
	network.truthValues = make(map[bool]int)

	return nil
}

// Create a new network
func (network *Network) create(layerLayout []*Layout.LayerLayout) {
	// Create the layers
	network.layers = make([]*Layer.Layer, len(layerLayout))
	for index, element := range layerLayout {
		layer := new(Layer.Layer)
		layer.Initialize(element)

		network.layers[index] = layer
	}

	// Create the connectome
	connectionCount := len(network.layers) - 1
	network.connections = make([]*Connection.Connection, connectionCount)
	for i := 0; i < connectionCount; i++ {
		network.connections[i] = new(Connection.Connection)
		network.connections[i].Initialize(network.layers[i].GetRealSize(), network.layers[i+1].GetSize())
	}
}

// Input new data
func (network *Network) Input(input []float64) {
	targets, inputs := input[network.targetColumns[0]:network.targetColumns[1]], input[network.targetColumns[1]:]
	if network.targetInterpretation == "index" {
		for index := range network.targets {
			network.targets[index] = 0
		}
		network.targets[int(targets[0])] = 1
	} else if network.targetInterpretation == "direct" {
		network.targets = targets
	}

	network.inputs = inputs
	network.layers[0].SetNeurons(network.inputs)

	network.preProcessing.Process(network.layers[0].GetNeurons())
}

// Think with forward propagation
func (network *Network) Think() {
	network.layers[0].Think()
	connectionCount := len(network.layers) - 1
	for i := 0; i < connectionCount; i++ {
		for y := 0; y < network.layers[i+1].GetSize(); y++ {
			for x, ex := range network.layers[i].GetNeurons() {
				network.layers[i+1].ExciteNeuron(y, ex*network.connections[i].GetWeight(x, y))
			}
		}
		network.layers[i+1].Think()
	}
}

// Train with backward propagation
func (network *Network) Train() {
	// Get first delta
	errorValue := make([]float64, len(network.targets))

	for index, element := range network.layers[len(network.layers)-1].GetNeurons() {
		errorValue[index] = element - network.targets[index]
	}

	// Go backward through all layers
	for i := len(network.layers) - 1; i > 0; i-- {
		network.layers[i].Train(errorValue)

		errorValue = make([]float64, network.layers[i-1].GetRealSize())

		// Set new weights
		neurons := network.layers[i].GetNeurons()
		for x, output := range network.layers[i-1].GetNeurons() {
			for y := 0; y < network.layers[i].GetSize(); y++ {
				currentWeight := network.connections[i-1].GetWeight(x, y)

				// Get new delta
				errorValue[x] += neurons[y] * currentWeight

				// 2w
				currentWeight *= 2

				// sgn(w)
				if currentWeight > 0 {
					currentWeight += 1
				} else if currentWeight < 0 {
					currentWeight -= 1
				}

				// Add weight with elastic net regularization
				network.connections[i-1].UpdateWeight(x, y, network.learningRate*(neurons[y]*output-network.lambda*currentWeight))
			}
		}
	}

	network.reset()
}

// Check outputs
func (network *Network) Check() {
	neurons := network.layers[len(network.layers)-1].GetNeurons()
	if network.targetInterpretation == "index" {
		max := neurons[0]
		maxVal := 0

		for index, element := range neurons {
			if element > max {
				max = element
				maxVal = index
			}
			network.errorRate += 0.5 * math.Pow(network.targets[index]-element, 2)
			/*if network.targets[index] == 1 {
				network.errorRate -= math.Log(neurons[index])
			}*/
		}

		if network.targets[maxVal] == 1 {
			network.truthValues[true]++
		} else {
			network.truthValues[false]++
		}
	} else if network.targetInterpretation == "direct" {
		for index, element := range neurons {
			if network.targets[index]*0.95 < element && network.targets[index]*1.05 > element {
				network.truthValues[true]++
			} else {
				network.truthValues[false]++
			}
		}
	}

	network.reset()
}

// Quadratic loss
func (network *Network) quadraticLoss(currentWeight *float64) {
	// sgn(w)
	if *currentWeight > 0 {
		*currentWeight += 0.5
	} else if *currentWeight < 0 {
		*currentWeight -= 0.5
	}
}

// Cross entropy loss
func (network *Network) crossEntropyLoss(currentWeight *float64) {
	// 2w
	*currentWeight *= 2

	// sgn(w)
	if *currentWeight > 0 {
		*currentWeight += 1
	} else if *currentWeight < 0 {
		*currentWeight -= 1
	}
}

// Set weight state of network
func (network *Network) SetState(connections []*Connection.Connection) {
	for index, element := range network.connections {
		for x, ex := range element.GetState() {
			for y := range ex {
				network.connections[index].SetWeight(x, y, connections[index].GetWeight(x, y))
			}
		}
	}
}

// Get weight state of network
func (network *Network) GetState() []*Connection.Connection {
	return network.connections
}

// Get truth values
func (network *Network) GetTruthValues() map[bool]int {
	return network.truthValues
}

// Get network error
func (network *Network) GetResetError() float64 {
	networkError := network.errorRate
	network.errorRate = 0.0
	return networkError
}

// Reset layers
func (network *Network) reset() {
	for index := range network.layers {
		network.layers[index].Reset()
	}
}
