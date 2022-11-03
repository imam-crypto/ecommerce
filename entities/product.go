package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	Base

	CategoryID  uuid.UUID `gorm:"type:uuid;"`
	Title       string
	Description string
	Price       int
	Variant     []Variant `gorm:"foreignKey:ProductID"`
	//Category    Category  `gorm:"foreignKey:CategoryID"`
}

func (t *Product) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.NewV4()
	if t.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
