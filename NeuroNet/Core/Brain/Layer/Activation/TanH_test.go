package Activation

import (
	"math"
	"testing"
)

func TestTanH_Think(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{math.Tanh(values[0])}

	tanH := new(TanH)
	tanH.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestTanH_Train(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{1 - values[0]*values[0]}

	tanH := new(TanH)
	tanH.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
