package service

import (
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/repository"
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(user models.SignInInput) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, entity models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Update(userId, listId int, entity models.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create( userId,listId int, entity models.TodoItem) (int, error)
	GetAll(userId,listId int) ([]models.TodoItem, error)
	GetById(userId, listId,itemId int) (models.TodoItem, error)
	Update(userId, listId,itemId int, entity models.UpdateItemInput) error
	Delete(userId, listId,itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {

	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem: NewTodoItemService(repo.TodoItem,repo.TodoList),
	}
}
