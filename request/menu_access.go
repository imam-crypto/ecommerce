package request

import "ecommerce/entities"

type MenuAccessRequest struct {
	MenuID       []string `json:"menu_id"`
	RoleID       string   `json:"role_id"`
	ReadAccess   bool     `json:"read_access"`
	CreateAccess bool     `json:"create_access"`
	UpdateAccess bool     `json:"update_access"`
	DeleteAccess bool     `json:"delete_access"`
	Role         entities.Role
}
