package Activation

import "math"

// Struct for Exponential Linear Unit activation
type ELU struct {
}

// When the value is smaller than 0, calculate alpha * (e^x - 1)
func (elu *ELU) Think(values []float64) {
	for index, element := range values {
		if element < 0 { // Spare check for equality - don't need to calculate alpha * 0, if it already is 0
			values[index] = 0.5 * (math.Pow(math.E, element) - 1) // TODO: Let user set alpha
		}
	}
}

// When smaller than 0, add 0.5, otherwise set it to 1
// We can simply add 0.5 as the array element still has the result of the original calculation and the derivative is
// this result + 0.5. (alpha * (e^x - 1))' = alpha * (e^x - 1) + alpha
func (elu *ELU) Train(values []float64) {
	for index, element := range values {
		if element > 0 {
			values[index] = 1
		} else {
			values[index] += 0.5 // TODO: Let user set alpha
		}
	}
}
