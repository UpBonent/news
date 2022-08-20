package rest

import (
	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
	"net/http"
)

type homePageStruct struct {
	way string
}

func NewHandlerHomePage(way string) services.RESTMethod {
	return &homePageStruct{way: way}
}

func (h *homePageStruct) Register(c *echo.Echo) {
	c.GET(h.way, h.homePage)
}

func (h *homePageStruct) homePage(c echo.Context) (err error) {
	hint := `
	Welcome! There's a news portal.
		You can use these options:
			- /authors -- to view all author in our portal;
			- /authors/create -- to create a new author;
			- /author/delete -- to delete exists author.

			- /article -- to view all article from the latest to the earliest;
			- /article/headers -- to light view article with header only;
			- /article/create -- to create new article (use GET method to hint).
`
	return c.String(http.StatusOK, hint)
}
