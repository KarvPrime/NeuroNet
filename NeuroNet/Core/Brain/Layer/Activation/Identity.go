package Activation

// Struct for identity activation
type Identity struct {
}

// Don't change the values
func (identity *Identity) Think(values []float64) {
}

// Set all values to 1
func (identity *Identity) Train(values []float64) {
	for index := range values {
		values[index] = 1
	}
}
