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

//can DELETE
var _ services.RESTMethod = &handlerAuthor{}

const (
	toCreateAuthor = "/create"
	toDeleteAuthor = "/delete"
)

type handlerAuthor struct {
	ctx        context.Context
	way        string
	repository services.AuthorRepository
}

func NewHandlerAuthor(ctx context.Context, s string, r services.AuthorRepository) services.RESTMethod {
	return &handlerAuthor{ctx, s, r}
}

func (h *handlerAuthor) Register(e *echo.Echo) {
	g := e.Group(h.way)
	g.GET("", h.all)
	g.POST(toCreateAuthor, h.insert)
	g.DELETE(toDeleteAuthor, h.delete)
}

func (h *handlerAuthor) insert(c echo.Context) (err error) {
	var read []byte
	auth := models.Author{}

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

	err = json.Unmarshal(read, &auth)
	if err != nil {
		return err
	}

	err = h.repository.Insert(h.ctx, auth)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "yeah, author has been created")
}

func (h *handlerAuthor) all(c echo.Context) error {
	authors, err := h.repository.All(h.ctx)
	if err != nil {
		return err
	}

	var newLine = byte(10)

	for _, auth := range authors {
		res, err := json.Marshal(auth)
		res = append(res, newLine)
		_, err = c.Response().Write(res)
		if err != nil {
			return err
		}

	}

	return c.String(http.StatusOK, "There are authors")
}

func (h *handlerAuthor) delete(c echo.Context) (err error) {
	var read []byte
	auth := models.Author{}

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

	err = json.Unmarshal(read, &auth)
	if err != nil {
		return err
	}

	err = h.repository.Delete(h.ctx, auth)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "yeah, author has been deleted")
}
