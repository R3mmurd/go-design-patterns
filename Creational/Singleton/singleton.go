package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	loglevel int
}

// Log a message using the logger
func (l *MyLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

// SetLogLevel sets the Log level of the Logger
func (l *MyLogger) SetLogLevel(level int) {
	l.loglevel = level
}

// the Logger instance
var logger *MyLogger

// a variable to enforce goroutine safety
var once sync.Once

func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating MyLogger instance")
		logger = &MyLogger{}
	})
	fmt.Println("Returning MyLogger instance")
	return logger
}
