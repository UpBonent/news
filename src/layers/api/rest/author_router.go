package rest

//Returning errors into handler:
//- send it to user?
//- or log it in a server side?

import (
	"context"
	"github.com/UpBonent/news/src/common/models"
	"github.com/pkg/errors"
	"net/http"
	"time"

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

	g.GET("/create", h.newAuthorForm)
	g.POST("/create", h.newAuthor)

	profile := g.Group("/profile")
	profile.POST("", h.viewUserProfile)
	profile.GET("/auth", h.viewAuthenticationForm)
	profile.POST("/auth", h.authentication)

	profile.GET("/new", h.viewNewArticleForm)
	profile.POST("/new", h.newArticle)
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

func (h *handlerAuthor) newAuthorForm(c echo.Context) (err error) {
	return c.File("./static/html/new_author_form.html")
}

func (h *handlerAuthor) newAuthor(c echo.Context) (err error) {
	defer func() {
		err := c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	author := models.Author{
		Name:     c.FormValue("name"),
		Surname:  c.FormValue("surname"),
		UserName: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	checkPassword := c.FormValue("check_password")

	_, cookieValue, err := h.application.CreateNewAuthor(h.ctx, author, checkPassword)
	if err != nil {
		return
	}

	cookie, err := h.application.SetUserCookie(cookieValue)
	c.SetCookie(&cookie)

	return c.String(http.StatusCreated, "yeah, author has been created")
}

func (h *handlerAuthor) viewAuthorsArticles(c echo.Context) (err error) {

	cookie, err := c.Cookie("i")
	if err != nil {
		return err
	}

	author, err := h.application.GetAuthorByCookie(cookie.String())
	if err != nil {
		return err
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

func (h *handlerAuthor) viewUserProfile(c echo.Context) (err error) {
	err = c.File("./static/html/author_profile.html")
	if err != nil {
		return err
	}

	cookie, err := c.Cookie("i")
	if err != nil {
		return err
	}

	author, err := h.application.GetAuthorByCookie(cookie.String())
	if err != nil {
		return err
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

func (h *handlerAuthor) viewAuthenticationForm(c echo.Context) error {
	return c.File("./static/html/auth_author_form.html")
}

func (h *handlerAuthor) authentication(c echo.Context) (err error) {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.String(http.StatusBadRequest, "fill all fields")
	}

	err = h.application.CheckUserAuthentication(username, password)
	if err != nil {
		if errors.As(err, "Wrong username and/or password") {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusInternalServerError, "Try again")
	}

	cookieValue, err := h.application.GetCookieByUserName(username)
	cookie, err := h.application.SetUserCookie(cookieValue)
	if err != nil {
		return
	}

	c.SetCookie(&cookie)

	return c.Redirect(http.StatusPermanentRedirect, "/authors/profile")
}

func (h *handlerAuthor) viewNewArticleForm(c echo.Context) (err error) {
	return c.File("./static/html/new_article_form.html")
}

func (h *handlerAuthor) newArticle(c echo.Context) (err error) {
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			return
		}
	}()

	datePublish, err := time.Parse("2006-01-02", c.FormValue("date_publish"))
	if err != nil {
		return c.String(http.StatusBadRequest, "incorrect date, try again")
	}

	article := models.Article{
		Header:      c.FormValue("header"),
		Text:        c.FormValue("text"),
		Annotation:  c.FormValue("annotation"),
		DatePublish: datePublish,
	}

	//id, err := h.application.GetIDByAuthor(h.ctx, author)
	//if err != nil {
	//	return err
	//}

	err = h.application.CreateNewArticle(h.ctx, article, 1)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "yeah, article has been created")
}
