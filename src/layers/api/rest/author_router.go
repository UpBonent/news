package rest

//Returning errors into handler:
//- send it to user?
//- or log it in a server side?

import (
	"context"
	"io"
	"net/http"

	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
)

const (
	wayToCreateAuthor     = "/create"
	wayToArticlesByAuthor = "/articles"
)

const (
	paramDiscernibly = "view"
)

type handlerAuthor struct {
	ctx               context.Context
	way               string
	articleRepository services.ArticleRepository
	authorRepository  services.AuthorRepository
}

func NewHandlerAuthor(ctx context.Context, s string, article services.ArticleRepository, author services.AuthorRepository) services.Handler {
	return &handlerAuthor{ctx, s, article, author}
}

func (h *handlerAuthor) Register(e *echo.Echo) {
	g := e.Group(h.way)
	g.GET("", h.allAuthor)
	g.POST(wayToCreateAuthor, h.create)
	g.GET(wayToArticlesByAuthor, h.articlesByAuthor)
}

func (h *handlerAuthor) allAuthor(c echo.Context) (err error) {
	var allAuthorsJSON []AuthorJSON
	authors, err := h.authorRepository.All(h.ctx)
	if err != nil {
		return err
	}

	for _, author := range authors {
		authorJSON := convertAuthorModelToJSON(author)
		allAuthorsJSON = append(allAuthorsJSON, authorJSON)
	}

	return c.JSON(http.StatusOK, allAuthorsJSON)
}

func (h *handlerAuthor) create(c echo.Context) (err error) {
	defer func() {
		err := c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	reader, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	author, err := convertAuthorJSONtoModel(reader)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.authorRepository.Insert(h.ctx, author)
	if err != nil {
		return
	}

	return c.String(http.StatusCreated, "yeah, author has been created")
}

func (h *handlerAuthor) articlesByAuthor(c echo.Context) (err error) {
	defer func() {
		err := c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	reader, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	author, err := convertAuthorJSONtoModel(reader)
	if author.Id == 0 {
		return c.String(http.StatusBadRequest, "author ID is empty or equal zero")
	}

	articles, err := h.articleRepository.GetByAuthorID(h.ctx, author.Id)
	if err != nil {
		return
	}

	var allArticlesJSON []ArticleJSON
	for _, article := range articles {
		articleJSON := convertArticleModelToJSON(article)
		allArticlesJSON = append(allArticlesJSON, articleJSON)
	}

	return c.String(http.StatusOK, "There are articles by this author")
}
