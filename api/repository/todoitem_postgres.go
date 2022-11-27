package repository

import (
	"errors"
	"fmt"
	"strings"

	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
	"github.com/jmoiron/sqlx"

)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (p *TodoItemPostgres) Create(listId int, entity models.TodoItem) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, nil
	}
	var itemId int

	query := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1, $2) RETURNING id", todoItemsTable)

	err = tx.QueryRow(query, entity.Title, entity.Description).Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	itemsListQuery := fmt.Sprintf("INSERT INTO %s (item_id,list_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(itemsListQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()

}

func (p *TodoItemPostgres) GetAll(userId, listId int) ([]models.TodoItem, error) {
	query := fmt.Sprintf(`
				SELECT
				 ti.id,
				 ti.title,
				 ti.description,
				 ti.done
				FROM %s as ti
	 			JOIN %s as l ON ti.id = l.item_id
				JOIN %s as u ON l.list_id = u.list_id 
	 			WHERE l.list_id = $1
				AND u.user_id=$2`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)

	rows, err := p.db.Queryx(query, listId, userId)
	if err != nil {
		return nil, err
	}
	var todoItem []models.TodoItem

	for rows.Next() {
		var t models.TodoItem

		err := rows.Scan(
			&t.Id,
			&t.Title,
			&t.Description,
			&t.Done,
		)
		if err != nil {
			return todoItem, err
		}
		todoItem = append(todoItem, t)
	}
	return todoItem, nil

}

func (p *TodoItemPostgres) GetById(userId, listId, itemId int) (models.TodoItem, error) {
	var t models.TodoItem
	query := fmt.Sprintf(`
		SELECT
			ti.id,
			ti.title,
			ti.description,
			ti.done
		FROM %s as ti
		JOIN %s as l ON ti.id = l.item_id
		JOIN %s as u ON l.list_id = u.list_id 
		WHERE l.list_id = $1
		AND u.user_id=$2
		AND ti.id=$3
		`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)
	err := p.db.QueryRow(query, listId, userId, itemId).Scan(
		&t.Id,
		&t.Title,
		&t.Description,
		&t.Done,
	)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (p *TodoItemPostgres) Update(userId,listId, itemId int, entity models.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if entity.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d ", argId))
		args = append(args, *entity.Title)
		argId++
	}
	if entity.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d ", argId))
		args = append(args, *entity.Description)
		argId++
	}
	if entity.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d ", argId))
		args = append(args, *entity.Description)
		argId++
	}

	s := strings.Join(setValues, ",")
	query := fmt.Sprintf(`
				UPDATE  %s ti 
				SET %s
				FROM %s li,%s ul 
				WHERE ti.id=li.item_id 
				AND li.list_id=ul.list_id
				AND ul.user_id=$%d,
				AND li.list_id=$%d 
				AND ti.id=$%d `,
		todoItemsTable,
		s,
		listsItemsTable,
		usersListsTable,
		argId,
		argId+1,
		argId+1,
	)

	args = append(args,userId, listId, itemId)

	result, err := p.db.Exec(query, args...)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n > 1 {
		return errors.New("internal server error")
	}
	if n == 0 {
		return errors.New("todo Item not found")
	}
	return nil
}

func (p *TodoItemPostgres) Delete(userId, listId, itemId int) error {

	query := fmt.Sprintf(`
			DELETE FROM %s as ti
			USING %s as ul, %s as li  
			WHERE ti.id=li.item_id
			AND li.list_id=ul.list_id 
			AND ul.user_id=$1 
			AND li.list_id=$2
			AND ti.id=$3
			`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)
	result, err := p.db.Exec(query, userId, listId, itemId)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n > 1 {
		return errors.New("internal server error")
	}
	if n == 0 {
		return errors.New("todo Item not found")
	}
	return nil
}
