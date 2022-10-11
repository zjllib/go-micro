module github.com/zjllib/go-micro/plugins/wrapper/breaker/gobreaker/v3

go 1.16

require (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/zjllib/go-micro/v3 v3.5.2-0.20210630062103-c13bb07171bc
	github.com/sony/gobreaker v0.4.1
)

replace (
	github.com/zjllib/go-micro/plugins/registry/memory/v3 => ../../../../plugins/registry/memory
	github.com/zjllib/go-micro/v3 => ../../../../../go-micro
)
