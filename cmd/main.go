package main

import (
	"context"
	"github.com/UpBonent/news/src/layers/domain/repositories/article"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/UpBonent/news/src/layers/api"
	"github.com/UpBonent/news/src/layers/infrastructure/config"
	"github.com/UpBonent/news/src/layers/infrastructure/logging"
	"github.com/UpBonent/news/src/layers/infrastructure/postgres"
)

func main() {
	var err error
	var cfg *config.Config

	err = cleanenv.ReadConfig("config.yml", cfg)
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	logger := logging.NewLogger(logFile)

	ctx := context.Background()
	db := postgres.NewClient(ctx, 5, cfg.Storage, logger)

	aRepository := article.NewRepository(db)

	api.StarServer()

}

//	var stru struct {
//		Date  string `json:"date"`
//		Check string `json:"check"`
//	}
//
//	str := `{"date": "` + string(pq.FormatTimestamp(time.Now())) + `", "check": "work"}`
//	_ = json.Unmarshal([]byte(str), &stru)
//	fmt.Println(stru)
//
//	stroka := string(pq.FormatTimestamp(time.Now()))
//	parse, _ := pq.ParseTimestamp(time.Local, stroka)
//	fmt.Printf("value: %v, \n Type: %T", parse, parse)
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
