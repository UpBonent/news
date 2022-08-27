package services

import "github.com/labstack/echo"

//go:generate mockgen -source=rest.go -destination=mocks/mock.go

type RESTMethod interface {
	Register(c *echo.Echo)
}
