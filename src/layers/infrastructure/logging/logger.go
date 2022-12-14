package logging

import (
	"fmt"
	"github.com/UpBonent/news/src/layers/infrastructure/config"
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

func newLogger(w io.Writer) *logrus.Logger {
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
	l.SetOutput(w)
	l.AddHook(&writerHook{
		Writer:    []io.Writer{w},
		LogLevels: logrus.AllLevels,
	})

	return logrus.NewEntry(l).Logger
}

func NewLogger(cfg config.Log) *logrus.Logger {
	w := setLoggerOutput(cfg.Output, cfg.PathToFile)
	l := newLogger(w)
	activatorLevels(l, cfg.ActiveLevels)
	return l
}
