package nats_test

import (
	"os"
	"testing"

	log "github.com/zjllib/go-micro/logger"
	"github.com/zjllib/go-micro/plugins/registry/nats"
	"github.com/zjllib/go-micro/registry"
)

type environment struct {
	registryOne   registry.IRegistry
	registryTwo   registry.IRegistry
	registryThree registry.IRegistry

	serviceOne registry.Service
	serviceTwo registry.Service

	nodeOne   registry.Node
	nodeTwo   registry.Node
	nodeThree registry.Node
}

var e environment

func TestMain(m *testing.M) {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		log.Infof("NATS_URL is undefined - skipping tests")
		return
	}

	e.registryOne = nats.Newregistry.IRegistry.Addrs(natsURL), nats.Quorum(1))
	e.registryTwo = nats.Newregistry.IRegistry.Addrs(natsURL), nats.Quorum(1))
	e.registryThree = nats.Newregistry.IRegistry.Addrs(natsURL), nats.Quorum(1))

	e.serviceOne.Name = "one"
	e.serviceOne.Version = "default"
	e.serviceOne.Nodes = []*registry.Node{&e.nodeOne}

	e.serviceTwo.Name = "two"
	e.serviceTwo.Version = "default"
	e.serviceTwo.Nodes = []*registry.Node{&e.nodeOne, &e.nodeTwo}

	e.nodeOne.Id = "one"
	e.nodeTwo.Id = "two"
	e.nodeThree.Id = "three"

	if err := e.registryOne.Register(&e.serviceOne); err != nil {
		log.Fatal(err)
	}

	if err := e.registryOne.Register(&e.serviceTwo); err != nil {
		log.Fatal(err)
	}

	result := m.Run()

	if err := e.registryOne.Deregister(&e.serviceOne); err != nil {
		log.Fatal(err)
	}

	if err := e.registryOne.Deregister(&e.serviceTwo); err != nil {
		log.Fatal(err)
	}

	os.Exit(result)
}
