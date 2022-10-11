package main

import (
	"github.com/zjllib/go-micro/examples/cache/handler"
	pb "github.com/zjllib/go-micro/examples/cache/proto"

	"github.com/zjllib/go-micro"
	log "github.com/zjllib/go-micro/logger"
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
