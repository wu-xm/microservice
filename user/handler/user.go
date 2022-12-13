package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"

	pb "user/proto/user"
)

type User struct {
	UserService service.IUserService
}

func (e *User) Register(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     request.UserName,
		FirstName:    request.FirstName,
		HashPassword: request.Pwd,
	}
	_, err := e.UserService.AddUser(userRegister)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (e *User) Login(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	isOk, err := e.UserService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

func (e *User) GetUserInfo(ctx context.Context, request *pb.UserInfoRequest, response *pb.UserInfoResponse) error {
	user, err := e.UserService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response.UserName = user.UserName
	response.FirstName = user.FirstName
	response.Pwd = user.HashPassword
	return nil
}

// Return a new handler
func New() *User {
	return &User{}
}
