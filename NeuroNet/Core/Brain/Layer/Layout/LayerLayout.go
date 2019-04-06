package Layout

import (
	"NeuroNet/Core/Brain/Layer/Activation"
	"NeuroNet/Core/Brain/Layer/Activation/Interface"
	"errors"
)

// Struct for the layer layout
type LayerLayout struct {
	size       int
	realSize   int
	activation Interface.ActivationInterface
	bias       bool
}

// Initialize the layout
func (layout *LayerLayout) Initialize(size int, bias bool, activation string) error {
	layout.size = size
	if bias {
		layout.realSize = size + 1
	}
	layout.bias = bias

	// Use the correct activation function
	switch activation {
	case "None":
		layout.activation = &Activation.None{}
		break
	case "Identity":
		layout.activation = &Activation.Identity{}
		break
	case "Logistic":
		layout.activation = &Activation.Logistic{}
		break
	case "TanH":
		layout.activation = &Activation.TanH{}
		break
	case "ReLU":
		layout.activation = &Activation.ReLU{}
		break
	case "LeakyReLU":
		layout.activation = &Activation.LeakyReLU{}
		break
	case "ELU":
		layout.activation = &Activation.ELU{}
		break
	case "SoftMax":
		layout.activation = &Activation.SoftMax{}
		break
	default:
		return errors.New("incorrect activation: " + activation + " not found")
		break
	}
	return nil
}

// Get the activation function
func (layout *LayerLayout) GetActivation() Interface.ActivationInterface {
	return layout.activation
}

// Get the size without bias neuron, which we don't need when layer x (w/o bias) backpropagates to layer x-1 (w bias)
func (layout *LayerLayout) GetSize() int {
	return layout.size
}

// Get the size including bias neuron, which we need for the x-1 layer to correctly calculate the connectome
func (layout *LayerLayout) GetRealSize() int {
	return layout.realSize
}

// Get the bias
func (layout *LayerLayout) GetBias() bool {
	return layout.bias
}
