package application

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
)

func (a *Application) CreateNewArticle(ctx context.Context, article models.Article, id int) error {
	return a.Article.CreateNew(ctx, article, id)
}

func (a *Application) GetAllArticles(ctx context.Context) (articles []models.Article, err error) {
	return a.Article.GetAll(ctx)
}

func (a *Application) UpdateArticle(ctx context.Context, existArticle int, article models.Article) error {
	return a.Article.Update(ctx, existArticle, article)
}

func (a *Application) GetArticlesByAuthorID(ctx context.Context, id int) (articles []models.Article, err error) {
	return a.Article.GetByAuthorID(ctx, id)
}
