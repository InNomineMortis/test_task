-- +migrate Up
CREATE EXTENSION pgcrypto;

CREATE TABLE person(
                       person_id int PRIMARY KEY,
                       avatar_url VARCHAR(70) DEFAULT '',
                       name VARCHAR(50) NOT NULL,
                       surname VARCHAR(70) NOT NULL,
                       patronymic VARCHAR(70) NOT NULL
);

CREATE TABLE post (
                      post_id int PRIMARY KEY,
                      person_id int NOT NULL,
                      header VARCHAR(120) NOT NULL,
                      post_text VARCHAR(360) NOT NULL,
                      timestamp VARCHAR(32) NOT NULL,
                      FOREIGN KEY (person_id) references person(person_id)
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
                        flat varchar(10) NOT NULL,
                        FOREIGN KEY (person_id) references person(person_id)
);
INSERT INTO person(person_id, avatar_url, name,
                   surname, patronymic) VALUES (1,
                                                'skdjsada', 'somename', 'somesurname', 'thirdname');
INSERT INTO person(person_id, avatar_url, name,
                   surname, patronymic) VALUES (2,
                                                'anotherone', 'somename2', 'somesurname2', 'thirdname2');
INSERT INTO post(post_id, person_id, header, post_text, timestamp
) VALUES (1,(select person_id from person where name ='somename'), 'someheader', 'so mmuch text, o my god','21.01.01');

INSERT INTO post(post_id, person_id, header, post_text, timestamp
) VALUES (2,(select person_id from person where name ='somename2'), 'someheader2323', 'so mmuch text, o my god, shock content','21.01.01');

INSERT INTO address(address_id, person_id, index, country, region,
                    city, street, metro, house_number, section, flat) VALUES (1,
                                                                              (select person_id from person where name = 'somename'),
                                                                              143112,
                                                                              'Russia',
                                                                              'Moscow reg',
                                                                              'kaliningrad',
                                                                              'lenina prospekt',
                                                                              'NET TAM METRO',
                                                                              123,
                                                                              'WHAT',
                                                                              '21323'
                                                                             );

INSERT INTO address(address_id, person_id, index, country, region,
                    city, street, metro, house_number, section, flat) VALUES (2,
                                                                              (select person_id from person where name = 'somename2'),
                                                                              223312,
                                                                              'POLAND',
                                                                              'WARSAW reg',
                                                                              'WARSAW',
                                                                              'lenina prospekt',
                                                                              'Dworzec Gdański',
                                                                              123,
                                                                              'kurwa',
                                                                              '21323'
                                                                             );

INSERT INTO comment(comment_id, post_id, comment_text, person_id) VALUES (1,
                                                                          (select post_id from post where header = 'someheader'),
                                                                          'VSE FIGNIA DAVAI PO NOVOI',
                                                                          (select person_id from person where name = 'somename')
                                                                         );
-- +migrate Down
DROP TABLE comment cascade ;
DROP TABLE post cascade ;
DROP TABLE address cascade ;
DROP TABLE person cascade;