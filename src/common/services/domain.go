package services

//go:generate mockgen -source=domain.go -destination=mocks/mock.go

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"time"
)

type ArticleRepository interface {
	CreateNew(ctx context.Context, article models.Article, dateCreate time.Time, id int) (err error)
	GetAll(ctx context.Context) (articles []models.Article, err error)
	Update(ctx context.Context, existArticle int, article models.Article) error
	GetByAuthorID(ctx context.Context, id int) (articles []models.Article, err error)
}

type AuthorRepository interface {
	CreateNew(ctx context.Context, author models.Author) (int, error)
	GetAll(ctx context.Context) (authors []models.Author, err error)

	GetByID(ctx context.Context, id int) (author models.Author, err error)
	GetIDByName(ctx context.Context, author models.Author) (id int, err error)

	CheckExisting(username string) (bool, error)

	GetSalt(username string) (salt, passwordHash string, err error)
}

type Authentication interface {
	CheckUser(ctx context.Context, username, password string) (bool, error)
}
