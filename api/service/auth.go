package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"time"

	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/repository"
	models "github.com/AbdulahadAbduqahhorov/gin/todo-api/models"
	"github.com/golang-jwt/jwt"
)

const (
	salt            = "adsjadshalkjdhasddasdasdsdxcvslvs"
	sampleSecretKey = "SecretYouShouldHide"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = genereatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

type tokenClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func (s *AuthService) GenerateToken(entity models.SignInInput) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	user, err := s.repo.GetUser(entity.Username, genereatePasswordHash(entity.Password))
	if err != nil {
		return "", err
	}

	claims := &tokenClaim{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(sampleSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken,&tokenClaim{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")

		}
		return []byte(sampleSecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaim)
	if !ok{
		return 0,errors.New("token claims are not of type *tokenClaim")
	}
	return claims.UserID,nil
}

func genereatePasswordHash(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum([]byte(salt)))
}
