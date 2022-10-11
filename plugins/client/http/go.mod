module github.com/zjllib/go-micro/plugins/client/http/v3

go 1.16

require (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/zjllib/go-micro v3.5.2-0.20210630062103-c13bb07171bc
	github.com/golang/protobuf v1.5.2
)

replace (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 => ../../../plugins/registry/memory
	github.com/zjllib/go-micro => ../../../../go-micro
)
