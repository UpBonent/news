package postgres

//	for the author table
const (
	NewAuthor     = `INSERT INTO authors(name, surname) VALUES($1, $2)`
	DeleteAuthor  = `DELETE FROM authors WHERE name = $1 AND surname = $2`
	AllAuthors    = `SELECT id, name, surname FROM authors`
	GetAuthorByID = `SELECT name, surname FROM authors WHERE id = $1`
)

//	for the article table
const (
	CreateArticle    = `INSERT INTO articles(header, text, date_create, date_publish, id_authors) VALUES ($1, $2, $3, $4, (SELECT id FROM authors WHERE name = $5 AND surname = $6))`
	AllArticles      = `SELECT header, text, date_publish, id_authors FROM articles`
	AllHeaders       = `SELECT header, date_publish FROM articles`
	GetHeadersByTime = `SELECT header, date_publish FROM articles WHERE date_publish < $1 && > $2`
)

// inner JOIN
const (
	ArticlesByAuthor = `SELECT header, text, authors.name, authors.surname
		FROM articles
		INNER JOIN authors ON articles.id = authors.id;`
)
