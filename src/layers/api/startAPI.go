package api

import (
	"github.com/UpBonent/news/src/layers/api/handlers/article"
	"github.com/UpBonent/news/src/layers/api/handlers/author"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StarServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}. Error: [${error}]` + "\n",
	}))
	e.Use(middleware.Recover())

	// Routes
	routeAuthor := author.NewHandlerAuthor()
	routeAuthor.Register("/author", e)

	routeArticle := article.NewHandlerArticle()
	routeArticle.Register("/article", e)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../html",
	}))
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
