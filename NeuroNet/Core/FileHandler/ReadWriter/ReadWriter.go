package ReadWriter

import (
	"NeuroNet/Core/Logging"
	"bufio"
	"io"
	"os"
)

// Struct for file handling
type ReadWriter struct {
	file   *os.File
	reader *bufio.Reader
	writer *bufio.Writer
}

// Count lines
func (readWriter *ReadWriter) CountLines() int {
	readWriter.file.Seek(0, 0)
	count := 0

	for {
		_, _, err := readWriter.reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				Logging.Fatal("Data reader error: " + err.Error())
			}
		}
		count++
	}
	readWriter.file.Seek(0, 0)

	return count
}

// Initialize the file handler
func (readWriter *ReadWriter) Initialize(filePath string, read, write, append bool) {
	modeInt := 0
	if read && write {
		modeInt = os.O_RDWR
	} else if read {
		modeInt = os.O_RDONLY
	} else if write {
		modeInt = os.O_WRONLY
	} else {
		// TODO: Error message
	}
	if append {
		modeInt += os.O_APPEND
	}

	file, err := os.OpenFile(filePath, modeInt, 0666)
	if err != nil {
		Logging.Fatal("File error: " + err.Error())
	}
	readWriter.file = file

	if read {
		readWriter.reader = bufio.NewReader(readWriter.file)
	}
	if write {
		readWriter.writer = bufio.NewWriter(readWriter.file)
	}
}

// Go to certain line
func (readWriter *ReadWriter) SetLine(line int) {
	for i := 0; i < line; i++ {
		_, _, err := readWriter.reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				Logging.Fatal("End of file reached: " + err.Error())
			} else {
				Logging.Fatal("Data reader error: " + err.Error())
			}
		}
	}
}

// Read a line
func (readWriter *ReadWriter) ReadLine() []byte {
	byteLine, isPrefix, err := readWriter.reader.ReadLine()

	if err != nil {
		if err == io.EOF {
			readWriter.file.Close()
		} else {
			Logging.Fatal("Data reader error: " + err.Error())
		}
		return nil
	}

	if isPrefix {
		Logging.Fatal("Data reader buffer overflow.")
	}

	return byteLine
}

// Write a line
func (readWriter *ReadWriter) WriteLine(line string) {
	_, err := readWriter.writer.WriteString(line + "\n")
	if err != nil {
		Logging.Fatal("Error writing data line" + err.Error())
	}
	if readWriter.writer.Flush() != nil {
		Logging.Fatal("Error writing to file" + err.Error())
	}
}

// Create file
func (readWriter *ReadWriter) CreateFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		os.Create(filePath)
		return true
	}
	return false
}
