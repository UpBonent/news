package author

import (
	"github.com/labstack/echo"
	"net/http"
)

func createAuthor(c echo.Context) error {
	return c.String(http.StatusCreated, "yeah, author has been created")
}
