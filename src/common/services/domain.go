package services

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
)

type ArticleRepository interface {
	Insert(ctx context.Context, article models.Article, author models.Author) error
	All(ctx context.Context) (articles []models.Article, err error)
	Delete(ctx context.Context, articles models.Article) error
	UpDate(ctx context.Context, existArticle int, article models.Article) error
}

type AuthorRepository interface {
	Insert(ctx context.Context, a models.Author) error
	All(ctx context.Context) (authors []models.Author, err error)
	Delete(ctx context.Context, a models.Author) error

	GetByID(ctx context.Context, id int) (author models.Author, err error)
}
