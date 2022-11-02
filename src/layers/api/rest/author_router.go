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
	"github.com/labstack/echo/middleware"
)

type handlerAuthor struct {
	ctx         context.Context
	application services.Application
}

func NewHandlersAuthor(ctx context.Context, app services.Application) services.Handler {
	return &handlerAuthor{ctx, app}
}

func (h *handlerAuthor) Register(e *echo.Echo) {
	g := e.Group("/authors")
	g.GET("", h.viewAllAuthor)

	g.GET("/articles", h.viewAuthorsArticles)

	g.GET("/create", h.viewCreateForm)
	g.POST("/create/new", h.createNewAuthor)

	a := g.Group("/auth")
	a.Use(middleware.BasicAuth())
	a.GET("", h.auth)
}

func (h *handlerAuthor) viewAllAuthor(c echo.Context) (err error) {
	var allAuthorsJSON []AuthorJSON
	authors, err := h.application.GetAllAuthors(h.ctx)
	if err != nil {
		return err
	}

	for _, author := range authors {
		authorJSON := convertAuthorModelToJSON(author)
		allAuthorsJSON = append(allAuthorsJSON, authorJSON)
	}

	return c.JSON(http.StatusOK, allAuthorsJSON)
}

func (h *handlerAuthor) viewCreateForm(c echo.Context) (err error) {
	return c.File("../static/html/authors.html")
}

func (h *handlerAuthor) createNewAuthor(c echo.Context) (err error) {
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

	_, err = h.application.CreateNewAuthor(h.ctx, author)
	if err != nil {
		return
	}

	return c.String(http.StatusCreated, "yeah, author has been created")
}

func (h *handlerAuthor) viewAuthorsArticles(c echo.Context) (err error) {
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

	articles, err := h.application.GetArticlesByAuthorID(h.ctx, author.Id)
	if err != nil {
		return
	}
	//
	//
	//mb it makes sense to restructure code below
	var allArticlesJSON []ArticleJSON
	for _, article := range articles {
		articleJSON := convertArticleModelToJSON(article)
		allArticlesJSON = append(allArticlesJSON, articleJSON)
	}

	return c.String(http.StatusOK, "There are articles by this author")
}

func (h *handlerAuthor) auth(c echo.Context) (err error) {

	return err
}
