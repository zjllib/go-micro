// Package broker is an interface used for asynchronous messaging
package broker

// IBroker is an interface used for asynchronous messaging.
type IBroker interface {
	Init(...Option) error
	Connect() error
	Disconnect() error
	Publish(topic string, body []byte, opts ...PublishOption) error
	Subscribe(topic string, h Handler, opts ...SubscribeOption) error
}

type IMessage interface {
	UUID() string
	Unmarshal(interface{}) error
	Finish() error
	Requeue() error
	Body() []byte
}

// Handler is used to process messages via a subscription of a topic.
// The handler is passed a publication interface which contains the
// message and optional Ack method to acknowledge receipt of the message.
type Handler func(IMessage) error
