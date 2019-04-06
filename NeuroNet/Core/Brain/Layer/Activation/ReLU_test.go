package Activation

import "testing"

func TestReLU_Think(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{2.6}

	reLU := new(ReLU)
	reLU.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestReLU_Think2(t *testing.T) {
	values := []float64{-2.6}
	valuesCheck := []float64{0}

	reLU := new(ReLU)
	reLU.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestReLU_Train(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{1}

	reLU := new(ReLU)
	reLU.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestReLU_Train2(t *testing.T) {
	values := []float64{-2.6}
	valuesCheck := []float64{0}

	reLU := new(ReLU)
	reLU.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
