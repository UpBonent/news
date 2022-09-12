package rest

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
	wayToDeleteAuthor     = "/delete"
	wayToArticlesByAuthor = "/"
)

type HandlerAuthor struct {
	ctx        context.Context
	way        string
	repository services.AuthorRepository
}

func NewHandlerAuthor(ctx context.Context, s string, rep services.AuthorRepository) services.AuthorHandler {
	return &HandlerAuthor{ctx, s, rep}
}

func (h *HandlerAuthor) Register(e *echo.Echo) {
	g := e.Group(h.way)
	g.GET("", h.All)
	g.POST(wayToCreateAuthor, h.Create)
	g.DELETE(wayToDeleteAuthor, h.Delete)
}

func (h *HandlerAuthor) All(c echo.Context) error {
	authors, err := h.repository.All(h.ctx)
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

func (h *HandlerAuthor) Create(c echo.Context) (err error) {
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
		return err
	}

	err = json.Unmarshal(read, &author)
	if err != nil {
		return err
	}

	err = h.repository.Insert(h.ctx, author)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "yeah, author has been created")
}

func (h *HandlerAuthor) Delete(c echo.Context) (err error) {
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
		return err
	}

	err = json.Unmarshal(read, &author)
	if err != nil {
		return err
	}

	err = h.repository.Delete(h.ctx, author.Id)
	if err != nil {
		return err
	}
	return c.String(http.StatusResetContent, "yeah, author has been deleted")
}

//
//func (h HandlerAuthor) articlesByAuthor(c echo.Context) (err error) {
//
//	return c.String(http.StatusOK, "There are articles by this author")
//}
