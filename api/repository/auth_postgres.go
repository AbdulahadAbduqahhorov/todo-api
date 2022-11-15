package repository

import (
	"fmt"

	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	err:= p.db.QueryRow(query,user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id,nil
}

