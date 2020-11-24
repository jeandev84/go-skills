package main

import (
	"log"
	"os"
)

var (
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
)

func initLoggers() {

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
}
