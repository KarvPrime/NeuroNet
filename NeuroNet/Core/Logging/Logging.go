package Logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Initialize the logger
func Initialize(mainPath string) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	logfile, err := os.OpenFile(mainPath+"/Data/Log/"+time.Now().Format("2006_01_02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Fatal("Could not open logfile: " + err.Error())
	}
	log.SetOutput(logfile)
}

// For info messages
func Info(line string) {
	fmt.Println(line)
	log.Println(line)
}

// For errors causing a panic
func Panic(line string) {
	fmt.Println("\033[01;31m", line, "\033[00;00m")
	log.Panicln(line)
}

// For unrecoverable fatal errors
func Fatal(line string) {
	fmt.Println("\033[00;31m", line, "\033[00;00m")
	log.Fatalln(line)
}
