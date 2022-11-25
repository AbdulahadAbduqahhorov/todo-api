package repository

import (
	"fmt"
	"strings"

	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (p *TodoListPostgres) Create(userID int,entity models.TodoList) (int, error) {
	tx, err := p.db.Begin()
	if err!=nil{
		return 0,nil
	}
	var listID int
	 

	query := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1, $2) RETURNING id", todoListsTable)

	
	err= tx.QueryRow(query,entity.Title, entity.Description).Scan(&listID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	usersListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_,err=tx.Exec(usersListQuery,userID,listID)
	if err!=nil{
		tx.Rollback()
		return 0,err
	}
	
	return listID, tx.Commit()

}

func (p *TodoListPostgres) GetAll(userID int) ([]models.TodoList, error){
	query := fmt.Sprintf("SELECT l.id,l.title,l.description FROM %s as l JOIN %s as u ON l.id = u.list_id WHERE u.user_id = $1",todoListsTable,usersListsTable)

	rows,err := p.db.Queryx(query,userID)
	if err!= nil{
        return nil,err
    }
	var todoList []models.TodoList
	
	for rows.Next() {
		var t models.TodoList

		err := rows.Scan(
			&t.Id,
			&t.Title,
			&t.Description,
			
		)
		if err != nil {
			return todoList,err
		}
		todoList = append(todoList, t)
	}
	return todoList,nil

}

func (p *TodoListPostgres) GetById(userId ,listId int)(models.TodoList,error){
	var t models.TodoList
	query := fmt.Sprintf(`SELECT l.id,l.title,l.description FROM %s as l 
							JOIN %s as u ON l.id = u.list_id
							 WHERE u.user_id = $1 AND l.id=$2`,todoListsTable,usersListsTable)
	err:=p.db.QueryRow(query,userId,listId).Scan(
		&t.Id,
		&t.Title,
		&t.Description,
		
	)
	if err!=nil{
        return t,err
    }
	return t,nil
}


func (p *TodoListPostgres) 	Update(userID ,listId int,entity models.UpdateListInput)  error{
	setValues:=make([]string,0)
	args:=make([]interface{},0)
	argId:=1
	if entity.Title !=nil{
		setValues=append(setValues, fmt.Sprintf("title=$%d ",argId))
		args=append(args, *entity.Title)
		argId++
	}
	if entity.Description !=nil{
		setValues=append(setValues, fmt.Sprintf("description=$%d ",argId))
		args=append(args, *entity.Description)
		argId++
	}

	s:=strings.Join(setValues,",")
	query := fmt.Sprintf("UPDATE  %s l SET %s FROM %s u WHERE l.id=u.list_id AND u.user_id=$%d AND u.list_id=$%d ", todoListsTable,s,usersListsTable,argId,argId+1)

	args=append(args, userID,listId)

	_,err:=p.db.Exec(query,args...)
	if err!=nil{
		return err
	}
	return nil

}

func (p *TodoListPostgres) Delete(userId ,listId int)(error){
	
	query := fmt.Sprintf(`DELETE FROM %s as l USING %s as u WHERE l.id=u.list_id AND user_id=$1 AND u.list_id=$2`,todoListsTable,usersListsTable)
	_,err:=p.db.Exec(query,userId,listId)
	if err!=nil{
		return err
	}
	return nil
}