// Created by	Karv Prime
// Twitter:		https://twitter.com/KarvPrime
// LinkedIn:	https://www.linkedin.com/in/karvprime/
// GitHub:		https://github.com/KarvPrime

// Package Persistence to save and load progress or results
package Persistence

import (
	"NeuroNet/Core/Brain/Connection"
	"NeuroNet/Core/Logging"
	"encoding/gob"
	"os"
	"strings"
)

// Struct for persistence handling
type Persistence struct {
	filePath    string
	TimeElapsed float64
	Connections []*Connection.Connection
	LineNumber  int
}

// Initialize persistence
func (persistence *Persistence) Initialize(filepath string) {
	_, err := os.Stat(filepath)
	if err != nil {
		splitPath := strings.Split(filepath, "/")
		rawPath := strings.Join(splitPath[:cap(splitPath)-1], "/")
		if os.MkdirAll(rawPath, os.ModePerm) != nil {
			Logging.Fatal("Could not create directories: " + err.Error())
		}
	}

	persistence.filePath = filepath
	_, err = os.Stat(persistence.filePath)
	if err != nil {
		os.Create(persistence.filePath)
	}
}

// Get persistence data
func (persistence *Persistence) Read() (int, float64, []*Connection.Connection) {
	file, err := os.OpenFile(persistence.filePath, os.O_RDONLY, 0666)
	if err != nil {
		Logging.Fatal("Persistence could not be loaded: " + err.Error())
	}

	fileStat, err := file.Stat()
	if err != nil {
		Logging.Fatal("Could not read file stats: " + err.Error())
	}
	if fileStat.Size() != 0 {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(persistence)
		if err != nil {
			Logging.Fatal("Persistence decoding error: " + err.Error())
		}
	} else {
		persistence.LineNumber = 0
		persistence.TimeElapsed = 0.0
		persistence.Connections = nil
	}
	file.Close()

	return persistence.LineNumber, persistence.TimeElapsed, persistence.Connections
}

// Set persistence data
func (persistence *Persistence) Write(lineNumber int, timeElapsed float64, connections []*Connection.Connection) {
	persistence.LineNumber = lineNumber
	persistence.TimeElapsed = timeElapsed
	persistence.Connections = connections

	file, err := os.OpenFile(persistence.filePath, os.O_WRONLY, 0666)
	if err != nil {
		Logging.Fatal("Persistence could not be loaded: " + err.Error())
	}

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(persistence)
	if err != nil {
		Logging.Fatal("Persistence encoding error: " + err.Error())
	}
	file.Close()
}
