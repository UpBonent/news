package postgres

import (
	"context"
	"fmt"
	"github.com/UpBonent/news/src/layers/domain"
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"
	"time"
)

var db *sqlx.DB

func NewClient(ctx context.Context, maxAttempts int, sc domain.StorageConfig, l *logging.Logger) *sqlx.DB {
	var once sync.Once
	once.Do(
		func() {
			var err error
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
				l.Panicf("problem with connect to the DB: [%v\n]. Chech status DB", err)
			}
		})

	return db
}
