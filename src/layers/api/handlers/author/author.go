package author

import (
	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
)

//can DELETE
var _ services.RESTMethod = &handlerAuthor{}

const (
	toCreateAuthor = "/create"
)

type handlerAuthor struct {
}

func NewHandlerAuthor() services.RESTMethod {
	return &handlerAuthor{}
}

func (h *handlerAuthor) Register(way string, e *echo.Echo) {
	g := e.Group(way)
	g.GET("", allAuthor)
	g.POST(toCreateAuthor, createAuthor)
}
