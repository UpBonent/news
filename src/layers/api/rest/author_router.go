package rest

//Returning errors into handler:
//- send it to user?
//- or log it in a server side?

import (
	"context"
	"github.com/labstack/echo/middleware"
	"io"
	"net/http"

	"github.com/UpBonent/news/src/common/services"

	"github.com/labstack/echo"
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
	g.POST("/create", h.createNewAuthor)

	a := g.Group("/profile")
	a.Use(middleware.BasicAuth(h.authentication))

	a.GET("", h.userProfile)
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
		err = c.Request().Body.Close()
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

	var allArticlesJSON []ArticleJSON
	for _, article := range articles {
		articleJSON := convertArticleModelToJSON(article)
		allArticlesJSON = append(allArticlesJSON, articleJSON)
	}

	return c.JSON(http.StatusOK, allArticlesJSON)
}

func (h *handlerAuthor) userProfile(c echo.Context) (err error) {

	return c.String(http.StatusOK, "Welcome")
}

func (h *handlerAuthor) authentication(username, passwordHash string, c echo.Context) (ok bool, err error) {
	ok, err = h.application.CheckUserAuthentication(username, passwordHash)
	if err != nil {
		return false, err
	}

	return
}
