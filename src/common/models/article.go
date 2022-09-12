package models

type Article struct {
	Id          int    `db:"id"`
	Header      string `db:"header"`
	Text        string `db:"text"`
	DateCreate  string `db:"date_create"`
	DatePublish string `db:"date_publish"`
	IdAuthor    int    `db:"id_author"`
}
