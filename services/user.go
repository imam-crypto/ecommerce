package services

import (
	"ecommerce/dtos"
	"ecommerce/entities"
	"ecommerce/mappers"
	"ecommerce/repositories"
	"ecommerce/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	FindUserByID(id string) (entities.User, error)
	Register(input dtos.RegisterRequest) (entities.User, error)
	CheckEmail(email string) (entities.User, error)
	Login(input dtos.LoginRequest) (entities.User, error)
	UpdateUser(id string, input dtos.UpdateRequest) (entities.User, error)
	GetUsers() ([]entities.User, error)
	FindUserAllPaginate(searchFilter string, pagination utils.Pagination) ([]*entities.User, utils.Pagination)
	//UpdateUserRole(id string, input request.UpdateUserRole) (entities.User, error)
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
func (s *userService) FindUserByID(id string) (entities.User, error) {

	getUser, err := s.userRepository.FindByIdUser(id)
	if err != nil {
		return getUser, err
	}
	//fmt.Println("data user nya", getUser)
	return getUser, nil
}

func (s *userService) Register(input dtos.RegisterRequest) (entities.User, error) {
	user := mappers.RegisterCreateUser(input)
	NewUser, errCreate := s.userRepository.CreateUser(user)
	if errCreate != nil {
		return user, errCreate
	}
	return NewUser, nil
}
func (s *userService) Login(input dtos.LoginRequest) (entities.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.userRepository.FindByEmail(email)
	fmt.Println("find by email", user.Password)
	fmt.Println("password di login", bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)))

	if err != nil {
		return user, err
	}

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if er != nil {
		return user, er
	}
	return user, nil
}
func (s *userService) UpdateUser(id string, input dtos.UpdateRequest) (entities.User, error) {
	oldUser, err := s.userRepository.FindByIdUser(id)
	if err != nil {
		return oldUser, err
	}
	newValue := mappers.UpdateUser(oldUser, input)
	userUpdate, errUpdate := s.userRepository.Update(newValue)
	if errUpdate != nil {
		return userUpdate, errUpdate
	}
	return userUpdate, nil
}
func (s *userService) GetUsers() ([]entities.User, error) {
	getUser, err := s.userRepository.GetUsersWithCache()
	if err != nil {
		return getUser, err
	}

	return getUser, nil
}
func (s *userService) FindUserAllPaginate(searchFilter string, pagination utils.Pagination) ([]*entities.User, utils.Pagination) {
	query := ""
	if searchFilter != "" && query != "" {
		query += " AND LOWER(username) LIKE LOWER('%" + searchFilter + "%') OR LOWER(email) LIKE LOWER('%" + searchFilter + "%')"
	} else if searchFilter != "" && query == "" {
		query += "LOWER(username) LIKE LOWER('%" + searchFilter + "%') OR LOWER(email) LIKE LOWER('%" + searchFilter + "%')"
	}
	users, pagination := s.userRepository.FindAllUsersPaginate(pagination, query)
	return users, pagination
}

//func (s *userService) UpdateUserRole(id string, input request.UpdateUserRole) (entities.User, error) {
//	oldRole, err := s.userRepository.FindByIdUser(id)
//	if err != nil {
//		return oldRole, err
//	}
//	oldRole.RoleID.String() = input.Role
//	updateRole, errRole := s.userRepository.Update(oldRole)
//	if errRole != nil {
//		return oldRole, errRole
//	}
//	return updateRole, nil
//}
