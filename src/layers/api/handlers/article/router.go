package article

import (
	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
)

//can DELETE
var _ services.RESTMethod = &handlerArticle{}

const (
	toCreateArticle = "/create"
)

type handlerArticle struct {
}

func NewHandlerArticle() services.RESTMethod {
	return &handlerArticle{}
}

func (h *handlerArticle) Register(way string, e *echo.Echo) {
	e.GET(way, allArticle)
	g := e.Group(way)
	g.POST(toCreateArticle, createArticle)
}
