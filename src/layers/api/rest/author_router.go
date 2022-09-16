package rest

//Returning errors into handler:
//- send it to user?
//- or log it in a server side?

import (
	"context"
	"encoding/json"
	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
	"io"
	"net/http"
)

const (
	wayToCreateAuthor     = "/create"
	wayToArticlesByAuthor = "/articles"
)

type handlerAuthor struct {
	ctx               context.Context
	way               string
	articleRepository services.ArticleRepository
	authorRepository  services.AuthorRepository
}

func NewHandlerAuthor(ctx context.Context, s string, article services.ArticleRepository, author services.AuthorRepository) services.Handler {
	return &handlerArticle{ctx, s, article, author}
}

func (h *handlerAuthor) Register(e *echo.Echo) {
	g := e.Group(h.way)
	g.GET("", h.all)
	g.POST(wayToCreateAuthor, h.create)
	g.GET(wayToArticlesByAuthor, h.articlesByAuthor)
}

func (h *handlerAuthor) all(c echo.Context) error {
	authors, err := h.authorRepository.All(h.ctx)
	if err != nil {
		return err
	}

	var newLine = byte(10)

	for _, author := range authors {
		res, err := json.Marshal(author)
		res = append(res, newLine)
		_, err = c.Response().Write(res)
		if err != nil {
			return err
		}
	}

	return c.String(http.StatusOK, "There are authors")
}

func (h *handlerAuthor) create(c echo.Context) (err error) {
	var read []byte
	authorJSON := AuthorJSON{}

	read, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(read, &authorJSON)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	author := models.Author{
		Id:      authorJSON.Id,
		Name:    authorJSON.Name,
		Surname: authorJSON.Surname,
	}

	err = h.authorRepository.Insert(h.ctx, author)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = c.Request().Body.Close()
	if err != nil {
		return
	}

	return c.String(http.StatusCreated, "yeah, author has been created")

}

func (h *handlerAuthor) articlesByAuthor(c echo.Context) (err error) {
	var read []byte
	author := models.Author{}

	defer func() {
		err := c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	read, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(read, &author)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.authorRepository.Insert(h.ctx, author)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "There are articles by this author")
}
