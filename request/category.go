package request

import "mime/multipart"

type CategoryRequestInsert struct {
	Name          string `form:"name" binding:"required"`
	PublicIDCloud string
	CreatedBy     string
}
type CategoryImageRequest struct {
	Image multipart.File `form:"image" binding:"required"`
}
