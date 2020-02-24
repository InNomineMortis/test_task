
INSERT INTO address(address_id, index, country, region,
                    city, street, metro, house_number, section, flat)
VALUES (1, 143112, 'Russia', 'Moscow reg', 'kaliningrad', 'lenina prospekt', 'NET TAM METRO', 123, 'WHAT', '21323'),
       (2, 223312, 'POLAND', 'WARSAW reg', 'WARSAW', 'lenina prospekt', 'Dworzec Gdański', 123, 'kurwa', '21323');

INSERT INTO person(person_id, address_id, avatar_url, name,
                   surname, patronymic)
VALUES (1, (select address_id from address where index = 143112), 'skdjsada', 'somename', 'somesurname', 'thirdname'),
       (2, (select address_id from address where index = 223312), 'anotherone', 'somename2', 'somesurname2',
        'thirdname2');

INSERT INTO post(post_id, person_id, address_id, header, text, timestamp)
VALUES (1, (select person_id from person where name = 'somename'),
        (select address_id from person where name = 'somename'), 'someheader', 'so mmuch text, o my god', '2020-02-10T12:02:44.073Z');

INSERT INTO post(post_id, person_id, address_id, header, text, timestamp)
VALUES (2, (select person_id from person where name = 'somename2'),
        (select address_id from address where index = 143112), 'someheader2323',
        'so mmuch text, o my god, shock content', '2020-02-10T12:02:44.073Z');


INSERT INTO comment(comment_id, post_id, comment_text, person_id)
VALUES (1,
        (select post_id from post where header = 'someheader'),
        'VSE FIGNIA DAVAI PO NOVOI',
        (select person_id from person where name = 'somename'));