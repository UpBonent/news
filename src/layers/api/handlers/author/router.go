package author

import (
	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
)

//can DELETE
var _ services.RESTMethod = &Author{}

const (
	toCreateAuthor = "/create"
)

type Author struct {
}

func NewHandlerAuthor() services.RESTMethod {
	return &Author{}
}

func (a *Author) Register(way string, e *echo.Echo) {
	g := e.Group(way)
	g.GET("", allAuthors)
	g.POST(toCreateAuthor, createAuthor)
}
