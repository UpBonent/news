package application

import (
	"context"
	"encoding/hex"
	"github.com/UpBonent/news/src/common/models"
	"github.com/labstack/echo"
)

func (a *Application) CreateNewAuthor(ctx context.Context, author models.Author, username, pwd string) (id int, err error) {
	id, err = a.Author.CheckAuthorExists(username, pwd)
	if err != nil || id == 0 {
		return 0, err
	}

	salt, err := getSalt()
	if err != nil {
		return 0, err
	}

	salt = append(salt, innerSalt...)

	hexPassword := hashing(pwd, salt, 8, 32)
	hexSalt := hex.EncodeToString(salt)

	id, err = a.Author.CreateNew(ctx, author, username, hexPassword, hexSalt)
	return
}

func (a *Application) GetAllAuthors(ctx context.Context) (authors []models.Author, err error) {
	return a.Author.GetAll(ctx)
}

func (a *Application) GetAuthorByID(ctx context.Context, id int) (author models.Author, err error) {
	return a.Author.GetByID(ctx, id)
}

func (a *Application) GetIDByAuthor(ctx context.Context, author models.Author) (id int, err error) {
	return a.Author.GetIDByName(ctx, author)
}

func (a *Application) CheckUserExist(username, password string, c echo.Context) (ok bool, err error) {
	id, err := a.Author.CheckAuthorExists(username, password)
	if err != nil {
		return false, err
	}

	if id == 0 {
		return false, nil
	}

	return true, nil
}
