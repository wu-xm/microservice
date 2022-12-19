package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

// 创建
func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func (c CartDataService) AddCart(cart *model.Cart) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c CartDataService) DeleteCart(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CartDataService) UpdateCart(cart *model.Cart) error {
	//TODO implement me
	panic("implement me")
}

func (c CartDataService) FindCartByID(i int64) (*model.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (c CartDataService) FindAllCart(i int64) ([]model.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (c CartDataService) CleanCart(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CartDataService) DecrNum(i int64, i2 int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CartDataService) IncrNum(i int64, i2 int64) error {
	//TODO implement me
	panic("implement me")
}
