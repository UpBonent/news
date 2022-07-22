package postgres

import (
	"context"
	"fmt"
	"github.com/UpBonent/news/src/layers/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

func NewClient(ctx context.Context, maxAttempts int, sc domain.StorageConfig, l *logrus.Logger) (db *sqlx.DB, err error) {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", sc.Username, sc.Password, sc.Host, sc.Database, sc.SSLMode)

	err = connect(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		db, err = sqlx.ConnectContext(ctx, "postgres", dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second, l)

	if err != nil {
		l.Error("error do with tries postgresql: %v", err)
	}
	return
}
