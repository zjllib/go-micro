package micro

import (
	"context"
	"sync"
	"testing"

	proto "github.com/zjllib/go-micro/debug/proto"
	"github.com/zjllib/go-micro/registry"
	"github.com/zjllib/go-micro/util/test"
)

func TestFunction(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	r := registry.NewMemoryregistry.IRegistry.Services(test.Data))

	// create service
	fn := NewFunction(
		Registry(r),
		Name("test.function"),
		AfterStart(func() error {
			wg.Done()
			return nil
		}),
	)

	// we can't test fn.Init as it parses the command line
	// fn.Init()

	ch := make(chan error, 2)

	go func() {
		// run service
		ch <- fn.Run()
	}()

	// wait for start
	wg.Wait()

	// test call debug
	req := fn.Client().NewRequest(
		"test.function",
		"Debug.Health",
		new(proto.HealthRequest),
	)

	rsp := new(proto.HealthResponse)

	err := fn.Client().Call(context.TODO(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.Status != "ok" {
		t.Fatalf("function response: %s", rsp.Status)
	}

	if err := <-ch; err != nil {
		t.Fatal(err)
	}
}
