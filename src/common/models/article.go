package models

import (
	"time"
)

type Article struct {
	Header      string
	Text        string
	DateCreate  time.Time
	DatePublish time.Time
}
