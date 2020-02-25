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
	defer database.Close()

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5000", "https://localhost:5000", "http://localhost:8081", "https://localhost:8081"}
	corsConfig.AllowCredentials = true
	corsMiddleware := cors.New(corsConfig)
	router.Use(corsMiddleware)
	router.Use(func(c *gin.Context) {
		c.Next()
	})
	router.GET("/api/posts", handler.GetPosts(database))
	go func() {
		if err := router.Run(":7777"); err != nil {
			log.Printf("Error http server starts: %s", err.Error())
			return
		}
	}()

	if err := router.RunTLS(":8080", "cert.pem", "key.pem"); err != nil {
		log.Fatal("Error handling HTTPS: ", err.Error())
	}
}
