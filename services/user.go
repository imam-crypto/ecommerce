package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	FindUser(id string) (entities.User, error)
	Register(input request.RegisterUserInput) (entities.User, error)
	CheckEmail(email string) (entities.User, error)
	Login(input request.LoginUserInput) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CheckEmail(email string) (entities.User, error) {
	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return user, err
	}
	if user.Username == "" {
		return user, err
	}

	return user, nil
}
func (s *userService) FindUser(id string) (entities.User, error) {

	getUser, err := s.userRepository.GetUser(id)
	if err != nil {
		return getUser, err
	}

	return getUser, nil
}

func (s *userService) Register(input request.RegisterUserInput) (entities.User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return entities.User{}, nil
	}
	Password := string(passwordHash)

	user := entities.User{
		Username: input.Username,
		Email:    input.Email,
		Password: Password,
		Role:     "ADMIN",
	}
	NewUser, errCreate := s.userRepository.CreateUser(user)
	if errCreate != nil {
		return user, errCreate
	}
	return NewUser, nil
}
func (s *userService) Login(input request.LoginUserInput) (entities.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.userRepository.FindByEmail(email)
	fmt.Println("find by email", user.Password)
	fmt.Println("password di login", bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)))

	if err != nil {
		return user, err
	}
	//if user.Username == " " {
	//	return user, errors.New("user not found")
	//}

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if er != nil {
		return user, er
	}
	return user, nil
}
