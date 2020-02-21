package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test_task/internal/pkg/db_connect"
	"test_task/internal/pkg/models"
)

func RunServer() {
	database := db_connect.InitDB("postgres@postgres:5432", "test")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/posts", func(c *gin.Context) {

		var res []models.Post
		var resOne models.Post
		qres, err := database.Queryx(`SELECT * FROM post
 JOIN person on post.person_id = person.id
 JOIN address a on person.id = a.person_id`)
		if err != nil {
			log.Fatal("Error while querying from db")
		}
		for qres.Next() {
			err = qres.StructScan(&resOne)
			fmt.Println(resOne)
			res = append(res, resOne)
		}
		c.JSON(http.StatusOK, res)
	})
	http.ListenAndServe(":7777", router)

	defer database.Close()
}
