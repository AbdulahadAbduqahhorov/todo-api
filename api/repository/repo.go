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
	Create(userID int,entity models.TodoList) (int, error)
	GetAll(userID int) ([]models.TodoList, error)
	GetById(userID ,listId int)(models.TodoList,error)
	Update(userID ,listId int,entity models.UpdateListInput)error
	Delete(userID ,listId int)(error)

}

type TodoItem interface {
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
	}
}
