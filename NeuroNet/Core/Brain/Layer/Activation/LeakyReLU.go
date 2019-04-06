package Activation

// Struct for Leaky Rectified Linear Unit activation
type LeakyReLU struct {
}

// If the value is smaller than 0, multiply it with 0.01
func (leakyReLU *LeakyReLU) Think(values []float64) {
	for index, element := range values {
		if element < 0 { // Spare check for equality - don't need to calculate 0 * 0.01, if it already is 0
			values[index] = element * 0.01
		}
	}
}

// If the value is smaller than 0, set it to 0.01, otherwise set it to 1
func (leakyReLU *LeakyReLU) Train(values []float64) {
	for index, element := range values {
		if element > 0 {
			values[index] = 1
		} else {
			values[index] = 0.01
		}
	}
}
