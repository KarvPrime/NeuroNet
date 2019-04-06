package PreProcessing

import "testing"

func TestProportional_Process(t *testing.T) {
	values := []float64{2.0, 4.0}
	valuesCheck := []float64{0.5, 1.0}

	proportional := new(Proportional)
	proportional.Process(values)

	for index, element := range values {
		if valuesCheck[index] != element {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
