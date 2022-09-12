package rest

import (
	"context"
	"github.com/UpBonent/news/src/common/services"
	"github.com/UpBonent/news/src/layers/infrastructure/url"
)

func NewService(ctx context.Context, w url.Ways, r *services.Repository) *services.Service {
	return &services.Service{
		HomePageHandler: NewHandlerHomePage(w.Home),
		AuthorHandler:   NewHandlerAuthor(ctx, w.Author, r.AuthorRepository),
		ArticleHandler:  NewHandlerArticle(ctx, w.Article, r.ArticleRepository, r.AuthorRepository),
	}
}
