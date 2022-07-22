package postgres

import (
	"github.com/sirupsen/logrus"
	"time"
)

func connect(fn func() error, attempts int, delay time.Duration, l *logrus.Logger) (err error) {

	for attempts > 0 {
		if err = fn(); err != nil {
			l.Errorf("problem with connect to the DB: [%v\n].", err)
			time.Sleep(delay)
			l.Info("new attempt to connect to the DB")
			attempts--
		} else {
			return nil
		}
	}
	//errorf or fatalf?
	l.Errorf("problem with connect to the DB: [%v\n]. Chech status DB", err)
	return
}
