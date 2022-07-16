package main

import (
	todo "github.com/Vladosya/rest-api-go"
	"github.com/Vladosya/rest-api-go/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)                               // берем обработчики, где роуты
	srv := new(todo.Server)                                        // создаем сервер
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil { // делаем проверку если сервер не запустился то вызываем ошибку
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}
