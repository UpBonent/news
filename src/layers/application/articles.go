package application

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
)

func (a *Application) CreateNewArticle(ctx context.Context, article models.Article, id int) error {
	//TODO implement me
	panic("implement me")
}

func (a *Application) GetAllArticles(ctx context.Context) (articles []models.Article, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) UpdateArticle(ctx context.Context, existArticle int, article models.Article) error {
	//TODO implement me
	panic("implement me")
}

func (a *Application) GetArticlesByAuthorID(ctx context.Context, id int) (articles []models.Article, err error) {
	//TODO implement me
	panic("implement me")
}
