package services

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"github.com/labstack/echo"
)

type Application interface {
	CreateNewArticle(ctx context.Context, article models.Article, id int) error
	GetAllArticles(ctx context.Context) (articles []models.Article, err error)
	UpdateArticle(ctx context.Context, existArticle int, article models.Article) error
	GetArticlesByAuthorID(ctx context.Context, id int) (articles []models.Article, err error)

	CreateNewAuthor(ctx context.Context, author models.Author) (int, error)
	GetAllAuthors(ctx context.Context) (authors []models.Author, err error)
	GetAuthorByID(ctx context.Context, id int) (author models.Author, err error)
	GetIDByAuthor(ctx context.Context, author models.Author) (id int, err error)

	CheckUserExist(username, password string, c echo.Context) (ok bool, err error)
}
