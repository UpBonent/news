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
	wayToCreate = "/create"
	wayToDelete = "/delete" //root
	wayToUpDate = "/update"
)

type handlerArticle struct {
	ctx               context.Context
	way               string
	articleRepository services.ArticleRepository
	authorRepository  services.AuthorRepository
}

func NewHandlerArticle(ctx context.Context, s string, art services.ArticleRepository, auth services.AuthorRepository) services.RESTMethod {
	return &handlerArticle{ctx, s, art, auth}
}

func (h *handlerArticle) Register(e *echo.Echo) {
	g := e.Group(h.way)
	g.GET("", h.all)
	g.POST(wayToCreate, h.create)
	g.GET(wayToCreate, h.example)
	g.DELETE(wayToDelete, h.delete)
	g.PUT(wayToUpDate, h.upDate)
}

func (h *handlerArticle) all(c echo.Context) error {
	articles, err := h.articleRepository.All(h.ctx)
	if err != nil {
		return err
	}

	var newLine = byte(10)

	for _, art := range articles {
		auth, err := h.authorRepository.GetByID(h.ctx, art.IdAuthor)
		if err != nil {
			return err
		}

		articleJSON, err := json.Marshal(art)
		if err != nil {
			return err
		}

		authorJSON, err := json.Marshal(auth)
		if err != nil {
			return err
		}

		articleJSON = append(articleJSON, authorJSON...)
		articleJSON = append(articleJSON, newLine)

		_, err = c.Response().Write(articleJSON)
		if err != nil {
			return err
		}
	}

	return c.String(http.StatusOK, "There are all articles")
}

func (h *handlerArticle) create(c echo.Context) (err error) {
	var read []byte
	art := models.Article{}
	auth := models.Author{}

	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	read, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(read, &art)
	if err != nil {
		return err
	}

	err = json.Unmarshal(read, &auth)
	if err != nil {
		return err
	}

	err = h.articleRepository.Insert(h.ctx, art, auth)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "yeah, article has been created")
}

// DELETE example letter
func (h *handlerArticle) example(c echo.Context) (err error) {
	q := `
{
	"header": "",
	"text": "",
	"date_publish": "02.8.22 15:04",
    "name": "Boris",
    "surname": "Pasternak"
}
`
	return c.String(http.StatusOK, q)
}

func (h *handlerArticle) delete(c echo.Context) (err error) {
	var read []byte
	article := models.Article{}
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	read, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(read, &article)
	if err != nil {
		return err
	}

	err = h.articleRepository.Delete(h.ctx, article.Id)
	if err != nil {
		return err
	}

	return c.String(http.StatusResetContent, "Article was deleted")
}

func (h *handlerArticle) upDate(c echo.Context) (err error) {
	var read []byte
	article := models.Article{}
	existArticle := struct {
		Id int `json:"id_exist"`
	}{}
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	read, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(read, &article)
	if err != nil {
		return err
	}

	err = json.Unmarshal(read, &existArticle)
	if err != nil {
		return err
	}

	err = h.articleRepository.UpDate(h.ctx, existArticle.Id, article)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "There are all headers")
}
