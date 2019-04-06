package Activation

import (
	"fmt"
	"math"
	"testing"
)

func TestSoftMax_Think(t *testing.T) {
	values := []float64{2, 4, 6, 8, 10, 5, 3, 9, 8, 1}
	total := 0.0
	valuesCheck := make([]float64, len(values))
	for index, element := range values {
		valuesCheck[index] = math.Pow(math.E, element)
		total += valuesCheck[index]
	}
	for index := range valuesCheck {
		valuesCheck[index] /= total
	}

	softMax := new(SoftMax)
	softMax.Think(values)

	for index, element := range values {
		deviation := valuesCheck[index] * 0.00000000001
		if element < valuesCheck[index]-deviation || element > valuesCheck[index]+deviation {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestSoftMax_Train(t *testing.T) {
	values := []float64{2.0, 4.0}
	valuesCheck := []float64{-2.0, -12.0}

	softMax := new(SoftMax)
	softMax.Train(values)

	for index, element := range values {
		if element != valuesCheck[index] {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestSoftMax(t *testing.T) {
	values := []float64{1, 5, 3}
	originalValues := make([]float64, len(values))
	for index, element := range values {
		originalValues[index] = element
	}

	softMax := new(SoftMax)
	softMax.Think(values)

	thinkValues := make([]float64, len(values))
	for index, element := range values {
		thinkValues[index] = element
	}

	softMax.Train(values)

	trainValues := make([]float64, len(values))
	for index, element := range values {
		trainValues[index] = element
	}

	fmt.Println("Original: ", originalValues)
	fmt.Println("Think: ", thinkValues)
	fmt.Println("Train: ", trainValues)

	jacobianValues := make([][]float64, len(values))
	for index := range jacobianValues {
		jacobianValues[index] = make([]float64, len(values))
	}

	for i, ei := range thinkValues {
		for j, ej := range thinkValues {
			if i == j {
				jacobianValues[i][j] = ei * (1 - ej)
			} else {
				jacobianValues[i][j] = -ei * ej
			}
		}
	}

	fmt.Println("Jacobian: ", jacobianValues)

}
