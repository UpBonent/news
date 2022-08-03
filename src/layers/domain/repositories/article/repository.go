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
	CreateArticleQuery = `INSERT INTO articles(header, text, date_create, date_publish, id_authors) VALUES ($1, $2, $3, $4, (SELECT id FROM authors WHERE name = $5 AND surname = $6))`
	AllArticlesQuery   = `SELECT header, text, date_publish, id_authors FROM articles`
	AllHeadersQuery    = `SELECT header, date_publish FROM articles`
	HeadersByTimeQuery = `SELECT header, date_publish FROM articles WHERE date_publish < $1 AND articles.date_publish > $2`

	ArticlesByAuthorQuery = `SELECT header, text, authors.name, authors.surname FROM articles INNER JOIN authors ON articles.id = authors.id;`
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

	result := r.db.QueryRowxContext(ctx, CreateArticleQuery, article.Header, article.Text, dateCreate, parseDatePublish, author.Name, author.Surname)
	return result.Err()
}

func (r *Repository) All(ctx context.Context) (articles []models.Article, err error) {
	var header, text, datePublish string
	var dateTimestamp time.Time
	var idAuthor int

	selector, err := r.db.QueryxContext(ctx, AllArticlesQuery)
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

func (r *Repository) AllHeaders(ctx context.Context) (articles []models.Article, err error) {
	selector, err := r.db.QueryxContext(ctx, AllHeadersQuery)
	if err != nil {
		return
	}

	for selector.Next() {
		var header string
		var dateTimestamp time.Time

		err = selector.Scan(&header, &dateTimestamp)
		if err != nil {
			return
		}

		datePublish := time.Time.String(dateTimestamp)

		nextArticle := models.Article{
			Header:      header,
			DatePublish: datePublish,
		}

		articles = append(articles, nextArticle)
	}
	return
}

func (r *Repository) HeadersByTime(ctx context.Context) (articles []models.Article, err error) {
	selector, err := r.db.QueryxContext(ctx, HeadersByTimeQuery)
	if err != nil {
		return
	}

	for selector.Next() {
		var header string
		var dateTimestamp time.Time

		err = selector.Scan(&header, &dateTimestamp)
		if err != nil {
			return
		}

		datePublish := time.Time.String(dateTimestamp)

		nextArticle := models.Article{
			Header:      header,
			DatePublish: datePublish,
		}

		articles = append(articles, nextArticle)
	}
	return
}
