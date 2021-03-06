package service

import "github.com/Vladosya/rest-api-go/pkg/repository"

type Authorization interface {
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
