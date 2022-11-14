package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

const(
	usersTable = "users"
	todoListsTable = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable = "todo_items"
	listsItemsTable = "lists_items"
)

func NewPostgres(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.DBName,
		config.SSLMode))
	if err != nil{
		return nil,err
	}

	return db, nil
}
