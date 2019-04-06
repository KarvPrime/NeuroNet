package FirstOrderMethods

import "math"

type AdaDelta struct {
	runningAverage, lastRMS, rho, rhoComplement, epsilon float64
	runningAverageArray                                  []float64
	iterator                                             int
}

func (adaDelta *AdaDelta) Initialize(rho, epsilon float64) {
	adaDelta.runningAverageArray = make([]float64, 10)
	// TODO: Let user set rho
	adaDelta.rho = rho
	adaDelta.rhoComplement = 1 - adaDelta.rho
	// TODO: Let user set epsilon
	adaDelta.epsilon = epsilon
}

func (adaDelta *AdaDelta) CalculateGradient(gradient float64) float64 {
	data := adaDelta.rho*adaDelta.runningAverage + adaDelta.rhoComplement*gradient*gradient

	adaDelta.runningAverage += data - adaDelta.runningAverageArray[adaDelta.iterator]
	adaDelta.runningAverageArray[adaDelta.iterator] = data

	if adaDelta.iterator < 9 {
		adaDelta.iterator++
	} else {
		adaDelta.iterator = 0
	}

	RMS := math.Sqrt(adaDelta.runningAverage + 0.1)

	update := adaDelta.lastRMS / RMS * gradient

	adaDelta.lastRMS = RMS

	return update
}
