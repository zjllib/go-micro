module github.com/zjllib/go-micro/plugins/wrapper/trace/datadog/v3

go 1.16

require (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/zjllib/go-micro v3.5.2-0.20210630062103-c13bb07171bc
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.38.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.31.1
)

replace (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 => ../../../../plugins/registry/memory
	github.com/zjllib/go-micro => ../../../../../go-micro
)