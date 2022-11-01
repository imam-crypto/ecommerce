package formatresponse

import "ecommerce/entities"

type CategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UrlImage  string `json:"url_image"`
	CreatedBy string `json:"created_by"`
}

func ConvCategoryResponse(category entities.Category) CategoryResponse {
	categoryResponse := CategoryResponse{
		ID:        category.ID.String(),
		Name:      category.Name,
		UrlImage:  category.UrlImage,
		CreatedBy: category.CreatedBy.String(),
	}
	return categoryResponse
}
