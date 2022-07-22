package article

import (
	"github.com/labstack/echo"
	"net/http"
)

func allArticle(c echo.Context) error {
	return c.String(http.StatusOK, "you got it")
}
