package main

import (
	"fmt"
	"log"
	"os"
)

var logFile *os.File

func initLog() {
	var err error
	logFile, err = os.OpenFile("log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Panicln(err)
	}
}

func printfLog(s string, a ...interface{}) {
	fmt.Fprintf(logFile, s, a...)
}
