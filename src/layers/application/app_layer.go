package application

import (
	"github.com/UpBonent/news/src/common/services"
	"github.com/UpBonent/news/src/layers/domain/repositories/article"
	"github.com/UpBonent/news/src/layers/domain/repositories/author"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Author  services.AuthorRepository
	Article services.ArticleRepository
	Logger  services.Logger
}

func SetApplicationLayer(db *sqlx.DB, logger *logrus.Entry) *Application {

	authorRep := author.NewRepository(db)
	articleRep := article.NewRepository(db)

	return &Application{
		Author:  authorRep,
		Article: articleRep,
		Logger:  logger.Logger,
	}
}
