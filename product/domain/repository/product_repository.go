package repository

import (
	"gorm.io/gorm"
	"product/domain/model"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}
type ProductRepository struct {
	mysqlDb *gorm.DB
}

// NewProductRepository 创建productRepository
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}
func (p ProductRepository) InitTable() error {
	err := p.mysqlDb.AutoMigrate(&model.Product{}, &model.ProductImage{}, &model.ProductSize{}, &model.ProductSeo{})
	if err != nil {
		return err
	}
	return nil
}

func (p ProductRepository) FindProductByID(i int64) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) DeleteProductByID(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) UpdateProduct(product *model.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) FindAll() ([]model.Product, error) {
	//TODO implement me
	panic("implement me")
}
