module github.com/zjllib/go-micro/plugins/wrapper/trace/opencensus/v3

go 1.16

require (
	github.com/zjllib/go-micro v3.5.2-0.20210630062103-c13bb07171bc
	go.opencensus.io v0.23.0
	google.golang.org/genproto v0.0.0-20210624195500-8bfb893ecb84
)

replace github.com/zjllib/go-micro => ../../../../../go-micro
