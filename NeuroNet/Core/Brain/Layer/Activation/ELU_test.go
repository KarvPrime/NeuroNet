package Activation

import (
	"math"
	"testing"
)

func TestELU_Think(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{2.6}

	elu := new(ELU)
	elu.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestELU_Think2(t *testing.T) {
	values := []float64{-2.6}
	valuesCheck := []float64{0.5 * (math.Pow(math.E, values[0]) - 1)}

	elu := new(ELU)
	elu.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestELU_Train(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{1}

	elu := new(ELU)
	elu.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestELU_Train2(t *testing.T) {
	values := []float64{-2.6}
	valuesCheck := []float64{values[0] + 0.5}

	elu := new(ELU)
	elu.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
