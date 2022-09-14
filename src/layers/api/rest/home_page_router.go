package rest

import (
	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
	"net/http"
)

type homePageStruct struct {
	way string
}

func NewHandlerHomePage(way string) services.Handler {
	return &homePageStruct{way: way}
}

func (h *homePageStruct) Register(c *echo.Echo) {
	c.GET(h.way, h.homePage)
}

func (h *homePageStruct) homePage(c echo.Context) (err error) {
	hint := `
	Welcome! There's a news portal.
		You can use these options:
			- /authors -- to view All author in our portal;
			- /authors/create -- to create a new author.

			- /articles -- to view All article from the latest to the earliest;
			- /articles/headers -- to light view article with header only;
			- /articles/create -- to create new article (use GET method to hint).
`
	return c.String(http.StatusOK, hint)
}
