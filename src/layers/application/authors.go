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

	ok, err := a.CheckUserExisting(author.UserName)
	if ok == true {
		return 0, errors.New("User already exists with the same username")
	}
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	s, err := generate(salt)
	if err != nil {
		return 0, err
	}
	c, err := generate(cookie)
	if err != nil {
		return 0, err
	}

	author.Password = hashing(author.Password, s)
	author.Salt = hex.EncodeToString(s)

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

func (a *Application) CheckUserExisting(username string) (bool, error) { //delete this intermediate stage
	return a.Author.CheckExisting(username)
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

func (a *Application) SetUserCookie() (c http.Cookie) {

	userCookie := hex.EncodeToString(c)

	newCookie := http.Cookie{
		Name:   "i",
		Value:  ,
		Path:   "/",
		Domain: "localhost:8080",
		MaxAge: 86400,
	}

}