package services

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"net/http"
)

type Application interface {
	CreateNewArticle(ctx context.Context, article models.Article, id int) error
	GetAllArticles(ctx context.Context) (articles []models.Article, err error)
	UpdateArticle(ctx context.Context, existArticle int, article models.Article) error
	GetArticlesByAuthorID(ctx context.Context, id int) (articles []models.Article, err error)

	CreateNewAuthor(ctx context.Context, author models.Author) (id int, c string, err error)
	GetAllAuthors(ctx context.Context) (authors []models.Author, err error)
	GetAuthorByID(ctx context.Context, id int) (author models.Author, err error)
	GetIDByAuthor(ctx context.Context, author models.Author) (id int, err error)

	CheckUserAuthentication(username, password string) (err error)
	SetUserCookie(cookieValue string) (newCookie http.Cookie)
	GetAuthorByCookie(cookieValue string) (author models.Author, err error)
	GetCookieByUserName(username string) (cookieValue string, err error)
}
