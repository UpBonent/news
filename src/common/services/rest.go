package services

import (
	"github.com/labstack/echo"
)

type Handler interface {
	Register(c *echo.Echo)
}
