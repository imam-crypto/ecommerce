package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Variant struct {
	Base
	ProductID  uuid.UUID `gorm:"type:uuid;"`
	Sku        string
	Colour     string
	Size       string
	Ingredient string
	Quantity   int
	//Image      []ProductImage
}

func (t *Variant) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.NewV4()
	if t.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
