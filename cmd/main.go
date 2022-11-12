package main

import (
	"context"

	"github.com/UpBonent/news/src/layers/application"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/UpBonent/news/src/layers/api/rest"
	"github.com/UpBonent/news/src/layers/infrastructure/config"
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
)

func main() {
	var err error
	var cfg config.Config

	ctx := context.Background()
	err = cleanenv.ReadConfig("../config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	logOutput := logging.SetLoggerOutput(cfg.Log.Output, cfg.Log.PathToFile)
	logger := logging.NewLogger(cfg.Log.ActiveLevels, logOutput)

	dataBaseConnection, err := postgres.NewClient(ctx, cfg.Storage)
	if err != nil {
		panic(err)
	}

	app := application.SetApplicationLayer(dataBaseConnection, logger)

	e := echo.New()
	e.Use(middleware.Static("../static"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}. Error: [${error}]` + "\n"}))
	e.Use(middleware.Recover())

	rest.NewHandlersAuthor(ctx, app).Register(e)
	rest.NewHandlersArticle(ctx, app).Register(e)

	logger.INFO("Server has started")
	e.Logger.Fatal(e.Start(cfg.Listen.Port))
}
