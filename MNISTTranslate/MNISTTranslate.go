package main

import (
	"NeuroNet/Core/Logging"
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"os"
	"path"
	"runtime"
	"strconv"
)

func parse(labelFile, imageFile, outputFile string, lines int) {
	labelFileO, _ := os.OpenFile(labelFile, os.O_RDONLY, 0666)
	imageFileO, _ := os.OpenFile(imageFile, os.O_RDONLY, 0666)
	outputFileO, _ := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0666)

	gzipLabelReader, _ := gzip.NewReader(labelFileO)
	gzipImageReader, _ := gzip.NewReader(imageFileO)

	labelFileReader := bufio.NewReader(gzipLabelReader)
	imageFileReader := bufio.NewReader(gzipImageReader)
	outputFileWriter := bufio.NewWriter(outputFileO)

	var checksum int32

	binary.Read(labelFileReader, binary.BigEndian, &checksum)
	if checksum != 0x00000801 {
		panic("Label file checksum mismatch.")
	}
	binary.Read(imageFileReader, binary.BigEndian, &checksum)
	if checksum != 0x00000803 {
		panic("Image file checksum mismatch.")
	}

	binary.Read(labelFileReader, binary.BigEndian, &checksum)
	if checksum != int32(lines) {
		panic("Label file length mismatch.")
	}
	binary.Read(imageFileReader, binary.BigEndian, &checksum)
	if checksum != int32(lines) {
		panic("Image file length mismatch.")
	}

	binary.Read(imageFileReader, binary.BigEndian, &checksum)
	if checksum != int32(28) {
		panic("Label file row mismatch.")
	}

	binary.Read(imageFileReader, binary.BigEndian, &checksum)
	if checksum != int32(28) {
		panic("Image file column mismatch.")
	}

	n := 28*28 + 1
	var number byte

	for i := 0; i < lines; i++ {
		binary.Read(labelFileReader, binary.BigEndian, &number)
		output := strconv.Itoa(int(number))
		for j := 1; j < n; j++ {
			binary.Read(imageFileReader, binary.BigEndian, &number)
			output += ","
			output += strconv.Itoa(int(number))
		}

		outputFileWriter.WriteString(output + "\n")
	}
	outputFileWriter.Flush()

	gzipLabelReader.Close()
	gzipImageReader.Close()

	labelFileO.Close()
	imageFileO.Close()
	outputFileO.Close()
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		Logging.Fatal("Could not get file path.")
	}
	root := path.Dir(filename)

	parse("train-labels-idx1-ubyte.gz", "train-images-idx3-ubyte.gz", root+"/mnist_train.csv", 60000)
	parse("t10k-labels-idx1-ubyte.gz", "t10k-images-idx3-ubyte.gz", root+"/mnist_test.csv", 10000)
}
