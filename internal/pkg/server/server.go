package server

import (
	"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"test_task/internal/pkg/db_connect"
	"test_task/internal/pkg/models"
)

func RunServer() {
	database := db_connect.InitDB("postgres@postgres:5432", "test")

	router := gin.Default()
	router.GET("/api/posts", func(c *gin.Context) {

		var res []models.Post
		var resOne models.Post
		qres, err := database.Queryx(`SELECT id, person_id, header, text, timestamp FROM post`)
		if err != nil {
			log.Fatal("Error while querying from db")
		}
		for qres.Next() {
			err = qres.StructScan(&resOne)
			res = append(res, resOne)
		}
		c.JSON(http.StatusOK, res)
	})
	http.ListenAndServe(":7777", router)

	defer database.Close()
}
