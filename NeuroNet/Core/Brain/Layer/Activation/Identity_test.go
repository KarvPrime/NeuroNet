package Activation

import "testing"

func TestIdentity_Think(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{2.6}

	identity := new(Identity)
	identity.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestIdentity_Train(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{1}

	identity := new(Identity)
	identity.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
