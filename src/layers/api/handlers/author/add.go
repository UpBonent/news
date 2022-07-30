package author

import (
	"encoding/json"
	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
	"github.com/labstack/echo"
	"io"
	"net/http"
)

func createAuthor(c echo.Context) (err error) {
	var read []byte
	newAuthor := models.Author{}

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

	err = json.Unmarshal(read, &newAuthor)
	if err != nil {
		return err
	}

	err = postgres.AddAuthorQuery(newAuthor)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "yeah, author has been created")
}
