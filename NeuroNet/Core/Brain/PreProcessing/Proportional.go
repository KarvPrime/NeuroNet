package PreProcessing

// Struct for proportional pre processing
type Proportional struct {
}

// Divide every value by the max value
func (proportional *Proportional) Process(values []float64) {
	max := values[0]
	for _, element := range values {
		if element > max {
			max = element
		}
	}

	for index, element := range values {
		values[index] = element / max
	}
}
