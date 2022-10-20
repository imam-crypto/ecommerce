package entities

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	Base
	Name               string
	UserName           string
	NoKTP              string
	Gender             string
	Address            string
	City               string
	Province           string
	PostalCode         string
	Password           string
	Image              string
	ResetPasswordToken string
	Role               string
}

func (t *User) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.NewV4()
	if t.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
