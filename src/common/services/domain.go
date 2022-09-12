package services

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
)

type ArticleRepository interface {
	Insert(ctx context.Context, article models.Article, id int) error
	All(ctx context.Context) (articles []models.Article, err error)
	Delete(ctx context.Context, id int) error
	UpDate(ctx context.Context, existArticle int, article models.Article) error
}

type AuthorRepository interface {
	Insert(ctx context.Context, author models.Author) error
	All(ctx context.Context) (authors []models.Author, err error)
	Delete(ctx context.Context, id int) error

	GetByID(ctx context.Context, id int) (author models.Author, err error)
	GetByName(ctx context.Context, author models.Author) (id int, err error)
}

type Repository struct {
	ArticleRepository
	AuthorRepository
}
