package main

import (
	"category/common"
	"category/domain/repository"
	"category/domain/service"
	"category/handler"
	pb "category/proto"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

//	{
//		"host":"127.0.0.1",
//		"user":"root",
//		"pwd":"123456",
//		"database":"micro",
//		"port":3306
//	}
func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		fmt.Println("配置中心", err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// Create service
	srv := micro.NewService(
		micro.Name("category"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul作为注册中心
		micro.Registry(consulRegistry),
	)
	srv.Init()
	//获取mysql配置，
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	//数据库连接
	db, err := gorm.Open(mysql.Open(mysqlInfo.User+":"+
		mysqlInfo.Pwd+"@tcp("+
		mysqlInfo.Host+":"+
		strconv.FormatInt(mysqlInfo.Port, 10)+")/"+
		mysqlInfo.Database+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接", err)
	}
	categoryRepository := repository.NewCategoryRepository(db)
	//建表
	if err = categoryRepository.InitTable(); err != nil {
		fmt.Println("数据库建表", err)
	}
	// Register handler
	categoryService := service.NewCategoryService(categoryRepository)

	err = pb.RegisterCategoryHandler(srv.Server(), &handler.Category{CategoryService: categoryService})
	if err != nil {
		fmt.Println("micro服务注册", err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println("micro服务运行", err)
	}
}
