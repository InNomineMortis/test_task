package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"test_task/internal/pkg/db_connect"
	"test_task/internal/pkg/server/handler"
)

func RunServer() {
	database := db_connect.InitDB("postgres")
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/posts", handler.GetPosts(database))
	err := router.Run(":7777")
	if err != nil {
		log.Fatal("There is an Error running gin router: ", err.Error())
	}

	defer database.Close()
}
