package service

import (
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/repository"
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(user models.SignInInput) (string, error)
	ParseToken(token string)(int,error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {

	return &Service{
		Authorization:NewAuthService(repo.Authorization),
	}
}

