package repository

import (
	"cart/domain/model"
	"gorm.io/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

// 创建cartRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (c *CartRepository) InitTable() error {
	return c.mysqlDb.AutoMigrate(&model.Cart{})
}

func (c *CartRepository) FindCartByID(i int64) (*model.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) DeleteCartByID(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) FindAll(i int64) ([]model.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) CleanCart(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) IncrNum(i int64, i2 int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) DecrNum(i int64, i2 int64) error {
	//TODO implement me
	panic("implement me")
}
