package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	uuid "github.com/satori/go.uuid"
)

type RoleService interface {
	Create(loggedUser uuid.UUID, input request.RoleRequest) (entities.Role, error)
}

type roleService struct {
	roleRepository repositories.RoleRepository
}

func NewRoleService(roleRepository repositories.RoleRepository) *roleService {
	return &roleService{
		roleRepository,
	}
}

func (s *roleService) Create(loggedUser uuid.UUID, input request.RoleRequest) (entities.Role, error) {
	mapCreate := entities.Role{
		Base: entities.Base{
			CreatedBy: loggedUser,
		},
		Name:       input.Name,
		MenuAccess: input.MenuAccesses,
	}
	createRole, errCreate := s.roleRepository.Create(mapCreate)
	if errCreate != nil {
		return mapCreate, errCreate
	}
	return createRole, nil
}
