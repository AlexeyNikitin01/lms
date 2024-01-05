CREATE TABLE
    course (
        id serial not null unique,
        name varchar(255) not null,
        progress int not null,
    );

CREATE TABLE
    list_course (
    );

CREATE TABLE
    info (
        id serial not null unique,
        name varchar(255) not null,
        text varchar(255) not null,
    );

CREATE TABLE
    list_info (
    );

CREATE TABLE
    user (
        id serial not null unique,
        name varchar(255) not null,
        email varchar(255) not null,
        login varchar(255) not null,
        password varchar(255) not null,
        course int references list_course (id) on delete cascade,
        info int references list_info (id) on delete cascade,
    );
