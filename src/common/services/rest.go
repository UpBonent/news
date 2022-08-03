package services

import "github.com/labstack/echo"

type RESTMethod interface {
	Register(c *echo.Echo)
}
