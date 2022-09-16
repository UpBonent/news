create table authors
(
    id       smallserial UNIQUE,
    name     varchar(20) NOT NULL,
    surname  varchar(30) NOT NULL,
    activity bool
);

create table articles
(
    id              smallserial NOT NULL,
    header          varchar(65) NOT NULL,
    annotation      varchar(150) NOT NULL,
    text            text NOT NULL,
    date_create     timestamp NOT NULL,
    date_publish    timestamp NOT NULL,
    author_id       int references authors(id)
);

drop table authors;
drop table articles;

-- !!! Indexes !!!

INSERT INTO authors(name, surname) VALUES ('Don', 'Don')
