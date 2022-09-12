package author

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/common/services"
	"github.com/jmoiron/sqlx"
)

const (
	NewAuthor    = `INSERT INTO authors(name, surname) VALUES($1, $2)`
	DeleteAuthor = `DELETE FROM authors WHERE id = $1`
	AllAuthors   = `SELECT id, name, surname FROM authors`

	GetAuthorByID   = `SELECT name, surname FROM authors WHERE id = $1`
	GetAuthorByName = `SELECT id FROM authors WHERE name = $1 AND surname = $2`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) services.AuthorRepository {
	return &Repository{db}
}

func (r *Repository) Insert(ctx context.Context, author models.Author) (err error) {
	_, err = r.db.ExecContext(ctx, NewAuthor, author.Name, author.Surname)
	return
}

func (r *Repository) All(ctx context.Context) (authors []models.Author, err error) {
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

func (r *Repository) Delete(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, DeleteAuthor, id)
	return
}

func (r *Repository) GetByID(ctx context.Context, id int) (author models.Author, err error) {
	a := models.Author{}
	result := r.db.QueryRowxContext(ctx, GetAuthorByID, id)
	err = result.Scan(&a.Name, &a.Surname)
	return
}

func (r *Repository) GetByName(ctx context.Context, author models.Author) (id int, err error) {
	result := r.db.QueryRowContext(ctx, GetAuthorByName, author.Name, author.Surname)
	err = result.Scan(&id)
	return
}
