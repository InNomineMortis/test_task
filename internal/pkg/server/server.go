package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"test_task/internal/pkg/config"
	"test_task/internal/pkg/db_connect"
	"test_task/internal/pkg/middleware"
	"test_task/internal/pkg/server/handler"
)

type Server struct {
	Router *mux.Router
}

func NewRouter() (*mux.Router, error) {

	router := mux.NewRouter()
	db_data  := db_connect.OpenSqlxViaPgxConnPool()
	db_data.QueryRowx("SELECT id,name, surname FROM person")
	AccessLogOut := new(middleware.AccessLogger)
	AccessLogOut.StdLogger = log.New(os.Stdout, "STD ", log.LUTC|log.Lshortfile)

	router.Use(middleware.CorsMiddleware)
	router.Use(AccessLogOut.AccessLogMiddleware)

	router.PathPrefix("/api/").Subrouter()

	router.HandleFunc("/posts/", handler.GetPosts).Methods(http.MethodGet, http.MethodOptions)

	return router, nil
}

func RunServer() {
	router, err := NewRouter()
	if err != nil {
		log.Fatal("Failed to create router")
	}
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.MainAppPort), router))
}