package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
)

type UserServices interface {
	FindUser(id string) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) FindUser(id string) (entities.User, error) {

	getUser, err := s.userRepository.GetUser(id)
	if err != nil {
		return getUser, err
	}

	return getUser, nil

}
