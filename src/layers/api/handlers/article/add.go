package article

import (
	"encoding/json"
	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
	"github.com/labstack/echo"
	"io"
	"net/http"
)

func createArticle(c echo.Context) (err error) {
	var read []byte
	article := models.Article{}
	author := models.Author{}

	defer func() {
		err := c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	read, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(read, &article)
	if err != nil {
		return err
	}

	err = json.Unmarshal(read, &author)
	if err != nil {
		return err
	}

	err = postgres.AddArticleQuery(article, author)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "yeah, article has been created")
}

func example(c echo.Context) (err error) {
	q := `
{
	"header": "",
	"text": "",
	"date_publish": "02.8.22 15:04",
    "name": "Boris",
    "surname": "Pasternak"
}
`
	return c.String(http.StatusOK, q)
}
