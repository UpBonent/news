package models

type Author struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
	Salt     string `db:"salt"`
	Cookie   string `db:"cookie"`

	Activity bool `db:"activity"`
}
