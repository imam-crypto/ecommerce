package mappers

import (
	"ecommerce/dtos"
	"ecommerce/entities"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterCreateUser(input dtos.RegisterRequest) entities.User {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return entities.User{}
	}
	Password := string(passwordHash)
	user := entities.User{
		Username: input.Username,
		Email:    input.Email,
		Password: Password,
		RoleID:   uuid.FromStringOrNil("17cbf61e-2d14-46de-8bc2-b6ca3a67ba16"),
	}
	return user
}

func UpdateUser(oldValue entities.User, newValue dtos.UpdateRequest) entities.User {
	oldValue.Address = newValue.Address
	oldValue.Gender = newValue.Gender
	oldValue.Name = newValue.Name
	oldValue.City = newValue.City
	oldValue.Province = newValue.Province
	oldValue.PostalCode = newValue.PostalCode

	return oldValue
}
