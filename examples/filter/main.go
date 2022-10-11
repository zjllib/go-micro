package main

import (
	"context"
	"fmt"

	"github.com/zjllib/go-micro"
	"github.com/zjllib/go-micro/examples/filter/version"
	proto "github.com/zjllib/go-micro/examples/service/proto"
)

func main() {
	service := micro.NewService()
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(
		// provide a context
		context.TODO(),
		// provide the request
		&proto.Request{Name: "John"},
		// set the filter
		version.Filter("latest"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
