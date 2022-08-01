package postgres

import (
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"time"
)

func connect(fn func() error, attempts int, delay time.Duration, l *logging.Logger) (err error) {

	for attempts > 0 {
		if err = fn(); err != nil {
			l.Errorf("problem with connect to the db: [%v\n].", err)
			time.Sleep(delay)
			l.Info("new attempt to connect to the db")
			attempts--
		} else {
			return nil
		}
	}
	return err
}
