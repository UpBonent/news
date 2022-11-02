package logging

import (
	"fmt"
	"github.com/UpBonent/news/src/common/services"
	"io"
	"os"
)

type LoggerStruct struct {
	ActiveLevels levelActivate
	Output       services.LoggerOutput
}

type levelActivate struct {
	INFO, ERROR, FATAL bool
}

func (l *LoggerStruct) INFO(message string) {
	if l.ActiveLevels.INFO {
		toByte := fmt.Sprintf("INFO -- message: %v\n", message)
		_, _ = l.Output.Write([]byte(toByte))
	}
}

func (l *LoggerStruct) ERROR(message string) {
	if l.ActiveLevels.ERROR {
		toByte := fmt.Sprintf("ERROR -- message: %v\n", message)
		_, _ = l.Output.Write([]byte(toByte))
	}
}

func (l *LoggerStruct) FATAL(message string) {
	if l.ActiveLevels.FATAL {
		toByte := fmt.Sprintf("FATAL -- message: %v\n", message)
		_, _ = l.Output.Write([]byte(toByte))
		os.Exit(2)
	}
}

func NewLogger(logLevels []string, output io.Writer) services.Logger {
	active := activatorLevels(logLevels)
	return &LoggerStruct{active, output}
}
