package models

type Author struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Surname string `db:"surname"`

	Activity bool `db:"activity"`
}
