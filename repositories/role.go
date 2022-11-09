package repositories

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role entities.Role) (entities.Role, error)
	FindRoleByID(id string) (entities.Role, error)
	Update(role entities.Role) (entities.Role, error)
	Delete(id string) error
	//GetByRoleId(id string) (entities.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{
		db,
	}
}

func (r *roleRepository) Create(role entities.Role) (entities.Role, error) {
	//db.Omit(clause.Associations).Create(&user)

	err := r.db.Create(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

func (r *roleRepository) FindRoleByID(id string) (entities.Role, error) {
	var role entities.Role
	errFind := r.db.Where("id =?", id).Find(&role).Error
	if errFind != nil {
		return role, errFind
	}
	return role, nil
}
func (r *roleRepository) Update(role entities.Role) (entities.Role, error) {
	err := r.db.Save(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

func (r *roleRepository) Delete(id string) error {
	deleteErr := r.db.Where("id=?", id).Delete(&entities.Role{}).Error
	return deleteErr
}
