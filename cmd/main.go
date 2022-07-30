package main

import (
	"context"
	"github.com/UpBonent/news/src/layers/api"
	"github.com/UpBonent/news/src/layers/domain"
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
)

func main() {
	logging.NewLogger()
	l := logging.GetLogger()
	cfg := domain.GetConfig(l)
	ctx := context.Background()

	_ = postgres.NewClient(ctx, 5, cfg.Storage, l)

	api.StarServer()

}

//
//
//
//
//
//
//

//
//g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
//	if username == "Tom" && password == "1234" {
//		return true, nil
//	}
//	return false, echo.NewHTTPError(http.StatusNotAcceptable, "pleas write incorrect name & password")
//}))
