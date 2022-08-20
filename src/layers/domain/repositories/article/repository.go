package article

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/common/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

const (
	create = `INSERT INTO articles(header, text, date_create, date_publish, id_authors) VALUES ($1, $2, $3, $4, (SELECT id FROM authors WHERE name = $5 AND surname = $6))`
	all    = `SELECT header, text, date_publish, id_authors FROM articles`
	del    = `DELETE FROM articles WHERE header = $1`

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

func (r *Repository) Insert(ctx context.Context, article models.Article, author models.Author) (err error) {
	dateCreate := time.Now().Round(time.Minute)

	parseDatePublish, err := time.Parse("02.01.06 15:04", article.DatePublish)
	if err != nil {
		return
	}

	_, err = r.db.ExecContext(ctx, create, article.Header, article.Text, dateCreate, parseDatePublish, author.Name, author.Surname)
	err = r.db.Beginx()
	return
}

func (r *Repository) All(ctx context.Context) (articles []models.Article, err error) {
	var header, text, datePublish string
	var dateTimestamp time.Time
	var idAuthor int

	selector, err := r.db.QueryxContext(ctx, all)
	if err != nil {
		return
	}

	for selector.Next() {
		err = selector.Scan(&header, &text, &dateTimestamp, &idAuthor)
		if err != nil {
			return
		}

		datePublish = time.Time.String(dateTimestamp)

		nextArticle := models.Article{
			Header:      header,
			Text:        text,
			DatePublish: datePublish,
			IdAuthor:    idAuthor,
		}

		articles = append(articles, nextArticle)
	}
	return
}

func (r *Repository) Delete(ctx context.Context, article models.Article) (err error) {

	_, err = r.db.ExecContext(ctx, del, article.Header)

	return
}

func (r *Repository) UpDate(ctx context.Context, existArticle int, article models.Article) (err error) {

	if article.Header != "" {
		_, err = r.db.ExecContext(ctx, updHeader, article.Header, existArticle)
		if err != nil {
			return
		}
	}

	if article.Text != "" {
		_, err = r.db.ExecContext(ctx, updText, article.Text, existArticle)
		if err != nil {
			return
		}
	}

	if article.DatePublish != "" {
		parseDatePublish, err := time.Parse("02.01.06 15:04", article.DatePublish)
		if err != nil {
			return
		}

		_, err = r.db.ExecContext(ctx, updPublish, parseDatePublish, existArticle)
		if err != nil {
			return
		}
	}

	return
}
