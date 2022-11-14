package repository

import (
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
	return 0,nil
}
