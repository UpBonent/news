package main

import (
	"github.com/UpBonent/news/src/layers/api"
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
)

type Cat struct {
	Name  string `json:"name"`
	Years string `json:"years"`
}

func main() {
	logging.NewLogger()
	l := logging.GetLogger()
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
