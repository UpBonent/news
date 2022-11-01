package services

//go:generate mockgen -source=domain.go -destination=mocks/mock.go

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
)

type ArticleRepository interface {
	Insert(ctx context.Context, article models.Article, id int) error
	All(ctx context.Context) (articles []models.Article, err error)
	Update(ctx context.Context, existArticle int, article models.Article) error
	GetByAuthorID(ctx context.Context, id int) (articles []models.Article, err error)
}

type AuthorRepository interface {
	Insert(ctx context.Context, author models.Author) (int, error)
	GetAll(ctx context.Context) (authors []models.Author, err error)

	GetAuthorByID(ctx context.Context, id int) (author models.Author, err error)
	GetIDByName(ctx context.Context, author models.Author) (id int, err error)
}
