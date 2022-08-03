package main

import (
	"context"
	"github.com/UpBonent/news/src/layers/api/handlers/author"
	"github.com/UpBonent/news/src/layers/domain/repositories/article"
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

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

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
