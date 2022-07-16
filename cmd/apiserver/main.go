package main

import (
	todo "github.com/Vladosya/rest-api-go"
	"github.com/Vladosya/rest-api-go/pkg/handler"
	"github.com/Vladosya/rest-api-go/pkg/repository"
	"github.com/Vladosya/rest-api-go/pkg/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	initConfig()
	repos := repository.NewRepository()                                       // сначала создаётся репозиторий (работа с базой данных)
	services := service.NewService(repos)                                     // потом сервис который зависит от репозитория (бизнес логика)
	handlers := handler.NewHandler(services)                                  // берем обработчики, где роуты, который зависит от сервиса (роуты и обработчики)
	srv := new(todo.Server)                                                   // создаем сервер
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil { // делаем проверку если сервер не запустился то вызываем ошибку
		log.Fatalf("error accured while running http server in main.go: %s", err.Error())
	}
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	for _, k := range []string{"PORT"} {
		if _, ok := os.LookupEnv(k); !ok {
			log.Fatal("set environment variable -> ", k)
		}
	}
}
