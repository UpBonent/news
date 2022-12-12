package rest

import (
	"encoding/json"
	"github.com/UpBonent/news/src/common/models"
	"time"
)

type AuthorJSON struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Salt     string `json:"salt"`

	Activity bool `json:"activity"`
}

type ArticleJSON struct {
	Id          int    `json:"id"`
	Header      string `json:"header"`
	Text        string `json:"text"`
	DateCreate  string `json:"date_create"`
	DatePublish string `json:"date_publish"`
	AuthorID    int    `json:"author_id"`
}

// JSON-string to Model
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
		UserName: authorJSON.UserName,
		Password: authorJSON.Password,
		Activity: authorJSON.Activity,
	}
	return
}

func convertArticleJSONtoModel(reader []byte) (article models.Article, err error) {
	articleJSON := ArticleJSON{}

	err = json.Unmarshal(reader, &articleJSON)
	if err != nil {
		return
	}

	datePublish, err := time.Parse("02.01.06 15:04", articleJSON.DatePublish)
	if err != nil {
		return
	}

	article = models.Article{
		Id:          articleJSON.Id,
		Header:      articleJSON.Header,
		Text:        articleJSON.Text,
		DatePublish: datePublish,
		AuthorID:    articleJSON.AuthorID,
	}
	return
}

// Model to JSON-struct
func convertAuthorModelToJSON(author models.Author) AuthorJSON {
	return AuthorJSON{
		Id:       author.Id,
		Name:     author.Name,
		Surname:  author.Surname,
		UserName: author.UserName,
		Activity: author.Activity,
	}
}

func convertArticleModelToJSON(article models.Article) ArticleJSON {
	return ArticleJSON{
		Id:          article.Id,
		Header:      article.Header,
		Text:        article.Text,
		DateCreate:  article.DateCreate.Format("02.01.06 15:04"),
		DatePublish: article.DatePublish.Format("2006-01-02 15:04:.000"),
		AuthorID:    article.AuthorID,
	}
}
