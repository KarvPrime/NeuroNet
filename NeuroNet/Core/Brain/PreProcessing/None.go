package PreProcessing

// Struct for no pre processing
type None struct {
}

// Do nothing
func (none *None) Process(values []float64) {
}
