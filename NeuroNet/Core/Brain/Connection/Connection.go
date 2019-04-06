package Connection

// Struct for the connectome
type Connection struct {
	size    [2]int
	Weights [][]float64
}

// Initialize the connectome
func (connection *Connection) Initialize(from, to int) {
	connection.size = [2]int{from, to}
	connection.Weights = make([][]float64, from)
	for index := range connection.Weights {
		connection.Weights[index] = make([]float64, to)
	}
}

// Get the size of the connectome
func (connection *Connection) GetSize() [2]int {
	return connection.size
}

// Set a single weight
func (connection *Connection) SetWeight(x, y int, weight float64) {
	connection.Weights[x][y] = weight
}

// Add value to a single weight
func (connection *Connection) UpdateWeight(x, y int, weight float64) {
	connection.Weights[x][y] -= weight
}

// Get a single weight
func (connection *Connection) GetWeight(x, y int) float64 {
	return connection.Weights[x][y]
}

// Set the state of the connectome
func (connection *Connection) SetState(weights [][]float64) {
	for x, element := range connection.Weights {
		for y := range element {
			connection.Weights[x][y] = weights[x][y]
		}
	}
}

// Get the current state of the connectome
func (connection *Connection) GetState() [][]float64 {
	return connection.Weights
}

// Add the weighed differences (changes due to learning from another connectome) to the main connectome
func (connection *Connection) AddWeighedDiff(oldWeights, newWeights [][]float64, weight int) {
	for x := range connection.Weights {
		for y := range connection.Weights[x] {
			connection.Weights[x][y] += (newWeights[x][y] - oldWeights[x][y]) / float64(weight)
		}
	}
}
