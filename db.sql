create table authors
(
    id       smallserial UNIQUE,
    name     varchar(20) NOT NULL,
    surname  varchar(30) NOT NULL,
    activity bool DEFAULT true,
    username varchar(20) NOT NULL,
    password char(64) NOT NULL,
    salt     char(64) NOT NULL
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

drop table articles;
drop table authors;
