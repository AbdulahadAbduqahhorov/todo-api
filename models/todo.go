package todoapi

import "errors"

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct{
	Title *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct{
	Title *string `json:"title"`
	Description *string `json:"description"`
	Done *bool `json:"done"`
}


func (u UpdateListInput) Validate()error{
	if u.Title==nil && u.Description==nil{
		return errors.New("update structure has no values")
	}
	return nil
}
func (u UpdateItemInput) Validate()error{
	if u.Title==nil && u.Description==nil && u.Done==nil{
		return errors.New("update structure has no values")
	}
	return nil
}