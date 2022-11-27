package repository

import (
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username,password string) (models.User, error)
}

type TodoList interface {
	Create(userId int,entity models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId ,listId int)(models.TodoList,error)
	Update(userId ,listId int,entity models.UpdateListInput)error
	Delete(userId ,listId int)(error)

}

type TodoItem interface {
	Create(listId int, entity models.TodoItem) (int, error)
	GetAll(userId,listId int) ([]models.TodoItem, error)
	GetById(userId,listId,itemId int)(models.TodoItem,error)
	Update(userId,listId,itemId int,entity models.UpdateItemInput)error
	Delete(userId, listId,itemId int)(error)

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}

