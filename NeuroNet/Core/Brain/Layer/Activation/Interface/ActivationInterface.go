package Interface

// Interface for different activations
type ActivationInterface interface {
	Think(values []float64)
	Train(values []float64)
}
