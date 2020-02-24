DROP TABLE comment;
DROP TABLE post;
DROP TABLE address;
DROP TABLE person;

-- CREATE EXTENSION pgcrypto;

CREATE TABLE person(
   person_id int PRIMARY KEY,
   avatar_url VARCHAR(70) DEFAULT '',
   name VARCHAR(50) NOT NULL,
   surname VARCHAR(70) NOT NULL,
   patronymic VARCHAR(70) NOT NULL,
   address int NOT NULL,
   FOREIGN KEY (address) references address(address_id)
);

CREATE TABLE post (
    post_id int PRIMARY KEY,
    person_id int NOT NULL,
    header VARCHAR(120) NOT NULL,
    post_text VARCHAR(360) NOT NULL,
    timestamp VARCHAR(32) NOT NULL,
    address_id int NOT NULL,
    FOREIGN KEY (person_id) references person(person_id),
    FOREIGN KEY (address_id) references address(address_id)
);

CREATE TABLE comment(
    comment_id int PRIMARY KEY,
    post_id int NOT NULL,
    comment_text VARCHAR(240) NOT NULL,
    person_id int NOT NULL,
    FOREIGN KEY (person_id) references person(person_id),
    FOREIGN KEY (post_id) references post(post_id)
);

CREATE TABLE address(
    address_id int PRIMARY KEY,
    person_id int NOT NULL,
    index int NOT NULL,
    country varchar(20) NOT NULL,
    region varchar(30) NOT NULL,
    city varchar(15) NOT NULL,
    street varchar(30) NOT NULL,
    metro varchar(15) NOT NULL,
    house_number int NOT NULL,
    section varchar(20) NOT NULL,
    flat varchar(10) NOT NULL
);