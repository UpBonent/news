package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo"

	"github.com/UpBonent/news/src/common/services"
)

type handlerArticle struct {
	ctx         context.Context
	application services.Application
}

func NewHandlersArticle(ctx context.Context, app services.Application) services.Handler {
	return &handlerArticle{ctx, app}
}

func (h *handlerArticle) Register(e *echo.Echo) {
	g := e.Group("/articles")
	g.GET("", h.all)
	g.PUT("/update", h.update)
}

func (h *handlerArticle) all(c echo.Context) (err error) {
	var allArticlesJSON []ArticleJSON
	articles, err := h.application.GetAllArticles(h.ctx)
	if err != nil {
		return err
	}

	for _, article := range articles {
		articleJSON := convertArticleModelToJSON(article)
		allArticlesJSON = append(allArticlesJSON, articleJSON)
	}

	return c.JSON(http.StatusOK, allArticlesJSON)
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

	err = h.application.UpdateArticle(h.ctx, existsArticle.IdExists, article)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Update succeeded")
}
