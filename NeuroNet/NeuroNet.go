// Package main contains the main function
package main

import "NeuroNet/Core"

// Main function to run the neural network
func main() {
	core := new(Core.Core)
	core.Start("./Batchfile/Batch")

	// Debugging neural networks... (┘ò_Ó)┘ ~ ┴──┴

	// TODO: Write missing tests
	// TODO: Real batch (multiple forward passes, one backward pass), differentiate from data batch
	// TODO: Show approximate time until complete
	// TODO: Recovery functionality after small errors
}
