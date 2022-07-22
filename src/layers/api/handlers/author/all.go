package author

import (
	"github.com/labstack/echo"
	"net/http"
)

func allAuthor(c echo.Context) error {
	return c.String(http.StatusOK, "you got it")
}
