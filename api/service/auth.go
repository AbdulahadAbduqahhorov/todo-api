package service

import (
	"crypto/sha1"
	"encoding/hex"


	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/repository"
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
)

var salt="adsjadshalkjdhasddasdasdsdxcvslvs"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password=genereatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func  genereatePasswordHash(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum([]byte(salt)))
}
