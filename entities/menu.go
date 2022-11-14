package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Menu struct {
	Base
	MenuName string
	Uri      string
	//MenuID   MenuAccess `gorm:"foreignKey:MenuID"`
}

func (t *Menu) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.NewV4()
	if t.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
