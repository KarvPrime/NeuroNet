package Brain

import (
	"NeuroNet/Core/Brain/Connection"
	"NeuroNet/Core/Brain/Layer/Layout"
	"math"
	"testing"
)

func TestNetwork_Initialize(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, true, "Logistic")
	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(4, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2, layerLayout3}

	network := new(Network)
	err := network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)
	if err != nil {
		t.Errorf("Initialization failed.")
	}

	if len(network.layers) != 3 {
		t.Errorf("Not enough layers.")
	}

	if len(network.connections) != 2 {
		t.Errorf("Not enough connections.")
	}

	if network.targetColumns != [2]int{0, 1} {
		t.Errorf("Wrong target columns.")
	}
	if network.targetInterpretation != "index" {
		t.Errorf("Wrong index target interpretation.")
	}
	if network.learningRate != 0.5 {
		t.Errorf("Wrong learn multiplier.")
	}

	network1 := new(Network)
	err = network1.Initialize(layerLayout, "None", [2]int{0, 4}, 0.5, 0.5)
	if err != nil {
		t.Errorf("Initialization failed.")
	}

	if network1.targetInterpretation != "direct" {
		t.Errorf("Wrong direct target interpretation.")
	}

	network2 := new(Network)
	err = network2.Initialize(layerLayout, "None", [2]int{0, 2}, 0.5, 0.5)
	if err == nil {
		t.Errorf("Initialization did not fail.")
	}
}

func TestNetwork_Initialize_None(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, true, "Logistic")
	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(4, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2, layerLayout3}

	network := new(Network)
	err := network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)
	if err != nil {
		t.Errorf("Initialization failed.")
	}

	network.Input([]float64{1, 2.0, 4.0, 6.0})
	valuesCheck := []float64{2.0, 4.0, 6.0}

	for index, element := range network.layers[0].GetNeurons() {
		if valuesCheck[index] != element {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestNetwork_Initialize_Proportional(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, true, "Logistic")
	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(4, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2, layerLayout3}

	network := new(Network)
	err := network.Initialize(layerLayout, "Proportional", [2]int{0, 1}, 0.5, 0.5)
	if err != nil {
		t.Errorf("Initialization failed.")
	}

	network.Input([]float64{1, 2.0, 4.0, 8.0})
	valuesCheck := []float64{0.25, 0.5, 1.0}

	for index, element := range network.layers[0].GetNeurons() {
		if valuesCheck[index] != element {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestNetwork_Initialize_MeanSubtraction(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, true, "Logistic")
	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(4, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2, layerLayout3}

	network := new(Network)
	err := network.Initialize(layerLayout, "MeanSubtraction", [2]int{0, 1}, 0.5, 0.5)
	if err != nil {
		t.Errorf("Initialization failed.")
	}

	network.Input([]float64{1, 2.0, 4.0, 6.0})
	valuesCheck := []float64{-2.0, 0, 2.0}

	for index, element := range network.layers[0].GetNeurons() {
		if valuesCheck[index] != element {
			t.Errorf("Expected value %f, different value %f.", valuesCheck[index], element)
		}
	}
}

func TestNetwork_Input(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, true, "Logistic")
	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(4, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2, layerLayout3}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	network.Input([]float64{1, 2.3, 4.6})
	targets := []float64{0, 1, 0, 0}
	for index, element := range network.targets {
		if element != targets[index] {
			t.Errorf("Target mismatch.")
		}
	}
	inputs := []float64{2.3, 4.6}
	for index, element := range network.inputs {
		if element != inputs[index] {
			t.Errorf("Input mismatch.")
		}
	}
}

func TestNetwork_Think(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "TanH")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(1, false, "None")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	network.connections[0].SetWeight(0, 0, 0.5)
	network.connections[0].SetWeight(1, 0, 0.8)

	network.Input([]float64{1, 2.5, 5.0})

	network.Think()

	if network.layers[1].GetNeurons()[0] != math.Tanh(2.5)*0.5+math.Tanh(5.0)*0.8 {
		t.Errorf("Output mismatch.")
	}
}

func TestNetwork_Train_Single(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "TanH")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(1, false, "TanH")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	network.connections[0].SetWeight(0, 0, 0.5)
	network.connections[0].SetWeight(1, 0, 0.8)

	network.Input([]float64{1, 2.5, 5.0})

	network.Think()

	network.Train()

	input := math.Tanh(2.5)*0.5 + math.Tanh(5.0)*0.8
	output := math.Tanh(input)
	delta := 1 - output
	train := 1 - output*output
	outputError := train * delta

	if network.connections[0].Weights[0][0] != 0.5+1.0*outputError*math.Tanh(2.5) {
		t.Errorf("Weight 0 mismatch.")
	}
	if network.connections[0].Weights[1][0] != 0.8+1.0*outputError*math.Tanh(5.0) {
		t.Errorf("Weight 1 mismatch.")
	}
}

