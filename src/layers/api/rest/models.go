package rest

type AuthorJSON struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`

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
