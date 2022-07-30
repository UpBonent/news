package postgres

import (
	"github.com/UpBonent/news/src/common/models"
)

func AddAuthorQuery(a models.Author) error {
	result := db.QueryRowx(NewAuthor, a.Name, a.Surname)
	return result.Err()
}

func AllAuthorsQuery() (authors []models.Author, err error) {

	selector, err := db.Queryx(AllAuthors)
	if err != nil {
		return
	}

	for selector.Next() {
		var name, surname string
		var id int
		err = selector.Scan(&id, &name, &surname)
		if err != nil {
			return
		}
		nextAuthor := models.Author{
			Id:      id,
			Name:    name,
			Surname: surname,
		}
		authors = append(authors, nextAuthor)
	}
	return
}

func DeleteAuthorQuery(a models.Author) error {
	result := db.QueryRowx(DeleteAuthor, a.Name, a.Surname)
	return result.Err()
}
