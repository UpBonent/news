create table authors
(
    id      smallserial UNIQUE,
    name    varchar(20) NOT NULL,
    surname varchar(30) NOT NULL
);

create table articles
(
    id              smallserial NOT NULL,
    header          varchar(65) NOT NULL,
    text            text NOT NULL,
    date_create     timestamp NOT NULL,
    date_publish    timestamp NOT NULL,
    id_authors      int references authors(id)
);

drop table authors;
drop table articles;

-- !!! Indexes !!!

--for author
INSERT INTO authors(name, surname) VALUES($1, $2);
SELECT name, surname FROM authors;

--for article
INSERT INTO articles(header, text, date_create, date_publish, id_authors) VALUES ($1, $2, $3, $4, (SELECT id FROM authors WHERE name = $5 AND surname = $6));   -- CreateArticle
SELECT header, text, date_publish FROM articles;    -- AllArticles
SELECT header FROM articles WHERE date_publish < $1 && > $2;    -- AllHeaders for the time

--for JOIN
SELECT header, text, authors.name, authors.surname
FROM articles
INNER JOIN authors ON articles.id = authors.id;
