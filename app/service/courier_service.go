package service

import (
	"github.com/kurir-go/app/dto"
)

type UserService interface {
	GetAll() ([]dto.CourierAll, error)
}
