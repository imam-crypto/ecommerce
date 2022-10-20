package repositories

import (
	"ecommerce/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id string) (entities.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUser(id string) (entities.User, error) {

	var user entities.User

	findUser := r.db.Where("id =?", id).First(&user).Error

	if findUser != nil {
		return user, findUser
	}
	return user, nil
}
