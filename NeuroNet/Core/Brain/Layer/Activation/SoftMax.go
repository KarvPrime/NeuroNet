package Activation

import "math"

// Struct for softmax activation
type SoftMax struct {
}

// Classify the values
// Substracting the maximum is necessary to prevent exploding gradients (which would cause the numerical representation
// to exceed the limits, which then leads to +Inf and therefore with the next calculation NaN.
// It won't prevent vanishing gradients though.
func (softMax *SoftMax) Think(values []float64) {
	total := 0.0

	// Find max value
	max := values[0]
	for _, element := range values {
		if element > max {
			max = element
		}
	}

	// Largest value becomes 0 and the rest negative while summing up the values
	for index, element := range values {
		values[index] = math.Pow(math.E, element-max)
		total += values[index]
	}

	// Divide by total values
	for index, element := range values {
		values[index] = element / total
	}
}

// Calculating everything with a Jacobi Matrix and Kronecker Delta would be slow. If we derive, calculate the dot product,
// etc., we find out that we only need the i = j part of the whole calculation which speeds up computation times.
// So the derivative is (1 - x) * x, which we already know from the logistic regression.
func (softMax *SoftMax) Train(values []float64) {
	for index, element := range values {
		values[index] = (1 - element) * element
	}
}