func TestNetwork_Train_Multiple(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(5, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	network.connections[0].SetWeight(0, 0, 0.2)
	network.connections[0].SetWeight(0, 1, 0.3)
	network.connections[0].SetWeight(0, 2, 0.4)
	network.connections[0].SetWeight(1, 0, 0.2)
	network.connections[0].SetWeight(1, 1, 0.3)
	network.connections[0].SetWeight(1, 2, 0.4)
	network.connections[0].SetWeight(2, 0, 0.2)
	network.connections[0].SetWeight(2, 1, 0.3)
	network.connections[0].SetWeight(2, 2, 0.4)
	network.connections[0].SetWeight(3, 0, 0.2)
	network.connections[0].SetWeight(3, 1, 0.3)
	network.connections[0].SetWeight(3, 2, 0.4)
	network.connections[0].SetWeight(4, 0, 0.2)
	network.connections[0].SetWeight(4, 1, 0.3)
	network.connections[0].SetWeight(4, 2, 0.4)

	network.Input([]float64{1, 1.0, 2.0, 3.0, 4.0, 5.0})
	network.Think()
	network.Train()

	total := 0.0
	outputs := []float64{15 * 0.2, 15 * 0.3, 15 * 0.4}
	for index, element := range outputs {
		outputs[index] = math.Pow(math.E, element)
		total += outputs[index]
	}
	for index := range outputs {
		outputs[index] /= total
	}

	train := make([]float64, 3)
	outputError := make([]float64, 3)
	delta := make([]float64, len(network.targets))
	for index, element := range network.layers[len(network.layers)-1].GetNeurons() {
		delta[index] = network.targets[index] - element
		train[index] = (1 - element) * element
		outputError[index] = train[index] * delta[index]

		expected := 0.2 + 1.0*outputError[index]*outputs[index]
		if network.connections[0].Weights[0][index] != expected {
			t.Errorf("Weight 0 %d mismatch. Expected %f, got %f.", index, expected, network.connections[0].Weights[0][index])
		}
		expected = 0.3 + 1.0*outputError[index]*outputs[index]
		if network.connections[0].Weights[1][index] != expected {
			t.Errorf("Weight 1 %d mismatch. Expected %f, got %f.", index, expected, network.connections[0].Weights[1][index])
		}
		expected = 0.4 + 1.0*outputError[index]*outputs[index]
		if network.connections[0].Weights[2][index] != expected {
			t.Errorf("Weight 2 %d mismatch. Expected %f, got %f.", index, expected, network.connections[0].Weights[2][index])
		}
	}
}

func TestNetwork_Check(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(1, false, "Identity")

	layerLayoutDirect := []*Layout.LayerLayout{layerLayout1, layerLayout2}

	network := new(Network)
	network.Initialize(layerLayoutDirect, "None", [2]int{0, 1}, 0.5, 0.5)

	network.connections[0].SetWeight(0, 0, 0.4)
	network.connections[0].SetWeight(1, 0, 0.2)

	network.Input([]float64{1, 2.5, 5.0})
	network.Think()
	network.Check()

	if network.truthValues[true] != 0 || network.truthValues[false] != 1 {
		t.Errorf("Check failed.")
	}

	network.Input([]float64{4, 2.5, 5.0})
	network.Think()
	network.Check()

	if network.truthValues[true] != 1 || network.truthValues[false] != 1 {
		t.Errorf("Check failed.")
	}

	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(2, true, "TanH")
	layerLayout4 := new(Layout.LayerLayout)
	layerLayout4.Initialize(2, false, "TanH")

	layerLayoutIndex := []*Layout.LayerLayout{layerLayout3, layerLayout4}

	network1 := new(Network)
	network1.Initialize(layerLayoutIndex, "None", [2]int{0, 1}, 0.5, 0.5)

	network1.connections[0].SetWeight(0, 0, 0.5)
	network1.connections[0].SetWeight(0, 1, 0.7)
	network1.connections[0].SetWeight(1, 0, 0.8)
	network1.connections[0].SetWeight(1, 1, 0.9)

	network1.Input([]float64{1, 2.5, 5.0})
	network1.Think()
	network1.Check()

	if network1.truthValues[true] != 1 || network1.truthValues[false] != 0 {
		t.Errorf("Check failed.")
	}

	network1.Input([]float64{0, 2.5, 5.0})
	network1.Think()
	network1.Check()

	if network1.truthValues[true] != 1 || network1.truthValues[false] != 1 {
		t.Errorf("Check failed.")
	}
}

func TestNetwork_SetState(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, true, "Logistic")
	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(4, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2, layerLayout3}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	connection1 := new(Connection.Connection)
	connection1.Initialize(2, 3)
	connection1.SetWeight(0, 0, 3.14)
	connection2 := new(Connection.Connection)
	connection2.Initialize(3, 4)
	connection2.SetWeight(1, 1, 4.13)
	connections := []*Connection.Connection{connection1, connection2}

	network.SetState(connections)

	connection1.SetWeight(0, 0, 3.15)
	connection2.SetWeight(1, 1, 4.14)

	if network.connections[0].GetWeight(0, 0) != 3.14 {
		t.Errorf("Connection value mismatch.")
	}
	if network.connections[1].GetWeight(1, 1) != 4.13 {
		t.Errorf("Connection value mismatch.")
	}
}

