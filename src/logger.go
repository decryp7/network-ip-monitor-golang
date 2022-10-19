package main

import (
	"fmt"
)

type logger struct {
	//log *eventlog.Log
}

func (l *logger) init() bool {
	// log, err := eventlog.Open("nwipmon")

	// if err != nil {
	// 	fmt.Println("Unable to open nwipmon eventlog.")
	// 	return false
	// }

	// l.log = log
	return true
}

func (l *logger) writeLog(logMessage string) {
	// if l.log == nil && !l.init() {
	// 	return
	// }

	//l.log.Info(1000, logMessage)
	fmt.Printf("[INFO] %v\n", logMessage)
}

func (l *logger) writeError(errorMessage string) {
	// if l.log == nil && !l.init() {
	// 	return
	// }

	//l.log.Error(1000, errorMessage)
	fmt.Printf("[ERROR] %v\n", errorMessage)
}
