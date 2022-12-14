package services

import "io"

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type LoggerOutput = io.Writer
