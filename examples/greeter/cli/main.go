package main

import (
	"context"
	"fmt"
	"github.com/zjllib/go-micro/registry"
	"github.com/zjllib/go-micro/selector"
	"math/rand"
	"sync"
	"time"

	"github.com/zjllib/go-micro"
	hello "github.com/zjllib/go-micro/examples/greeter/srv/proto/hello"
)


// Built in random hashed node selector
type dcSelector struct {
	opts selector.Options
}

var (
	datacenter = "local"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (n *dcSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}
	return nil
}

func (n *dcSelector) Options() selector.Options {
	return n.opts
}
var i int
func (n *dcSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	services, err := n.opts.Registry.GetService(service)
	if err != nil {
		return nil, err
	}



	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var nodes []*registry.Node

	// Filter the nodes for datacenter
	for _, service := range services {
		for _, node := range service.Nodes {
			nodes = append(nodes, node)
		}
	}

	if len(nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	var i int
	var mtx sync.Mutex

	return func() (*registry.Node, error) {
		mtx.Lock()
		defer mtx.Unlock()
		i++
		return nodes[1], nil
	}, nil
}

func (n *dcSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (n *dcSelector) Reset(service string) {
	return
}

func (n *dcSelector) Close() error {
	return nil
}

func (n *dcSelector) String() string {
	return "dc"
}
// Return a new first node selector
func DCSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	if sopts.Registry == nil {
		sopts.Registry = registry.DefaultRegistry
	}
	return &dcSelector{sopts}
}
func main() {
	// create a new service
	service := micro.NewService(
		micro.
	/*	micro.Selector(m)*/)

	// parse command line flags
	service.Init()

	micro.

	// Use the generated client stub
	cl := hello.NewSayService("go.micro.greeter", service.Client())

	for i := 0; i < 10; i++ {
		// Make request
		rsp, err := cl.Hello(context.Background(), &hello.Request{
			Name: "John ---",
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp.Msg)
	}

}
