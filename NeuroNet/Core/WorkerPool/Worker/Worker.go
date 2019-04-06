package Worker

import "NeuroNet/Core/Brain"

// Worker thread
func Worker(id int, mode string, network *Brain.Network, dataInput <-chan [][]float64, output chan<- int) {
	for input := range dataInput {
		inputs := len(input)
		for i := 0; i < inputs; i++ {
			network.Input(input[i])
			network.Think()
			if mode == "train" {
				network.Train()
			} else if mode == "test" {
				network.Check()
			}
		}
		output <- id
	}
}
