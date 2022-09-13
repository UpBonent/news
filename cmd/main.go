package main

import (
	"context"
	"github.com/UpBonent/news/src/layers/api/rest"
	"github.com/UpBonent/news/src/layers/domain/repositories/article"
	"github.com/UpBonent/news/src/layers/domain/repositories/author"
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

	authorRepository := author.NewRepository(db)
	articleRepository := article.NewRepository(db)

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}. Error: [${error}]` + "\n"}))
	e.Use(middleware.Recover())
	// Routes
	rest.NewHandlerHomePage("/").Register(e)
	rest.NewHandlerAuthor(ctx, "/authors", authorRepository).Register(e)
	rest.NewHandlerArticle(ctx, "/articles", articleRepository, authorRepository).Register(e)

	// Start server
	e.Logger.Fatal(e.Start(cfg.Listen.Port))
}
