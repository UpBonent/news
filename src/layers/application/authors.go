package application

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
)

func (a *Application) CreateNewAuthor(ctx context.Context, author models.Author) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) GetAllAuthors(ctx context.Context) (authors []models.Author, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) GetAuthorByID(ctx context.Context, id int) (author models.Author, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) GetIDByAuthor(ctx context.Context, author models.Author) (id int, err error) {
	//TODO implement me
	panic("implement me")
}
