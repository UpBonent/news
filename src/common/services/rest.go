package services

import "github.com/labstack/echo"

type RESTMethod interface {
	Register(way string, c *echo.Echo)
}
