package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/UpBonent/news/src/layers/infrastructure/config"
)

func NewClient(ctx context.Context, sc config.StorageConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", sc.Username, sc.Password, sc.Host, sc.Database, sc.SSLMode)
	db, err = sqlx.ConnectContext(ctx, "postgres", dsn)
	return nil, err
}
