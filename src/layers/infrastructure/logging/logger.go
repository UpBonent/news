package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"path"
	"runtime"
)

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

func NewLogger(w io.Writer) *logrus.Entry {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s in line:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: false,
	}
	l.SetOutput(io.Discard)
	l.AddHook(&writerHook{
		Writer:    []io.Writer{w},
		LogLevels: logrus.AllLevels,
	})

	return logrus.NewEntry(l)
}
