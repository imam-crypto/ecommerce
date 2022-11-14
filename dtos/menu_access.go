package dtos

import "ecommerce/entities"

type MenuAccessResponse struct {
	ID           string       `json:"id"`
	Menu         MenuResponse `json:"menu"`
	ReadAccess   bool         `json:"read_access"`
	CreateAccess bool         `json:"create_access"`
	UpdateAccess bool         `json:"update_access"`
	DeleteAccess bool         `json:"delete_access"`
}

func ConvMenuAccessResponse(menuAccess entities.MenuAccess) MenuAccessResponse {
	conv := MenuAccessResponse{
		ID:           menuAccess.ID.String(),
		Menu:         ConvMenuResponse(menuAccess.Menu),
		ReadAccess:   menuAccess.ReadAccess,
		CreateAccess: menuAccess.CreateAccess,
		UpdateAccess: menuAccess.UpdateAccess,
		DeleteAccess: menuAccess.DeleteAccess,
	}
	return conv
}
