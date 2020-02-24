package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"test_task/internal/pkg/db_connect"
	"test_task/internal/pkg/models"
)

func RunServer() {
	database := db_connect.InitDB("127.0.0.1:5432", "postgres")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/posts", func(c *gin.Context) {
		var res []models.Post
		var resOne models.Post
		var parse models.GetPost
		var commentsParse models.Comments
		var commentsOne []models.Comments
		qres, err := database.Queryx(
					`select post.post_id, header, post_text, timestamp, person.person_id, avatar_url, name, surname, patronymic, index, country, region, city, street, metro, house_number, section, flat from post
							join person on post.person_id = person.person_id
							join address a on person.person_id = a.person_id`)
		if err != nil {
			log.Fatal("Error while querying from db: ", err.Error())
		}
		var qresComment *sqlx.Rows
		for qres.Next() {
			err = qres.StructScan(&parse)
			if err != nil {
				fmt.Printf("Error while scanning: %s", err.Error())
			}
			//JSON forming
			resOne.ID = parse.PostId
			resOne.Header = parse.Header
			resOne.Text = parse.Text
			resOne.Persons.ID = parse.PersonId
			resOne.Persons.AvatarURL = parse.AvatarURL
			resOne.Persons.Name = parse.Name
			resOne.Persons.Surname = parse.Surname
			resOne.Timestamp = parse.Timestamp
			resOne.Persons.Patronymic = parse.Patronymic
			resOne.Persons.Address.City = parse.City
			resOne.Persons.Address.Index = parse.Index
			resOne.Persons.Address.Country = parse.Country
			resOne.Persons.Address.Flat = parse.Flat
			resOne.Persons.Address.HouseNumber = parse.HouseNumber
			resOne.Persons.Address.Section = parse.Section
			resOne.Persons.Address.Street = parse.Street
			resOne.Persons.Address.Metro = parse.Metro
			resOne.Persons.Address.Region = parse.Region
			qresComment, err = database.Queryx(`SELECT * from comment WHERE post_id = $1`, resOne.ID)
			if err != nil {
				log.Fatal("Error while querying comment: ", err.Error())
			}
			commentsOne = nil
			for qresComment.Next() {
				err := qresComment.StructScan(&commentsParse)
				if err != nil {
					log.Fatal("Error while scanning comments: ", err.Error())
				}
				commentsOne = append(commentsOne, commentsParse)
			}
			resOne.Responses = commentsOne
			res = append(res, resOne)
		}
		c.JSON(http.StatusOK, res)
	})

	http.ListenAndServe(":7777", router)

	defer database.Close()
}
