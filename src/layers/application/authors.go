package application

import (
	"context"
	"database/sql"
	"encoding/hex"
	"github.com/UpBonent/news/src/common/models"
	"github.com/pkg/errors"
	"net/http"
)

func (a *Application) CreateNewAuthor(ctx context.Context, author models.Author, checkPWD string) (id int, err error) {
	if author.Password != checkPWD {
		return 0, errors.New("Passwords are different")
	}

	err = a.Author.CheckExisting(author.UserName)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if err == nil {
		return 0, errors.New("User already exists with the same username")
	}

	s, err := generate(salt)
	if err != nil {
		return 0, err
	}
	c, err := generate(cookieValue)
	if err != nil {
		return 0, err
	}

	author.Password = hashing(author.Password, s)
	author.Salt = hex.EncodeToString(s)

	id, err = a.Author.CreateNew(ctx, author)

	return a.Author.
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

	return errors.New("Wrong username and/or password")
}

func (a *Application) SetUserCookie(userName string) (newCookie http.Cookie, err error) {

if userName	cake, err := generate(cookieValue)
	if err != nil {
		return
	}

	userCookie := hex.EncodeToString(cake)

	newCookie = http.Cookie{
		Name:   "i",
		Value:  userCookie,
		Path:   "/",
		Domain: "localhost:8080",
		MaxAge: 86400,
	}
	return
}

func (a *Application) GetAuthorByCookie(c string) (author models.Author) {

}