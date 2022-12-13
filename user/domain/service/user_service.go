package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domain/model"
	"user/domain/repository"
)

type IUserService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(userId int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(name string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}
type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{UserRepository: userRepository}
}

// 加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// 验证用户密码
func ValidatePassword(hashed string, userPassword string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

// 插入用户
func (u UserService) AddUser(user *model.User) (int64, error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

// 删除用户
func (u UserService) DeleteUser(userId int64) error {
	return u.UserRepository.DeleteUserByID(userId)
}

// 更新用户
func (u UserService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	//判断是否更新了密码
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u UserService) FindUserByName(name string) (*model.User, error) {
	return u.UserRepository.FindUserByName(name)
}
func (u UserService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(user.HashPassword, pwd)
}
