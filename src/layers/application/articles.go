package application

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"time"
)

func (a *Application) CreateNewArticle(ctx context.Context, article models.Article, id int) (err error) {
	dateCreate := time.Now().Round(time.Minute)
	err = a.Article.CreateNew(ctx, article, dateCreate, id)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) GetAllArticles(ctx context.Context) (articles []models.Article, err error) {
	articles, err = a.Article.GetAll(ctx)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) UpdateArticle(ctx context.Context, existArticle int, article models.Article) (err error) {
	err = a.Article.Update(ctx, existArticle, article)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) GetArticlesByAuthorID(ctx context.Context, id int) (articles []models.Article, err error) {
	articles, err = a.Article.GetByAuthorID(ctx, id)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}
