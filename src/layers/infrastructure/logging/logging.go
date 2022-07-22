package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type Logger struct {
	*logrus.Entry
}

var e *logrus.Entry

func GetLogger() *Logger {
	return &Logger{e}
}

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		_, err := w.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	return nil
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func NewLogger() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s in line:%d", filename, frame.Line)
		},
		DisableColors: true,
		FullTimestamp: true,
	}

	err := os.Mkdir("logs", 0777)
	if err != nil {
		if err.Error() != "mkdir logs: file exists" {
			fmt.Printf("text: %s ...\n", err)
			panic(err)
		}
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("discard")
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
