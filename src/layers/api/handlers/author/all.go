package author

import (
	"encoding/json"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
	"github.com/labstack/echo"
	"net/http"
)

func allAuthors(c echo.Context) error {
	authors, err := postgres.AllAuthorsQuery()
	if err != nil {
		return err
	}

	var newLine = byte(10)

	for _, author := range authors {
		res, err := json.Marshal(author)
		res = append(res, newLine)
		_, err = c.Response().Write(res)
		if err != nil {
			return err
		}

	}

	return c.String(http.StatusOK, "There are authors")
}
