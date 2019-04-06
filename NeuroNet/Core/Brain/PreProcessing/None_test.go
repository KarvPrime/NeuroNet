package PreProcessing

import "testing"

func TestNone_Process(t *testing.T) {
	values := []float64{2.0, 4.0}
	valuesCheck := []float64{2.0, 4.0}

	none := new(None)
	none.Process(values)

	for index, element := range values {
		if valuesCheck[index] != element {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
