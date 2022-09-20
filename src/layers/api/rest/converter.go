package rest

import (
	"encoding/json"
	"github.com/UpBonent/news/src/common/models"
)

func convertAuthorJSONtoModel(reader []byte) (author models.Author, err error) {
	authorJSON := AuthorJSON{}

	err = json.Unmarshal(reader, &authorJSON)
	if err != nil {
		return
	}

	author = models.Author{
		Id:       authorJSON.Id,
		Name:     authorJSON.Name,
		Surname:  authorJSON.Surname,
		Activity: authorJSON.Activity,
	}
	return
}

func convertAuthorModelToJSON(author models.Author) (writer []byte, err error) {
	authorJSON := AuthorJSON{
		Id:       author.Id,
		Name:     author.Name,
		Surname:  author.Surname,
		Activity: author.Activity,
	}

	return json.Marshal(authorJSON)
}

func convertArticleJSONtoModel(reader []byte) (article models.Article, err error) {
	articleJSON := ArticleJSON{}

	err = json.Unmarshal(reader, &articleJSON)
	if err != nil {
		return
	}

	article = models.Article{
		Id:          articleJSON.Id,
		Header:      articleJSON.Header,
		Text:        articleJSON.Text,
		DatePublish: articleJSON.DatePublish,
		AuthorID:    articleJSON.AuthorID,
	}
	return
}

func convertArticleModelToJSON(article models.Article) (writer []byte, err error) {
	articleJSON := ArticleJSON{
		Id:          article.Id,
		Header:      article.Header,
		Text:        article.Text,
		DateCreate:  article.DateCreate,
		DatePublish: article.DatePublish,
		AuthorID:    article.AuthorID,
	}

	return json.Marshal(articleJSON)
}
