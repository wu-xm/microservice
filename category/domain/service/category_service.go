package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type ICategoryService interface {
	AddCategory(category *model.Category) (int64, error)
	DeleteCategory(categoryId int64) error
	UpdateCategory(category *model.Category) error
	FindCategoryByID(categoryId int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(categoryName string) (*model.Category, error)
	FindCategoryByLevel(categoryLevel uint32) ([]model.Category, error)
	FindCategoryByParent(categoryParent int64) ([]model.Category, error)
}
type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}

func NewCategoryService(categoryRepository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{
		CategoryRepository: categoryRepository,
	}
}
func (c CategoryService) AddCategory(category *model.Category) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) DeleteCategory(categoryId int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) UpdateCategory(category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) FindCategoryByID(categoryId int64) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) FindAllCategory() ([]model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) FindCategoryByName(categoryName string) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) FindCategoryByLevel(categoryLevel uint32) ([]model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) FindCategoryByParent(categoryParent int64) ([]model.Category, error) {
	//TODO implement me
	panic("implement me")
}
