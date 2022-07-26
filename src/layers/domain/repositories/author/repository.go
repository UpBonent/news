package author

import (
	"context"

	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/common/services"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	NewAuthor  = `INSERT INTO authors(name, surname, username, password, salt, cookie) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	AllAuthors = `SELECT id, name, surname FROM authors`

	GetAuthorByID   = `SELECT name, surname FROM authors WHERE id = $1`
	GetAuthorByName = `SELECT id FROM authors WHERE name = $1 AND surname = $2`

	CheckUserNameExisting = `SELECT id FROM authors WHERE username = $1`
	GetSalt               = `SELECT salt, password FROM authors WHERE username = $1`

	GetCookieValue    = `SELECT cookie FROM authors WHERE username = $1`
	GetAuthorByCookie = `SELECT id FROM authors WHERE cookie = $1`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) services.AuthorRepository {
	return &Repository{db}
}

func (r *Repository) CreateNew(ctx context.Context, a models.Author) (id int, err error) {
	result := r.db.QueryRowxContext(ctx, NewAuthor, a.Name, a.Surname, a.UserName, a.Password, a.Salt, a.CookieValue)
	if result.Err() != nil {
		return 0, result.Err()
	}
	err = result.Scan(&id)
	return
}

func (r *Repository) GetAll(ctx context.Context) (authors []models.Author, err error) {
	selector, err := r.db.QueryxContext(ctx, AllAuthors)
	if err != nil {
		return
	}

	for selector.Next() {
		var name, surname string
		var id int
		err = selector.Scan(&id, &name, &surname)
		if err != nil {
			return
		}
		nextAuthor := models.Author{
			Id:      id,
			Name:    name,
			Surname: surname,
		}
		authors = append(authors, nextAuthor)
	}
	return
}

func (r *Repository) GetByID(ctx context.Context, id int) (author models.Author, err error) {
	result := r.db.QueryRowxContext(ctx, GetAuthorByID, id)
	err = result.Scan(&author.Name, &author.Surname)
	return
}

func (r *Repository) GetIDByName(ctx context.Context, author models.Author) (id int, err error) {
	result := r.db.QueryRowContext(ctx, GetAuthorByName, author.Name, author.Surname)
	err = result.Scan(&id)
	return
}

func (r *Repository) CheckExisting(username string) (err error) {
	var id int

	result := r.db.QueryRowx(CheckUserNameExisting, username)
	err = result.Scan(&id)
	return
}

func (r *Repository) GetSalt(username string) (salt, passwordHash string, err error) {
	result := r.db.QueryRowx(GetSalt, username)
	err = result.Scan(&salt, &passwordHash)
	return
}

func (r *Repository) GetCookieValue(username string) (cookieValue string, err error) {
	result := r.db.QueryRowx(GetCookieValue, username)
	err = result.Scan(&cookieValue)
	return
}

func (r *Repository) GetAuthorByCookie(cookieValue string) (author models.Author, err error) {
	result := r.db.QueryRowx(GetAuthorByCookie, cookieValue)
	err = result.Scan(&author.Id)
	return
}
