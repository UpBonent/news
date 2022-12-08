package services

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
)

type Application interface {
	CreateNewArticle(ctx context.Context, article models.Article, id int) error
	GetAllArticles(ctx context.Context) (articles []models.Article, err error)
	UpdateArticle(ctx context.Context, existArticle int, article models.Article) error
	GetArticlesByAuthorID(ctx context.Context, id int) (articles []models.Article, err error)

	CreateNewAuthor(ctx context.Context, author models.Author, checkPassword string) (id int, err error)
	GetAllAuthors(ctx context.Context) (authors []models.Author, err error)
	GetAuthorByID(ctx context.Context, id int) (author models.Author, err error)
	GetIDByAuthor(ctx context.Context, author models.Author) (id int, err error)

	CheckUserAuthentication(username, password string) (err error)
	CheckUserExisting(username string) (bool, error)
}
