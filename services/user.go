package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	"ecommerce/utils"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	FindUserByID(id string) (entities.User, error)
	Register(input request.RegisterUserInput) (entities.User, error)
	CheckEmail(email string) (entities.User, error)
	Login(input request.LoginUserInput) (entities.User, error)
	UpdateUser(id string, input request.UpdateUserInput) (entities.User, error)
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
		RoleID:   uuid.FromStringOrNil("17cbf61e-2d14-46de-8bc2-b6ca3a67ba16"),
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

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if er != nil {
		return user, er
	}
	return user, nil
}
func (s *userService) UpdateUser(id string, input request.UpdateUserInput) (entities.User, error) {
	oldUser, err := s.userRepository.FindByIdUser(id)
	if err != nil {
		return oldUser, err
	}
	oldUser.Address = input.Address
	oldUser.Gender = input.Gender
	oldUser.Name = input.Name
	oldUser.City = input.City
	oldUser.Province = input.Province
	oldUser.PostalCode = input.PostalCode

	userUpdate, errUpdate := s.userRepository.Update(oldUser)
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
