package article

import (
	"encoding/json"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
	"github.com/labstack/echo"
	"net/http"
)

func allArticle(c echo.Context) error {
	articles, err := postgres.AllArticleQuery()
	if err != nil {
		return err
	}

	var newLine = byte(10)

	for _, article := range articles {
		author, err := postgres.GetAuthorByIDQuery(article.IdAuthor)
		if err != nil {
			return err
		}

		articleJSON, err := json.Marshal(article)
		if err != nil {
			return err
		}

		authorJSON, err := json.Marshal(author)
		if err != nil {
			return err
		}

		articleJSON = append(articleJSON, authorJSON...)
		articleJSON = append(articleJSON, newLine)

		_, err = c.Response().Write(articleJSON)
		if err != nil {
			return err
		}
	}

	return c.String(http.StatusOK, "There are all articles")
}
