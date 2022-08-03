package article

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return ArticleRepository
}

func (r *Repository) Insert(ctx context.Context) error {
	parseDateCreate, err := pq.ParseTimestamp(time.Local, article.DateCreate)
	if err != nil {
		return
	}
	parseDatePublish, err := pq.ParseTimestamp(time.Local, article.DatePublish)
	if err != nil {
		return
	}

	result := r.db.QueryRowx(CreateArticle, article.Header, article.Text, parseDateCreate, parseDatePublish, author.Name, author.Surname)

	return result.Err()
}

func AddArticleQuery(article models.Article, author models.Author) (err error) {

}
