package util

import (
	"io"
	"log"
	"os"
)

var (
	ErrorLogger *log.Logger
)

func init() {
	logFile, err := os.OpenFile("htop.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
