package article

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/common/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"time"
)

const (
	create = `INSERT INTO articles(header, text, date_create, date_publish, id_authors) VALUES ($1, $2, $3, $4, $5)`
	all    = `SELECT id, header, text, date_publish, id_authors FROM articles`
	del    = `DELETE FROM articles WHERE id = $1`

	byAuthorID = `SELECT header, text, authors.name, authors.surname FROM articles INNER JOIN authors ON articles.id = authors.id;`

	updHeader  = `UPDATE articles SET header = $1 WHERE id = $2`
	updText    = `UPDATE articles SET text = $1 WHERE id = $2`
	updPublish = `UPDATE articles SET date_publish = $1 WHERE id = $2`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) services.ArticleRepository {
	return &Repository{db}
}

func (r *Repository) Insert(ctx context.Context, article models.Article, id int) (err error) {
	dateCreate := time.Now().Round(time.Minute)

	datePublish, err := time.Parse("02.01.06 15:04", article.DatePublish)
	if err != nil {
		return
	}

	_, err = r.db.ExecContext(ctx, create, article.Header, article.Text, dateCreate, datePublish, id)
	return
}

func (r *Repository) All(ctx context.Context) (articles []models.Article, err error) {
	var timestampPublish time.Time
	art := models.Article{}

	selector, err := r.db.QueryxContext(ctx, all)
	if err != nil {
		return
	}

	for selector.Next() {
		err = selector.Scan(&art.Id, &art.Header, &art.Text, &timestampPublish, &art.IdAuthor)
		if err != nil {
			return
		}

		nextArticle := models.Article{
			Id:          art.Id,
			Header:      art.Header,
			Text:        art.Text,
			DatePublish: timestampPublish.Format("02.01.06 15:04"),
			IdAuthor:    art.IdAuthor,
		}

		articles = append(articles, nextArticle)
	}
	return
}

func (r *Repository) Delete(ctx context.Context, id int) (err error) {

	_, err = r.db.ExecContext(ctx, del, id)

	return
}

func (r *Repository) UpDate(ctx context.Context, existArticle int, article models.Article) (err error) {
	count := 1

	if article.Header != "" {
		_, err = r.db.ExecContext(ctx, updHeader, article.Header, existArticle)
		if err != nil {
			return
		}
		count--
	}

	if article.Text != "" {
		_, err = r.db.ExecContext(ctx, updText, article.Text, existArticle)
		if err != nil {
			return
		}
		count--
	}

	if article.DatePublish != "" {
		var parseDatePublish time.Time
		parseDatePublish, err = time.Parse("02.01.06 15:04", article.DatePublish)
		if err != nil {
			return
		}

		_, err = r.db.ExecContext(ctx, updPublish, parseDatePublish, existArticle)
		if err != nil {
			return
		}
		count--
	}

	if count > 0 {
		return errors.New("Nothing to change")
	}

	return
}
