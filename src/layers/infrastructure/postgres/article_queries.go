package postgres

import (
	"github.com/UpBonent/news/src/common/models"
	"time"
)

func AddArticleQuery(article models.Article, author models.Author) (err error) {

	dateCreate := time.Now().Round(time.Minute)
	parseDatePublish, err := time.Parse("02.01.06 15:04", article.DatePublish)
	if err != nil {
		return
	}

	result := db.QueryRowx(CreateArticle, article.Header, article.Text, dateCreate, parseDatePublish, author.Name, author.Surname)

	return result.Err()
}

func AllArticleQuery() (articles []models.Article, err error) {
	var header, text, datePublish string
	var dateTimestamp time.Time
	var idAuthor int

	selector, err := db.Queryx(AllArticles)
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

func AllHeadersQuery() (articles []models.Article, err error) {

	selector, err := db.Queryx(AllHeaders)
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
