package logging

import "github.com/sirupsen/logrus"

func activatorLevels(l *logrus.Logger, levels []string) {
	for _, level := range levels {
		switch level {
		case "all":
			l.SetLevel(logrus.InfoLevel)
			l.SetLevel(logrus.ErrorLevel)
			l.SetLevel(logrus.FatalLevel)
			return
		case "info":
			l.SetLevel(logrus.InfoLevel)
		case "error":
			l.SetLevel(logrus.ErrorLevel)
		case "fatal":
			l.SetLevel(logrus.FatalLevel)
		}
	}
	return
}
