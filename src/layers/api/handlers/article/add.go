package article

import (
	"github.com/labstack/echo"
	"net/http"
)

func createArticle(c echo.Context) error {
	return c.String(http.StatusCreated, "yeah, article has been created")
}
