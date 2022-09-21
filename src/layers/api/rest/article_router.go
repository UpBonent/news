package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo"

	"github.com/UpBonent/news/src/common/services"
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

func NewHandlerArticle(ctx context.Context, s string, article services.ArticleRepository, author services.AuthorRepository) services.Handler {
	return &handlerArticle{ctx, s, article, author}
}

func (h *handlerArticle) Register(e *echo.Echo) {
	g := e.Group(h.way)
	g.GET("", h.all)
	g.POST(wayToCreate, h.create)
	g.GET(wayToCreate, h.example)
	g.PUT(wayToUpDate, h.update)
}

func (h *handlerArticle) all(c echo.Context) error {
	articles, err := h.articleRepository.All(h.ctx)
	if err != nil {
		return err
	}

	var newLine = byte(10)

	for _, article := range articles {
		author, err := h.authorRepository.GetAuthorByID(h.ctx, article.AuthorID)
		if err != nil {
			return err
		}

		articleJSON, err := convertArticleModelToJSON(article)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		authorJSON, err := convertAuthorModelToJSON(author)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
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

func (h *handlerArticle) create(c echo.Context) (err error) {
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	reader, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	article, err := convertArticleJSONtoModel(reader)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	author, err := convertAuthorJSONtoModel(reader)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id, err := h.authorRepository.GetIDByName(h.ctx, author)
	if err != nil {
		return err
	}

	err = h.articleRepository.Insert(h.ctx, article, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "yeah, article has been created")
}

// DELETE example later
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

func (h *handlerArticle) update(c echo.Context) (err error) {
	//another option: get all info in one structure (models.Article) without 'existsArticle'
	existsArticle := struct {
		IdExists int `json:"id_exist"`
	}{}
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	reader, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	article, err := convertArticleJSONtoModel(reader)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(reader, &existsArticle)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.articleRepository.Update(h.ctx, existsArticle.IdExists, article)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Update succeeded")
}
