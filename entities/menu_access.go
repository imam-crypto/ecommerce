package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MenuAccess struct {
	Base
	MenuID uuid.UUID `gorm:"type:uuid;" json:"menu_id"`
	RoleID uuid.UUID `gorm:"type:uuid;" json:"role_id"`
	//Menu   Menu      `gorm:"foreignKey:MenuID"`
	//Role         []Role    `gorm:"foreignKey:RoleID;"`
	ReadAccess   bool `json:"read_access"`
	CreateAccess bool `json:"create_access"`
	UpdateAccess bool `json:"update_access"`
	DeleteAccess bool `json:"delete_access"`
}

func (t *MenuAccess) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.NewV4()
	if t.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
