package repositories

import (
	"ecommerce/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
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
func (r *userRepository) FindByEmail(email string) (entities.User, error) {

	var user entities.User

	findUser := r.db.Where("email =?", email).First(&user).Error

	if findUser != nil {
		return user, findUser
	}
	return user, nil
}
func (r *userRepository) CreateUser(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
