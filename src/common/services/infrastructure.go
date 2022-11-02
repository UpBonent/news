package services

import "io"

type Logger interface {
	INFO(message string)
	ERROR(message string)
	FATAL(message string)
}

type LoggerOutput = io.Writer
