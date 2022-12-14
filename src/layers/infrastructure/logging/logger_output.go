package logging

import (
	"fmt"
	"github.com/UpBonent/news/src/common/services"
	"io"
	"os"
)

type OutPutToSTD struct{}

func (o *OutPutToSTD) Write(b []byte) (int, error) {
	n, err := os.Stdout.Write(b)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Got failed to write to the stdout: %v\n", err)
		os.Exit(2)
	}
	return n, err
}

type OutPutToFile struct {
	toFile io.Writer
}

func (o *OutPutToFile) Write(b []byte) (int, error) {
	n, err := o.toFile.Write(b)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Got failed to write to the file: %v\n", err)
		os.Exit(2)
	}
	return n, err
}

func setLoggerOutput(output, pathToFile string) services.LoggerOutput {
	switch output {
	case "file":
		logFile, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		return &OutPutToFile{logFile}
	default:
		return &OutPutToSTD{}
	}
}
