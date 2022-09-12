package services

import (
	"github.com/labstack/echo"
)

type HomePageHandler interface {
	Register(c *echo.Echo)
	HomePage(c echo.Context) error
}

type AuthorHandler interface {
	Register(c *echo.Echo)
	All(c echo.Context) error
	Create(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
}

type ArticleHandler interface {
	Register(c *echo.Echo)
	All(c echo.Context) error
	Create(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
	Update(c echo.Context) (err error)
}
