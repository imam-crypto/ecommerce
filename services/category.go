package services

import (
	"ecommerce/entities"
	"ecommerce/helpers"
	"ecommerce/repositories"
	"ecommerce/request"
	uuid "github.com/satori/go.uuid"
)

type CategoryService interface {
	CreateCategory(loggedUser uuid.UUID, input request.CategoryRequestInsert, url string) (entities.Category, error)
	ImageUpload(reqImage request.CategoryImageRequest, publicID string) (string, string, error)
	FindByID(id string) (entities.Category, error)
	Update(id string, input request.CategoryRequestInsert, url string) (entities.Category, error)
	Delete(id string, loggedUser uuid.UUID) (bool, error)
}
type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}
func (s *categoryService) ImageUpload(reqImage request.CategoryImageRequest, publicID string) (string, string, error) {
	uploadUrl, publicID, errUpload := helpers.ImageUpload(reqImage.Image, publicID)
	if errUpload != nil {
		return "", "", errUpload
	}
	return uploadUrl, publicID, nil
}
func (s *categoryService) CreateCategory(loggedUser uuid.UUID, input request.CategoryRequestInsert, url string) (entities.Category, error) {
	mapCreate := entities.Category{
		Base: entities.Base{
			CreatedBy: loggedUser,
		},
		Name:          input.Name,
		PublicIDCloud: input.PublicIDCloud,
		UrlImage:      url,
	}
	newCat, err := s.categoryRepository.Create(mapCreate)
	if err != nil {
		return newCat, err
	}
	return newCat, nil
}

func (s *categoryService) FindByID(id string) (entities.Category, error) {
	category, errGet := s.categoryRepository.FindByID(id)
	if errGet != nil {
		return category, errGet
	}
	return category, nil
}

func (s *categoryService) Update(id string, input request.CategoryRequestInsert, url string) (entities.Category, error) {
	oldCategory, err := s.categoryRepository.FindByID(id)
	if err != nil {
		return oldCategory, err
	}
	var newCode = input.PublicIDCloud
	if newCode != "" {
		oldCategory.PublicIDCloud = newCode
	}
	oldCategory.Name = input.Name
	oldCategory.UrlImage = url

	update, errUpdate := s.categoryRepository.Update(oldCategory)
	if errUpdate != nil {
		return update, errUpdate
	}
	return update, nil
}

func (s *categoryService) Delete(id string, loggedUser uuid.UUID) (bool, error) {
	getCategory, errGet := s.categoryRepository.FindByID(id)

	if errGet != nil {
		return false, errGet
	}
	getCategory.DeletedBy = loggedUser
	_, errUpdate := s.categoryRepository.Update(getCategory)
	if errUpdate != nil {
		return false, errUpdate
	}
	//s.categoryRepository.Delete(id)

	return true, nil
}
