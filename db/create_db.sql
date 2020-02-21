DROP TABLE comment;
DROP TABLE post;
DROP TABLE address;
DROP TABLE person;

CREATE EXTENSION pgcrypto;

CREATE TABLE person(
   id uuid PRIMARY KEY,
   avatar_url VARCHAR(70) DEFAULT '',
   name VARCHAR(50) NOT NULL,
   surname VARCHAR(70) NOT NULL,
   patronymic VARCHAR(70) NOT NULL
);

CREATE TABLE post (
    id uuid PRIMARY KEY,
    person_id uuid NOT NULL,
    header VARCHAR(120) NOT NULL,
    text VARCHAR(360) NOT NULL,
    timestamp VARCHAR(32) NOT NULL,
    FOREIGN KEY (person_id) references person(id)
);

CREATE TABLE comment(
    id uuid PRIMARY KEY,
    post_id uuid NOT NULL,
    text VARCHAR(240) NOT NULL,
    person_id uuid NOT NULL,
    FOREIGN KEY (person_id) references person(id),
    FOREIGN KEY (post_id) references post(id)
);

CREATE TABLE address(
    id uuid PRIMARY KEY,
    person_id uuid NOT NULL,
    index int NOT NULL,
    country varchar(20) NOT NULL,
    region varchar(30) NOT NULL,
    city varchar(15) NOT NULL,
    street varchar(30) NOT NULL,
    metro varchar(15) NOT NULL,
    house_number int NOT NULL,
    section varchar(20) NOT NULL,
    flat varchar(10) NOT NULL,
    FOREIGN KEY (person_id) references person(id)
);