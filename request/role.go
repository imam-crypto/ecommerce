package request

import "ecommerce/entities"

type RoleRequest struct {
	Name         string                `json:"name"`
	MenuAccesses []entities.MenuAccess `json:"menu_accesses"`
}
