package server

import (
	"../middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"../handler"
)

type Server struct {
	Router *mux.Router
}

func NewServer() (*Server, error) {
	server := new(Server)

	router := mux.NewRouter()

	AccessLogOut := new(middleware.AccessLogger)
	AccessLogOut.StdLogger = log.New(os.Stdout, "STD ", log.LUTC|log.Lshortfile)

	router.Use(middleware.CorsMiddleware)
	router.Use(AccessLogOut.AccessLogMiddleware)

	router.PathPrefix("/api/").Subrouter()

	router.HandleFunc("/posts/", handler.GetPosts).Methods(http.MethodGet, http.MethodOptions)

	server.Router = router

	return server, nil
}