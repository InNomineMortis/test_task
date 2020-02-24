package main

import (
	"github.com/kabukky/httpscerts"
	"log"
	"test_task/internal/pkg/server"
)

func main() {
	// Проверяем, доступен ли cert файл.
	err := httpscerts.Check("cert.pem", "key.pem")
	// Если он недоступен, то генерируем новый.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8080")
		if err != nil {
			log.Fatal("Error generating cert: ", err.Error())
		}
	}
	server.RunServer()
}