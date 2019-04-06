package Activation

import (
	"math"
)

// Struct for logistic activation
type Logistic struct {
}

// Set the value to 1 / (1 + e^-x)
func (logistic *Logistic) Think(values []float64) {
	for index, element := range values {
		values[index] = 1 / (1 + math.Pow(math.E, -element))
	}
}

// Set the value to (1 - x) * x
func (logistic *Logistic) Train(values []float64) {
	for index, element := range values {
		values[index] = (1 - element) * element
	}
}
