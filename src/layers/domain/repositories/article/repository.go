package article

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {

}

func (r *Repository) Insert(ctx context.Context) error {
	r.db.Queryx()
}
