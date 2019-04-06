package Layer

import (
	"NeuroNet/Core/Brain/Layer/Activation"
	"NeuroNet/Core/Brain/Layer/Layout"
	"testing"
)

func TestLayer_Initialize(t *testing.T) {
	// TODO: Test without bias
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	if layer.size != 2 {
		t.Errorf("Wrong size.")
	}

	if layer.realSize != 3 {
		t.Errorf("Wrong bias size.")
	}

	if len(layer.neurons) != 3 {
		t.Errorf("Neuron count different.")
	}

	test := &Activation.Logistic{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layer.activation.Think(layerVar)
	layer.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayer_ExciteNeuron(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	layer.ExciteNeuron(1, 2.5)
	if layer.neurons[1] != 2.5 {
		t.Errorf("Wrong value after first excite.")
	}

	layer.ExciteNeuron(1, 2.5)
	if layer.neurons[1] != 5.0 {
		t.Errorf("Wrong value after second excite.")
	}
}

func TestLayer_SetNeurons(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	layer.SetNeurons([]float64{1.5, 2.3})
}

func TestLayer_GetNeurons(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	layer.SetNeurons([]float64{1.5, 2.3})

	neurons := layer.GetNeurons()
	if neurons[0] != 1.5 {
		t.Errorf("Wrong first value.")
	}
	if neurons[1] != 2.3 {
		t.Errorf("Wrong second value.")
	}
}

func TestLayer_GetSize(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	if layer.GetSize() != 2 {
		t.Errorf("Wrong size.")
	}
}

func TestLayer_GetBiasSize(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	if layer.GetRealSize() != 3 {
		t.Errorf("Wrong bias size.")
	}
}

func TestLayer_Think(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	layer.SetNeurons([]float64{1.5, 2.3})
	test := Activation.Logistic{}
	testVar := []float64{1.5 / 3, 2.3 / 3}

	layer.Think()
	test.Think(testVar)

	if layer.neurons[0] != testVar[0] {
		t.Errorf("Wrong activation output in first neuron: %f is not %f.", layer.neurons[0], testVar[0])
	}
	if layer.neurons[1] != testVar[1] {
		t.Errorf("Wrong activation output in last neuron: %f is not %f.", layer.neurons[1], testVar[1])
	}
}

func TestLayer_Train(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	layer.SetNeurons([]float64{1.5, 2.3})
	test := Activation.Logistic{}
	testVar := []float64{1.5, 2.3, 0}

	layer.Think()
	for index := range testVar {
		testVar[index] /= float64(len(testVar))
	}
	test.Think(testVar)
	testVar[2] = 1

	delta := []float64{0.5, -0.3, 0.2}

	layer.Train(delta)
	test.Train(testVar)
	for index := range testVar {
		testVar[index] *= delta[index]
	}

	if layer.neurons[0] != testVar[0] {
		t.Errorf("Wrong backpropagation output in first neuron.")
	}
	if layer.neurons[1] != testVar[1] {
		t.Errorf("Wrong backpropagation output in middle neuron.")
	}
	if layer.neurons[2] != testVar[2] {
		t.Errorf("Wrong backpropagation output in last neuron.")
	}
}

func TestLayer_Reset(t *testing.T) {
	layerLayout := new(Layout.LayerLayout)
	layerLayout.Initialize(2, true, "Logistic")

	layer := new(Layer)
	layer.Initialize(layerLayout)

	layer.SetNeurons([]float64{1.5, 2.3})

	layer.Reset()

	if layer.neurons[0] != 0 {
		t.Errorf("Reset failed on fist element.")
	}
	if layer.neurons[1] != 0 {
		t.Errorf("Reset failed on last element.")
	}
}
