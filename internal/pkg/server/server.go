package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"test_task/internal/app"
	"test_task/internal/pkg/db_connect"
	"test_task/internal/pkg/server/handler"
)

func RunServer() {
	database := db_connect.InitDB("postgres")
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/posts", handler.GetPosts(database))
	errs := make(chan error)
	go func() {
		err := http.ListenAndServe(":7777", handlers.CORS()(http.HandlerFunc(app.Redirect)))
		if err != nil {
			errs <- err
		}
	}()
	err := router.RunTLS(":8080", "cert.pem", "key.pem")
	if err != nil {
		log.Fatal("Error handling HTTPS: ", err.Error())
	}

	defer database.Close()
}
