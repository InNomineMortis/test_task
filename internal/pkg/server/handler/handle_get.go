package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"test_task/internal/pkg/models"
)

/*
 Returning all posts and comments to them
*/
func GetPosts(database *sqlx.DB) gin.HandlerFunc{
		return func(c *gin.Context){
			c.Request.Header.Set("Access-Control-Allow-Origin", "*")
			c.Request.Header.Set("Access-Control-Allow-Credentials", "true")
			var res []models.Post
			var resOne models.Post
			var parse models.GetPost
			var commentsParse models.Comments
			var commentsOne []models.Comments
			qres, err := database.Queryx(`SELECT post.post_id, header, post_text, timestamp,
												person.person_id, avatar_url, name, surname,
												patronymic, index, country, region, city, street,
												metro, house_number, section, flat FROM post
												JOIN person ON post.person_id = person.person_id
												JOIN address a ON person.person_id = a.person_id`)
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
		}
}