package main

import (
	todo "github.com/Vladosya/rest-api-go"
	"github.com/Vladosya/rest-api-go/pkg/handler"
	"github.com/Vladosya/rest-api-go/pkg/repository"
	"github.com/Vladosya/rest-api-go/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()                            // сначала создаётся репозиторий
	services := service.NewService(repos)                          // потом сервис который зависит от репозитория
	handlers := handler.NewHandler(services)                       // берем обработчики, где роуты, который зависит от сервиса
	srv := new(todo.Server)                                        // создаем сервер
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil { // делаем проверку если сервер не запустился то вызываем ошибку
		log.Fatalf("error accured while running http server: %s", err.Error())
	}
}
