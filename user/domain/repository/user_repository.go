package repository

import (
	"gorm.io/gorm"
	"user/domain/model"
)

type IUserRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// FindUserByName 根据用户名称查找用户信息
	FindUserByName(name string) (*model.User, error)
	// FindUserByID 根据用户ID查找用户信息
	FindUserByID(userId int64) (*model.User, error)
	// CreateUser 创建用户
	CreateUser(*model.User) (int64, error)
	// DeleteUserByID 根据用户ID删除用户
	DeleteUserByID(userId int64) error
	// UpdateUser 更新用户信息
	UpdateUser(*model.User) error
	// FindALL 查找所有用户
	FindALL([]model.User) error
}
type UserRepository struct {
	mysqlDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDb.AutoMigrate(&model.User{})
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDb.Where("user_name=?", name).Find(user).Error
}

func (u *UserRepository) FindUserByID(userId int64) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDb.First(user, userId).Error
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(userId int64) error {
	return u.mysqlDb.Where("id = ?", userId).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Updates(user).Error
}

func (u *UserRepository) FindALL(users []model.User) error {
	return u.mysqlDb.Find(&users).Error
}
