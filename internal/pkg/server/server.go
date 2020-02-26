package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"test_task/internal/pkg/database"
	"test_task/internal/pkg/server/handlers"
)

func RunServer() {
	database.InitPostgres("postgres")

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("static", true)))
	api := router.Group("api")
	handlers.SetupPosts(api)

	if err := router.RunTLS(":443", "cert.pem", "key.pem"); err != nil {
		log.Fatal("Error handling HTTPS: ", err.Error())
	}
}
