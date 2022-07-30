package models

type Article struct {
	Id          int    `json:"id"`
	Header      string `json:"header"`
	Text        string `json:"text"`
	DateCreate  string `json:"date_create"`
	DatePublish string `json:"date_publish"`
	IdAuthor    int    `json:"id_author"`
}
