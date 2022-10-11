package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/zjllib/go-micro"
	proto "github.com/zjllib/go-micro/examples/helloworld/proto"
	"github.com/zjllib/go-micro/examples/mocking/mock"
)

func main() {
	var c proto.GreeterService

	service := micro.NewService(
		micro.Flags(&cli.StringFlag{
			Name:  "environment",
			Value: "testing",
		}),
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			env := ctx.String("environment")
			// use the mock when in testing environment
			if env == "testing" {
				c = mock.NewGreeterService()
			} else {
				c = proto.NewGreeterService("helloworld", service.Client())
			}
			return nil
		}),
	)

	// call hello service
	rsp, err := c.Hello(context.TODO(), &proto.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
