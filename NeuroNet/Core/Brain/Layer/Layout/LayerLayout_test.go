package Layout

import (
	"NeuroNet/Core/Brain/Layer/Activation"
	"testing"
)

func TestLayerLayout_Initialize(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "Logistic")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	if layerLayout.size != 3 {
		t.Errorf("Wrong size.")
	}
}

func TestLayerLayout_Initialize_None(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "None")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.None{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_Identity(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "Identity")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.Identity{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_Logistic(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "Logistic")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.Logistic{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_TanH(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "TanH")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.TanH{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_ReLU(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "ReLU")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.ReLU{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_LeakyReLU(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "LeakyReLU")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.LeakyReLU{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_ELU(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "ELU")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.ELU{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_SoftMax(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "SoftMax")

	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.SoftMax{}

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layerLayout.activation.Think(layerVar)
	layerLayout.activation.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_Initialize_Error(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "Error")

	if err == nil {
		t.Errorf("No Error thrown")
	}
}

func TestLayerLayout_GetActivation(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "Logistic")
	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	test := &Activation.Logistic{}
	layer := layerLayout.GetActivation()

	layerVar := []float64{1.2, 2.4}
	testVar := []float64{1.2, 2.4}

	layer.Think(layerVar)
	layer.Train(layerVar)
	test.Think(testVar)
	test.Train(testVar)

	if layerVar[0] != testVar[0] {
		t.Errorf("Activation different.")
	}
}

func TestLayerLayout_GetSize(t *testing.T) {
	layerLayout := new(LayerLayout)
	err := layerLayout.Initialize(3, true, "Logistic")
	if err != nil {
		t.Errorf("Error when initializing: " + err.Error())
	}

	if layerLayout.GetSize() != 3 {
		t.Errorf("Wrong size.")
	}
}
