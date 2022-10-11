package main

import (
	"github.com/zjllib/go-micro/examples/v3/cache/handler"
	pb "github.com/zjllib/go-micro/examples/v3/cache/proto"

	"github.com/zjllib/go-micro/v3"
	log "github.com/zjllib/go-micro/v3/logger"
)

var (
	service = "go.micro.srv.cache"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterCacheHandler(srv.Server(), handler.NewCache())

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
