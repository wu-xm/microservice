package main

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/domain/repository"
	"user/domain/service"
	"user/handler"
	pb "user/proto/user"
)

func main() {

	//数据库初始化
	//创建服务
	svc := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)
	//服务初始化
	svc.Init()
	//数据库连接
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接", err)
	}
	userRepository := repository.NewUserRepository(db)
	//建表
	if err = userRepository.InitTable(); err != nil {
		fmt.Println("数据库建表", err)
	}
	//服务注册
	userService := service.NewUserService(userRepository)
	if err = pb.RegisterUserHandler(svc.Server(), &handler.User{UserService: userService}); err != nil {
		fmt.Println("micro服务注册", err)
	}
	//服务运行
	if err = svc.Run(); err != nil {
		fmt.Println("micro服务运行", err)
	}

}
