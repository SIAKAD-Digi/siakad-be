package services

import (
	"os"
	"siakad-digi/internal/exception"
	"siakad-digi/internal/models/database"
	"siakad-digi/internal/models/request"
	"siakad-digi/internal/repository"
	"siakad-digi/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func (s *AuthService) Login(loginReq *request.LoginRequest) string {
	user, err := s.UserRepository.FindByNik(loginReq.Nik)
	match := utils.Compare(user.Password, loginReq.Password)

	if !match || err != nil {
		panic(exception.NewUnauthorizeError("nik atau password salah"))
	}

	jwt := s.generateJwt(user)
	return jwt
}

func (a *AuthService) generateJwt(u *database.User) string {
	key := os.Getenv("JWT_SECRET_KEY")
	exp := time.Now().Add(time.Hour * 24).Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": map[string]any{
		"id": u.ID, "name": u.Name, "role": u.RoleId,
	}, "exp": exp})

	jwt, _ := t.SignedString([]byte(key))
	return jwt
}
