package main

import (
	"context"
	"github.com/UpBonent/news/src/layers/api/rest"
	"github.com/UpBonent/news/src/layers/domain/repositories"
	"github.com/UpBonent/news/src/layers/infrastructure/url"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/UpBonent/news/src/layers/infrastructure/config"
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
)

func main() {
	var err error
	var cfg config.Config

	ctx := context.Background()
	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile(cfg.LogsFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	_ = logging.NewLogger(logFile)

	db, err := postgres.NewClient(ctx, cfg.Storage)
	if err != nil {
		panic(err)
	}

	rep := repositories.NewRepository(db)

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}. Error: [${error}]` + "\n"}))
	e.Use(middleware.Recover())
	// Routes
	ways := url.Ways{
		Home: "/hint", Author: "/authors", Article: "/articles",
	}

	service := rest.NewService(ctx, ways, rep)

	service.HomePageHandler.Register(e)
	service.AuthorHandler.Register(e)
	service.ArticleHandler.Register(e)
	//rest.NewHandlerHomePage("/hint").Register(e)
	//rest.NewHandlerAuthor(ctx, "/authors", rep.AuthorRepository).Register(e)
	//rest.NewHandlerArticle(ctx, "/articles", rep).Register(e)
	// Start server
	e.Logger.Fatal(e.Start(cfg.Listen.Port))
}
