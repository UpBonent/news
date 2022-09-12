package repositories

import (
	"github.com/UpBonent/news/src/common/services"
	"github.com/UpBonent/news/src/layers/domain/repositories/article"
	"github.com/UpBonent/news/src/layers/domain/repositories/author"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *services.Repository {
	return &services.Repository{
		ArticleRepository: article.NewRepository(db),
		AuthorRepository:  author.NewRepository(db),
	}
}
