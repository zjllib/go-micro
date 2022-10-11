package main

import (
	"context"
	"log"

	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/client"
	hello "github.com/zjllib/go-micro/examples/greeter/srv/proto/hello"
	"github.com/zjllib/go-micro/registry"
)

func main() {
	CallGrpcByHttp()
}

func CallGrpcByHttp() {
	// create a new service
	service := micro.NewService(
		micro.Registry(registry.NewRegistry()),
	)
	// parse command line flags
	service.Init()
	c := service.Client()

	req := &hello.Request{Name: "call grpc server by http client"}
	// create request/response
	request := client.NewRequest("go.micro.srv.greeter", "Say.Hello", req)

	response := new(hello.Response)
	// call service
	err := c.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)
}
