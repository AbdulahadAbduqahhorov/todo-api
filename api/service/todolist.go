package service

import (
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/repository"
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (t *TodoListService) Create(userID int, entity models.TodoList) (int, error){
	return t.repo.Create(userID, entity)
}

func (t *TodoListService) GetAll(userID int)([]models.TodoList,error){
	return t.repo.GetAll(userID)
}

func (t *TodoListService) GetById(userID ,listId int)(models.TodoList,error){
	return t.repo.GetById(userID, listId)
}
func (t *TodoListService) 	Update(userID ,listId int,entity models.UpdateListInput) error{
	if err:=entity.Validate();err!=nil{
		return err
	}
	return t.repo.Update(userID,listId, entity)

}

func (t *TodoListService) 	Delete(userID ,listId int) error{
	return t.repo.Delete(userID,listId)

}
