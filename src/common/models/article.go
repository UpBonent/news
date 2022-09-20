package models

type Article struct {
	Id          int    `db:"id"`
	Header      string `db:"header"`
	Text        string `db:"text"`
	DateCreate  string `db:"date_create"`
	DatePublish string `db:"date_publish"`
	AuthorID    int    `db:"author_id"`
}