func TestNetwork_GetState(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(3, true, "Logistic")
	layerLayout3 := new(Layout.LayerLayout)
	layerLayout3.Initialize(4, false, "SoftMax")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2, layerLayout3}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	connection1 := new(Connection.Connection)
	connection1.Initialize(2, 3)
	connection1.SetWeight(0, 0, 3.14)
	connection2 := new(Connection.Connection)
	connection2.Initialize(3, 4)
	connection2.SetWeight(1, 1, 4.13)
	connections := []*Connection.Connection{connection1, connection2}

	network.SetState(connections)

	state := network.GetState()

	if state[0].GetWeight(0, 0) != 3.14 {
		t.Errorf("Connection value mismatch.")
	}
	if state[1].GetWeight(1, 1) != 4.13 {
		t.Errorf("Connection value mismatch.")
	}
}

func TestNetwork_GetTruthValues(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(1, false, "Identity")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	network.connections[0].SetWeight(0, 0, 0.4)
	network.connections[0].SetWeight(1, 0, 0.2)

	network.Input([]float64{1, 2.5, 5.0})
	network.Think()
	network.Check()

	network.Input([]float64{4, 2.5, 5.0})
	network.Think()
	network.Check()

	truthValues := network.GetTruthValues()

	if truthValues[false] != 1 || truthValues[true] != 1 {
		t.Errorf("Truth values not correct.")
	}
}

func TestNetwork_reset(t *testing.T) {
	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(2, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(1, false, "Identity")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	network.connections[0].SetWeight(0, 0, 0.4)
	network.connections[0].SetWeight(1, 0, 0.2)

	network.Input([]float64{1, 2.5, 5.0})
	network.Think()

	network.reset()

	for _, element := range network.layers {
		for _, e := range element.GetNeurons() {
			if e != 0 {
				t.Errorf("Value not reset.")
			}
		}
	}
}

func TestNetwork_crossEntropyError(t *testing.T) {
	values := []float64{0.8, 0.08, 0.04, 0.02, 0.02, 0.01, 0.01, 0.01, 0.005, 0.005}
	target := []float64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	layerLayout1 := new(Layout.LayerLayout)
	layerLayout1.Initialize(5, true, "Identity")
	layerLayout2 := new(Layout.LayerLayout)
	layerLayout2.Initialize(10, false, "Identity")

	layerLayout := []*Layout.LayerLayout{layerLayout1, layerLayout2}

	network := new(Network)
	network.Initialize(layerLayout, "None", [2]int{0, 1}, 0.5, 0.5)

	network.layers[1].SetNeurons(values)
	network.targets = target

	// TODO: Write test
}
