package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID            uuid.UUID `json:",omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	CreatedBy     uuid.UUID
	UpdatedBy     uuid.UUID
	DeletedBy     uuid.UUID
	CreatedByType string
	UpdatedByType string
	DeletedByType string
}
