package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	opentracingv3 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/opentracing/opentracing-go"
	"log"
	"product/common"
	product "product/proto"
)

func main() {
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	t, io, err := common.NewTracer("product.client", "localhost:6831")
	if err != nil {
		log.Println(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	srv := micro.NewService(
		micro.Name("categoryClient"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul作为注册中心
		micro.Registry(consulRegistry),
		//绑定链路追踪
		micro.WrapClient(opentracingv3.NewClientWrapper(opentracing.GlobalTracer())),
	)
	srv.Init()
	productService := product.NewProductService("category", srv.Client())
	addProduct, err := productService.AddProduct(context.Background(), &product.ProductInfo{ProductCategoryId: 10})
	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(addProduct)
}
