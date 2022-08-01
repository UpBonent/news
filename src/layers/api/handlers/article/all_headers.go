package article

import (
	"github.com/labstack/echo"
	"net/http"
)

func allHeaders(c echo.Context) error {

	return c.String(http.StatusOK, "There are all headers")
}
