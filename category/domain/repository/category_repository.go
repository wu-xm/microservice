package repository

import (
	"category/domain/model"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(categoryId int64) (*model.Category, error)
	CreateCategory(category *model.Category) (int64, error)
	DeleteCategoryByID(categoryId int64) error
	UpdateCategory(category *model.Category) error
	FindAll() ([]model.Category, error)
	FindCategoryByName(categoryName string) (*model.Category, error)
	FindCategoryByLevel(categoryLevel uint32) ([]model.Category, error)
	FindCategoryByParent(categoryParent int64) ([]model.Category, error)
}
type CategoryRepository struct {
	mysqlDb *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}
func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.AutoMigrate(&model.Category{})
}
func (u *CategoryRepository) FindCategoryByID(categoryId int64) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryRepository) DeleteCategoryByID(categoryId int64) error {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryRepository) FindAll() ([]model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryRepository) FindCategoryByName(categoryName string) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryRepository) FindCategoryByLevel(categoryLevel uint32) ([]model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryRepository) FindCategoryByParent(categoryParent int64) ([]model.Category, error) {
	//TODO implement me
	panic("implement me")
}
