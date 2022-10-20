package repository

import (
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repository *UserRepositoryImpl) FindByUsername(username string) dto.User {
	var user dto.User
	repository.db.Model(&model.User{}).Where("username = ?", username).Take(&user)
	return user
}

func (repository *UserRepositoryImpl) CheckUsernameExist(username string) bool {
	var user dto.User
	repository.db.Model(&model.User{}).Where("username = ?", username).Take(&user)
	return (user != dto.User{})
}

func (repository *UserRepositoryImpl) SaveUser(user *model.User) error {
	result := repository.db.Create(user)
	return result.Error
}

func (repository *UserRepositoryImpl) GetAllCourier() ([]dto.CourierAll, error) {
	roleQuery := "kurir"
	var couriers []dto.CourierAll
	result := repository.db.Model(&model.User{}).Select("name", "username", "role", "id").Where("role = ?", roleQuery).Find(&couriers)

	return couriers, result.Error
}
