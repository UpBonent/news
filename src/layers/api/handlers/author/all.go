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

	var jsonSingle string

	for _, author := range authors {
		jsonStruct, err := json.Marshal(author)
		if err != nil {
			return err
		}
		jsonSingle += string(jsonStruct) + "\n"
	}

	return c.JSON(http.StatusOK, jsonSingle)
}
