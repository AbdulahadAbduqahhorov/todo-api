package service

import (
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/repository"
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (t *TodoItemService) Create(userId, listId int, entity models.TodoItem) (int, error) {
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return t.repo.Create(listId, entity)
}

func (t *TodoItemService) GetAll(userId, listId int) ([]models.TodoItem, error) {

	return t.repo.GetAll(userId,listId)
}

func (t *TodoItemService) GetById(userId, listId,itemId int) (models.TodoItem, error){
	
	return t.repo.GetById(userId,listId,itemId)
}

func (t *TodoItemService) Update(userId, listId,itemId int, entity models.UpdateItemInput) error{
	if err:=entity.Validate();err!=nil{
		return err
	}
	return t.repo.Update(userId,listId,itemId, entity)

}

func (t *TodoItemService) Delete(userId, listId,itemId int) error{
	
	return t.repo.Delete(userId,listId,itemId)
}
