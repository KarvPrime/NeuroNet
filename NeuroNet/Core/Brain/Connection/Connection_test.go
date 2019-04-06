package Connection

import "testing"

func TestConnection_Initialize(t *testing.T) {
	connection := new(Connection)
	connection.Initialize(2, 3)

	if connection.size[0] != 2 {
		t.Errorf("Wrong x value.")
	}
	if connection.size[1] != 3 {
		t.Errorf("Wrong y value.")
	}

	if len(connection.Weights) != 2 {
		t.Errorf("Wrong x length.")
	}
	if len(connection.Weights[0]) != 3 {
		t.Errorf("Wrong y length.")
	}
}

func TestConnection_GetSize(t *testing.T) {
	connection := new(Connection)
	connection.Initialize(2, 3)

	if connection.GetSize() != [2]int{2, 3} {
		t.Errorf("Wrong return value.")
	}
}

func TestConnection_SetSingle(t *testing.T) {
	connection := new(Connection)
	connection.Initialize(2, 3)

	connection.SetWeight(1, 2, 3.4)
	if connection.Weights[1][2] != 3.4 {
		t.Errorf("Wrong weight.")
	}
}

func TestConnection_AddSingle(t *testing.T) {
	connection := new(Connection)
	connection.Initialize(2, 3)

	connection.UpdateWeight(1, 2, 3.4)
	if connection.Weights[1][2] != -3.4 {
		t.Errorf("Wrong starting weight.")
	}

	connection.UpdateWeight(1, 2, 3.4)
	if connection.Weights[1][2] != -6.8 {
		t.Errorf("Wrong weight.")
	}
}

func TestConnection_GetSingle(t *testing.T) {
	connection := new(Connection)
	connection.Initialize(2, 3)

	if connection.GetWeight(1, 2) != 0 {
		t.Errorf("Wrong starting weight.")
	}

	connection.SetWeight(1, 2, 3.4)

	if connection.GetWeight(1, 2) != 3.4 {
		t.Errorf("Wrong weight.")
	}
}

func TestConnection_SetState(t *testing.T) {
	connection := new(Connection)
	connection.Initialize(2, 3)

	weights := make([][]float64, 2)
	for index := range weights {
		weights[index] = make([]float64, 3)
	}

	for index, element := range weights {
		for i := range element {
			weights[index][i] = float64(index + i)
		}
	}

	connection.SetState(weights)

	if &weights == &connection.Weights {
		t.Errorf("Same address.")
	}

	for index, element := range connection.Weights {
		for i := range element {
			if connection.Weights[index][i] != weights[index][i] {
				t.Errorf("Weights different.")
			}
		}
	}
}

func TestConnection_GetState(t *testing.T) {
	connection := new(Connection)
	connection.Initialize(2, 3)

	weights := make([][]float64, 2)
	for index := range weights {
		weights[index] = make([]float64, 3)
	}

	for index, element := range weights {
		for i := range element {
			weights[index][i] = float64(index + i)
		}
	}

	connection.SetState(weights)

	getWeights := connection.GetState()

	for index, element := range getWeights {
		for i := range element {
			if getWeights[index][i] != weights[index][i] {
				t.Errorf("Weights different.")
			}
		}
	}
}
