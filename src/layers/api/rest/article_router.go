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

func NewHandlerArticle(ctx context.Context, s string, article services.ArticleRepository, author services.AuthorRepository) services.ArticleHandler {
	return &handlerArticle{ctx, s, article, author}
}

func (h *handlerArticle) Register(e *echo.Echo) {
	g := e.Group(h.way)
	g.GET("", h.All)
	g.POST(wayToCreate, h.Create)
	g.GET(wayToCreate, h.example)
	g.DELETE(wayToDelete, h.Delete)
	g.PUT(wayToUpDate, h.Update)
}

func (h *handlerArticle) All(c echo.Context) error {
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

	return c.String(http.StatusOK, "There are All articles")
}

func (h *handlerArticle) Create(c echo.Context) (err error) {
	var read []byte
	article := models.Article{}
	author := models.Author{}

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

	err = json.Unmarshal(read, &author)
	if err != nil {
		return err
	}

	id, err := h.authorRepository.GetByName(h.ctx, author)

	err = h.articleRepository.Insert(h.ctx, article, id)
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

func (h *handlerArticle) Delete(c echo.Context) (err error) {
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

func (h *handlerArticle) Update(c echo.Context) (err error) {
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

	return c.String(http.StatusOK, "There are All headers")
}
