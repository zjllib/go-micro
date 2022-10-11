module github.com/zjllib/go-micro/plugins/broker/gocloud/v3

go 1.16

require (
	github.com/zjllib/go-micro v3.5.2-0.20210630062103-c13bb07171bc
	github.com/streadway/amqp v1.0.0
	gocloud.dev v0.17.0
	gocloud.dev/pubsub/rabbitpubsub v0.17.0
)

replace github.com/zjllib/go-micro => ../../../../go-micro