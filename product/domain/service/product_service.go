package service

import (
	"product/domain/model"
	"product/domain/repository"
)

type IProductService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}
type ProductService struct {
	ProductRepository repository.IProductRepository
}

func (p ProductService) AddProduct(product *model.Product) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) DeleteProduct(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) UpdateProduct(product *model.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) FindProductByID(i int64) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) FindAllProduct() ([]model.Product, error) {
	//TODO implement me
	panic("implement me")
}

// 创建
func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{productRepository}
}
