package mock

import (
	"context"

	"github.com/zjllib/go-micro/client"
	proto "github.com/zjllib/go-micro/examples/helloworld/proto"
)

type mockGreeterService struct {
}

func (m *mockGreeterService) Hello(ctx context.Context, req *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
	return &proto.Response{
		Greeting: "Hello " + req.Name,
	}, nil
}

func NewGreeterService() proto.GreeterService {
	return new(mockGreeterService)
}
