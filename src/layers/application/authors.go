package application

import (
	"context"
	"database/sql"
	"encoding/hex"
	"github.com/UpBonent/news/src/common/models"
	"github.com/pkg/errors"
	"net/http"
)

func (a *Application) CreateNewAuthor(ctx context.Context, author models.Author) (id int, c string, err error) {
	err = a.Author.CheckExisting(author.UserName)
	if err != nil && err != sql.ErrNoRows {
		a.Logger.Error(err.Error())
		return 0, "", err
	}
	if err == nil {
		return 0, "", errors.New("User already exists with the same username")
	}

	s, err := generate(salt)
	if err != nil {
		a.Logger.Error(err.Error())
		return 0, "", err
	}
	cookieValue, err := generate(cookieValue)
	if err != nil {
		a.Logger.Error(err.Error())
		return 0, "", err
	}

	author.Password = hashing(author.Password, s)
	author.Salt = hex.EncodeToString(s)
	author.CookieValue = hex.EncodeToString(cookieValue)

	id, err = a.Author.CreateNew(ctx, author)
	if err != nil {
		a.Logger.Error(err.Error())
	}

	return
}

func (a *Application) GetAllAuthors(ctx context.Context) (authors []models.Author, err error) {
	authors, err = a.Author.GetAll(ctx)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) GetAuthorByID(ctx context.Context, id int) (author models.Author, err error) {
	author, err = a.Author.GetByID(ctx, id)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) GetIDByAuthor(ctx context.Context, author models.Author) (id int, err error) {
	id, err = a.Author.GetIDByName(ctx, author)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) CheckUserAuthentication(username, password string) (err error) {
	salt, existedPassword, err := a.Author.GetSalt(username)
	if err != nil {
		return err
	}

	byteSalt, err := hex.DecodeString(salt)
	if err != nil {
		return err
	}

	inputPassword := hashing(password, byteSalt)

	if inputPassword == existedPassword {
		return nil
	}

	err = errors.New("Wrong username and/or password")
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) SetUserCookie(cookieValue string) (newCookie http.Cookie) {
	newCookie = http.Cookie{
		Name:   "i",
		Value:  cookieValue,
		Path:   "/",
		Domain: "localhost:8080",
		MaxAge: 86400,
	}
	return
}

func (a *Application) GetAuthorByCookie(cookieValue string) (author models.Author, err error) {
	author, err = a.Author.GetAuthorByCookie(cookieValue)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}

func (a *Application) GetCookieByUserName(username string) (cookieValue string, err error) {
	cookieValue, err = a.Author.GetCookieValue(username)
	if err != nil {
		a.Logger.Error(err.Error())
	}
	return
}
