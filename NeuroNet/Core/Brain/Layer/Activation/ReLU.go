package Activation

// Struct for Rectified Linear Unit activation
type ReLU struct {
}

// If the value is smaller than 0, set it to 0
func (relu *ReLU) Think(values []float64) {
	for index, element := range values {
		if element < 0 { // Spare check for equality - don't need to set to 0, if it already is 0
			values[index] = 0
		}
	}
}

// If the element is smaller than 0, set it to 0, otherwise to 1
func (relu *ReLU) Train(values []float64) {
	for index, element := range values {
		if element > 0 {
			values[index] = 1
		} else {
			values[index] = 0
		}
	}
}
