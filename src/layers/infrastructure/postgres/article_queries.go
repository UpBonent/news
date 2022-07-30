package postgres

import (
	"github.com/UpBonent/news/src/common/models"
	"github.com/lib/pq"
	"time"
)

func AddArticleQuery(article models.Article, author models.Author) (err error) {

	parseDateCreate, err := pq.ParseTimestamp(time.Local, article.DateCreate)
	if err != nil {
		return
	}
	parseDatePublish, err := pq.ParseTimestamp(time.Local, article.DatePublish)
	if err != nil {
		return
	}

	result := db.QueryRowx(CreateArticle, article.Header, article.Text, parseDateCreate, parseDatePublish, author.Name, author.Surname)

	return result.Err()
}
