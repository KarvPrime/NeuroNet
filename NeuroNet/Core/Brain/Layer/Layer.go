package Layer

import (
	"NeuroNet/Core/Brain/Layer/Activation/Interface"
	"NeuroNet/Core/Brain/Layer/Layout"
)

// Struct for the layer
type Layer struct {
	activation Interface.ActivationInterface
	size       int
	realSize   int
	neurons    []float64
	bias       bool
}

// Initialize the layer
func (layer *Layer) Initialize(layout *Layout.LayerLayout) {
	layer.size = layout.GetSize()
	if layout.GetBias() {
		layer.bias = true
		layer.realSize = layer.size + 1
	} else {
		layer.realSize = layer.size
	}

	layer.neurons = make([]float64, layer.realSize)
	layer.activation = layout.GetActivation()
}

// Add a value to the neuron
func (layer *Layer) ExciteNeuron(index int, value float64) {
	layer.neurons[index] += value
}

// Set the values of all neurons for the layer
func (layer *Layer) SetNeurons(input []float64) {
	for index, element := range input {
		layer.neurons[index] = element
	}
}

// Get all neurons within the layer
func (layer *Layer) GetNeurons() []float64 {
	return layer.neurons
}

// Get the size without the bias neuron
func (layer *Layer) GetSize() int {
	return layer.size
}

// Get the size including the bias neuron
func (layer *Layer) GetRealSize() int {
	return layer.realSize
}

// Forward propagation
func (layer *Layer) Think() {
	layer.normalize()

	layer.activation.Think(layer.neurons)

	// Set bias
	if layer.bias {
		layer.neurons[layer.size] = 1
	}
}

// Backward propagation
func (layer *Layer) Train(errorValue []float64) {
	layer.activation.Train(layer.neurons)

	for index, element := range layer.neurons {
		layer.neurons[index] = errorValue[index] * element
	}
}

// Set all neurons to 0
func (layer *Layer) Reset() {
	// Reset all neurons to 0
	for index := range layer.neurons {
		layer.neurons[index] = 0
	}
}

// Normalize to prevent +Inf when adding neuron outputs in the next layer
func (layer *Layer) normalize() {
	for index := range layer.neurons {
		layer.neurons[index] /= float64(layer.realSize)
	}
}
