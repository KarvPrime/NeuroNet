package FileHandler

import (
	"NeuroNet/Core/FileHandler/ReadWriter"
	"NeuroNet/Core/Logging"
	"bytes"
	"strconv"
	"strings"
)

// Struct to read files
type FileHandler struct {
	readWriter *ReadWriter.ReadWriter
	valueSet   [][2]string
}

// Initialize batch handler
func (fileHandler *FileHandler) Initialize(filePath string, read, write, append, create bool) {
	if create {
		fileHandler.readWriter.CreateFile(filePath)
	}

	fileHandler.readWriter = new(ReadWriter.ReadWriter)
	fileHandler.readWriter.Initialize(filePath, read, write, append)
}

// Count lines when needed for reading line by line
/*
func (fileHandler *FileHandler) CountLines() int {
	return fileHandler.readWriter.CountLines()
}*/

// Read batchfile
func (fileHandler *FileHandler) ReadBatch(delimiter string) [][2]string {
	var lines [][2]string
	for {
		byteLine := fileHandler.readWriter.ReadLine()
		if byteLine == nil {
			break
		}
		if len(byteLine) != 0 {
			params := bytes.Split(byteLine, []byte(delimiter))
			lines = append(lines, [2]string{string(params[0]), string(params[1])})
		}
	}
	return lines
}

// Add line
func (fileHandler *FileHandler) AddValue(name, data string) {
	fileHandler.valueSet = append(fileHandler.valueSet, [2]string{name, data})
}

// Save the value line
func (fileHandler *FileHandler) SaveLine() {
	// Make csv header
	if fileHandler.readWriter.CountLines() == 0 {
		var header []string
		for _, element := range fileHandler.valueSet {
			header = append(header, element[0])
		}
		fileHandler.readWriter.WriteLine(strings.Join(header, ";"))
	}

	var parts []string
	for _, element := range fileHandler.valueSet {
		parts = append(parts, element[1])
	}
	fileHandler.readWriter.WriteLine(strings.Join(parts, ";"))

	// Clear value set
	fileHandler.valueSet = [][2]string{}
}

// Set line number when saving line by line
/*
func (fileHandler *FileHandler) SetLine(line int) {
	fileHandler.readWriter.SetLine(line)
}*/

// Read next line of CSV data
/*func (fileHandler *FileHandler) NextLine() []float64 {
	byteLine := fileHandler.readWriter.ReadLine()

	if byteLine == nil {
		return nil
	}

	params := bytes.Split(byteLine, []byte(","))
	inputs := make([]float64, len(params))
	for i := 0; i < len(params); i++ {
		param := string(params[i])
		value, err := strconv.ParseFloat(param, 64)
		if err != nil {
			Logging.Fatal("Data float conversion error: " + err.Error())
		}
		inputs[i] = value
	}
	return inputs
}*/

// Read all lines
func (fileHandler *FileHandler) ReadData() [][]float64 {
	var lines [][]float64
	for {
		byteLine := fileHandler.readWriter.ReadLine()
		if byteLine == nil {
			break
		}
		if len(byteLine) != 0 {
			params := bytes.Split(byteLine, []byte(","))
			inputs := make([]float64, len(params))
			for i := 0; i < len(params); i++ {
				param := string(params[i])
				value, err := strconv.ParseFloat(param, 64)
				if err != nil {
					Logging.Fatal("Data float conversion error: " + err.Error())
				}
				inputs[i] = value
			}
			lines = append(lines, inputs)
		}
	}
	return lines
}
