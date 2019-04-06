package PreProcessing

import "testing"

func TestMeanSubtraction_Process(t *testing.T) {
	values := []float64{2.0, 4.0, 6.0}
	valuesCheck := []float64{-2.0, 0.0, 2.0}

	meanSubtraction := new(MeanSubtraction)
	meanSubtraction.Process(values)

	for index, element := range values {
		if valuesCheck[index] != element {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
