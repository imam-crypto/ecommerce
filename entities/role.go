package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Role struct {
	Base
	Name       string
	MenuAccess []MenuAccess `gorm:"foreignKey:RoleID;"`
}

func (t *Role) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.NewV4()
	if t.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
