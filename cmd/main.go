package main

import (
	"context"
	"github.com/UpBonent/news/src/layers/application"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"

	"github.com/UpBonent/news/src/layers/api/rest"
	"github.com/UpBonent/news/src/layers/infrastructure/config"
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
)

func main() {
	var err error
	var cfg config.Config

	ctx := context.Background()
	err = cleanenv.ReadConfig("./config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	//logFile, err := os.OpenFile(cfg.PathToFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	//if err != nil {
	//	panic(err)
	//}
	logger := logging.NewLogger(os.Stdout)

	dataBaseConnection, err := postgres.NewClient(ctx, cfg.Storage)
	if err != nil {
		panic(err)
	}

	app := application.SetApplicationLayer(dataBaseConnection, logger)

	e := echo.New()
	e.Use(middleware.Static("./static"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}. Error: [${error}]` + "\n"}))
	e.Use(middleware.Recover())

	rest.NewHandlersAuthor(ctx, app).Register(e)
	rest.NewHandlersArticle(ctx, app).Register(e)

	app.Logger.Info("server has been started")
	e.Logger.Fatal(e.Start(cfg.Listen.Port))
}
