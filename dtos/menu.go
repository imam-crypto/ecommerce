package dtos

import "ecommerce/entities"

type MenuResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

func ConvMenuResponse(menu entities.Menu) MenuResponse {
	res := MenuResponse{
		Id:   menu.ID.String(),
		Name: menu.MenuName,
		Uri:  menu.Uri,
	}

	return res
}
