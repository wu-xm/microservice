package main

import (
	"cart/common"
	"cart/domain/repository"
	"cart/domain/service"
	"cart/handler"
	pb "cart/proto"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracingv3 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

//	{
//		"host":"127.0.0.1",
//		"user":"root",
//		"pwd":"123456",
//		"database":"micro",
//		"port":3306
//	}
const QPS = 100

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
	//链路追踪
	t, io, err := common.NewTracer("cart", "localhost:6831")
	if err != nil {
		log.Println(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// Create service
	srv := micro.NewService(
		micro.Name("cart"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8087"),
		//添加consul作为注册中心
		micro.Registry(consulRegistry),
		//绑定链路追踪
		micro.WrapHandler(opentracingv3.NewHandlerWrapper(opentracing.GlobalTracer())),
		//限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
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
	cartRepository := repository.NewCartRepository(db)
	//建表
	if err = cartRepository.InitTable(); err != nil {
		fmt.Println("数据库建表", err)
	}
	// Register handler
	productService := service.NewCartDataService(cartRepository)

	err = pb.RegisterCartHandler(srv.Server(), &handler.Cart{CartDataService: productService})
	if err != nil {
		fmt.Println("micro服务注册", err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println("micro服务运行", err)
	}
}
