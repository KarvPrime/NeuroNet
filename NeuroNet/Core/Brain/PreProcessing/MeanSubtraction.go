package PreProcessing

// Struct for mean subtraction pre processing
type MeanSubtraction struct {
}

// Substract the mean value from every value
func (meanSubtraction *MeanSubtraction) Process(values []float64) {
	min := values[0]
	max := values[0]
	for _, element := range values {
		if element > max {
			max = element
		} else if element < min {
			min = element
		}
	}

	mean := (max + min) / 2

	for index, element := range values {
		values[index] = element - mean
	}
}
