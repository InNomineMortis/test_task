-- DROP TABLE post;
-- DROP TABLE person;
-- DROP TABLE comment;


CREATE TABLE post (
    id uuid PRIMARY KEY,
    person_id uuid NOT NULL,
    header VARCHAR(120) NOT NULL,
    text VARCHAR(360) NOT NULL,
    timestamp VARCHAR(32) NOT NULL,
    address VARCHAR(240) NOT NULL,
    comment_id UUID,
    FOREIGN KEY (person_id) references person(id)
);

CREATE TABLE person(
    id uuid PRIMARY KEY,
    avatar_url VARCHAR(70) DEFAULT '',
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(70) NOT NULL,
    patronymic VARCHAR(70) NOT NULL,
    address VARCHAR(120) NOT NULL
);

CREATE TABLE comment(
    id uuid PRIMARY KEY,
    post_id uuid NOT NULL,
    text VARCHAR(240) NOT NULL,
    person_id uuid NOT NULL
);

