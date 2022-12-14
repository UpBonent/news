package article

import (
	"context"
	"time"

	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/common/services"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	create = `INSERT INTO articles(header, text, annotation, date_create, date_publish, author_id) VALUES ($1, $2, $3, $4, $5, $6)`
	all    = `SELECT id, header, text, date_publish, author_id FROM articles`

	byAuthorID = `SELECT id, header, text, date_publish FROM articles WHERE author_id = $1`

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

func (r *Repository) CreateNew(ctx context.Context, article models.Article, dateCreate time.Time, id int) (err error) {
	_, err = r.db.ExecContext(ctx, create, article.Header, article.Text, article.Annotation, dateCreate, article.DatePublish, id)
	return
}

func (r *Repository) GetAll(ctx context.Context) (articles []models.Article, err error) {
	art := models.Article{}

	selector, err := r.db.QueryxContext(ctx, all)
	if err != nil {
		return
	}

	for selector.Next() {
		err = selector.Scan(&art.Id, &art.Header, &art.Annotation, &art.DatePublish, &art.AuthorID)
		if err != nil {
			return
		}

		nextArticle := models.Article{
			Id:          art.Id,
			Header:      art.Header,
			Annotation:  art.Annotation,
			DatePublish: art.DatePublish,
			AuthorID:    art.AuthorID,
		}

		articles = append(articles, nextArticle)
	}
	return
}

func (r *Repository) Update(ctx context.Context, existArticle int, newArticle models.Article) (err error) {
	count := 1

	if newArticle.Header != "" {
		_, err = r.db.ExecContext(ctx, updHeader, newArticle.Header, existArticle)
		if err != nil {
			return
		}
		count--
	}

	if newArticle.Text != "" {
		_, err = r.db.ExecContext(ctx, updText, newArticle.Text, existArticle)
		if err != nil {
			return
		}
		count--
	}

	_, err = r.db.ExecContext(ctx, updPublish, newArticle.DatePublish, existArticle)
	if err != nil {
		return
	}

	if count > 0 {
		return errors.New("Nothing to change")
	}

	return
}

func (r *Repository) GetByAuthorID(ctx context.Context, id int) (articles []models.Article, err error) {
	var timestampPublish time.Time
	article := models.Article{}

	selector, err := r.db.QueryxContext(ctx, byAuthorID, id)
	if err != nil {
		return
	}

	for selector.Next() {
		err = selector.Scan(&article.Id, &article.Header, &article.Text, &timestampPublish)
		if err != nil {
			return
		}

		articles = append(articles, article)
	}
	return
}
