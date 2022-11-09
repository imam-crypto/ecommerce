package repositories

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type MenuAccessRepository interface {
	Create(menuAccess entities.MenuAccess) (bool, error)
}
type menuAccessRepository struct {
	db *gorm.DB
}

func NewMenuAccessRepository(db *gorm.DB) *menuAccessRepository {
	return &menuAccessRepository{db}
}

func (r *menuAccessRepository) Create(menuAccess entities.MenuAccess) (bool, error) {
	err := r.db.Create(&menuAccess).Error
	if err != nil {
		return false, err
	}
	return true, nil

}
