package models

import "time"

type Article struct {
	Id          int       `db:"id"`
	Header      string    `db:"header"`
	Text        string    `db:"text"`
	DateCreate  time.Time `db:"date_create"`
	DatePublish time.Time `db:"date_publish"`
	AuthorID    int       `db:"author_id"`
}
