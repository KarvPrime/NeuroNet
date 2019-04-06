package Activation

import "testing"

func TestLeakyReLU_Think(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{2.6}

	leakyReLU := new(LeakyReLU)
	leakyReLU.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestLeakyReLU_Think2(t *testing.T) {
	values := []float64{-2.6}
	valuesCheck := []float64{values[0] * 0.01}

	leakyReLU := new(LeakyReLU)
	leakyReLU.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestLeakyReLU_Train(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{1}

	leakyReLU := new(LeakyReLU)
	leakyReLU.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestLeakyReLU_Train2(t *testing.T) {
	values := []float64{-2.6}
	valuesCheck := []float64{0.01}

	leakyReLU := new(LeakyReLU)
	leakyReLU.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
