package Activation

import "math"

// Struct for TanH activation
type TanH struct {
}

// Set values to TanH(x)
func (tanH *TanH) Think(values []float64) {
	for index, element := range values {
		values[index] = math.Tanh(element)
	}
}

// Set values to 1 - x²
func (tanH *TanH) Train(values []float64) {
	for index, element := range values {
		values[index] = 1 - element*element
	}
}
