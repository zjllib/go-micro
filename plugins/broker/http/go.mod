module github.com/zjllib/go-micro/plugins/broker/http/v3

go 1.16

require (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/zjllib/go-micro v3.5.2-0.20210630062103-c13bb07171bc
	github.com/google/uuid v1.2.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
)

replace (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 => ../../../plugins/registry/memory
	github.com/zjllib/go-micro => ../../../../go-micro
)
