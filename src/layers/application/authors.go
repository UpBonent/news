package application

import (
	"context"
	"encoding/hex"
	"github.com/UpBonent/news/src/common/models"
	"github.com/pkg/errors"
)

func (a *Application) CreateNewAuthor(ctx context.Context, author models.Author) (id int, err error) {
	ok, err := a.CheckUserExisting(author.UserName)
	if err != nil || id == 0 {
		return 0, err
	}
	if ok == true {
		return 0, errors.New("The author already exists")
	}

	salt, err := generateSalt()
	if err != nil {
		return 0, err
	}

	author.Password = hashing(author.Password, salt)
	author.Salt = hex.EncodeToString(salt)

	id, err = a.Author.CreateNew(ctx, author)
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

func (a *Application) CheckUserExisting(username string) (bool, error) {
	return a.Author.CheckExisting(username)
}

func (a *Application) CheckUserAuthentication(username, password string) (ok bool, err error) {
	salt, existedPassword, err := a.Author.GetSalt(username)
	if err != nil {
		return false, err
	}

	byteSalt, err := hex.DecodeString(salt)
	if err != nil {
		return false, err
	}

	inputPassword := hashing(password, byteSalt)

	if inputPassword == existedPassword {
		return true, nil
	}

	return false, errors.New("Wrong username and/or password")
}
