package Activation

import "testing"

func TestNone_Think(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{2.6}

	none := new(None)
	none.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestNone_Train(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{2.6}

	none := new(None)
	none.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
