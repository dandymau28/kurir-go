package service

import (
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/repository"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) GetAll() ([]dto.CourierAll, error) {
	couriers, err := s.userRepo.GetAllCourier()

	return couriers, err
}
