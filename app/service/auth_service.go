package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/kurir-go/app/dto"
)

type AuthService interface {
	Login(username string, password string) (dto.LoginResponse, error)
	Register(username string, name string, password string, role string) error
	ValidateToken(token string) (*jwt.Token, error)
}
