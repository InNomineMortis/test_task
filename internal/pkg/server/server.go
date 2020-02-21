package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test_task/internal/pkg/db_connect"
)

func RunServer() {
	database := db_connect.InitDB("postgres@postgres:5432", "test")

	router := gin.Default()
	router.GET("/api/posts", func(c *gin.Context) {
		qres := database.QueryRowx(`SELECT id, person_id, header, text, timestamp, address FROM post`)
		fmt.Println(qres)
		c.JSON(200, gin.H{
			"id" : "it should be woriking",
			"Text" : "SOme text",
		})
	})
	http.ListenAndServe(":7777", router)

	defer database.Close()
}
