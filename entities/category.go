package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name          string
	PublicIDCloud string
	UrlImage      string
	//Category      Product `gorm:"foreignKey:CategoryID"`
}

func (t *Category) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.NewV4()
	if t.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
