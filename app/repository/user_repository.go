package repository

import (
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/model"
)

type UserRepository interface {
	FindByUsername(username string) dto.User
	CheckUsernameExist(username string) bool
	SaveUser(*model.User) error
	GetAllCourier() ([]dto.CourierAll, error)
}
