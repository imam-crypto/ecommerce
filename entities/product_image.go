package entities

import uuid "github.com/satori/go.uuid"

type ProductImage struct {
	Base
	VariantID uuid.UUID `gorm:"type:uuid;"`
	PublicID  string
	ImageUrl  string
}
