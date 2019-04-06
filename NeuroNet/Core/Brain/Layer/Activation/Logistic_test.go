package Activation

import (
	"math"
	"testing"
)

func TestLogistic_Think(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{1 / (1 + math.Pow(math.E, -values[0]))}

	logistic := new(Logistic)
	logistic.Think(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestLogistic_Train(t *testing.T) {
	values := []float64{2.6}
	valuesCheck := []float64{(1 - values[0]) * values[0]}

	logistic := new(Logistic)
	logistic.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}
