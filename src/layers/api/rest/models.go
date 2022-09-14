package rest

type AuthorJSON struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type ArticleJSON struct {
	Id          int    `json:"id"`
	Header      string `json:"header"`
	Text        string `json:"text"`
	DatePublish string `json:"date_publish"`
	IdAuthor    int    `json:"id_author"`
}
