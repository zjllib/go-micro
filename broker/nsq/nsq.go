// Package nsq provides an NSQ broker
package nsq

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nsqio/go-nsq"
	"github.com/zjllib/go-micro/broker"
	"github.com/zjllib/go-micro/codec/json"
)

type nsqBroker struct {
	lookupdAddrs []string
	addrs        []string
	opts         broker.Options
	config       *nsq.Config

	sync.Mutex
	running bool
	p       []*nsq.Producer
	c       []*subscriber
}

type subscriber struct {
	topic string
	opts  broker.SubscribeOptions

	c *nsq.Consumer

	// handler so we can resubcribe
	h nsq.HandlerFunc
	// concurrency
	n int
}

var (
	DefaultConcurrentHandlers = 1
)

func (n *nsqBroker) Init(opts ...broker.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}

	var addrs []string

	for _, addr := range n.opts.Addrs {
		if len(addr) > 0 {
			addrs = append(addrs, addr)
		}
	}

	if len(addrs) == 0 {
		addrs = []string{"127.0.0.1:4150"}
	}

	n.addrs = addrs
	return nil
}

func (n *nsqBroker) Connect() error {
	n.Lock()
	defer n.Unlock()

	if n.running {
		return nil
	}

	// create producers
	if n.opts.Side&broker.SIDE_PRODUCER >= 1 {
		producers := make([]*nsq.Producer, 0, len(n.addrs))
		for _, addr := range n.addrs {
			p, err := nsq.NewProducer(addr, n.config)
			if err != nil {
				return err
			}
			if err = p.Ping(); err != nil {
				return err
			}
			producers = append(producers, p)
		}
		n.p = producers
	}
	// create consumers
	//if n.opts.Side&broker.SIDE_CONSUMER >= 1 {
	//	for _, c := range n.c {
	//		channel := c.opts.Queue
	//		if len(channel) == 0 {
	//			channel = uuid.New().String() + "#ephemeral"
	//		}
	//
	//		cm, err := nsq.NewConsumer(c.topic, channel, n.config)
	//		if err != nil {
	//			return err
	//		}
	//
	//		cm.AddConcurrentHandlers(c.h, c.n)
	//
	//		c.c = cm
	//
	//		if len(n.lookupdAddrs) > 0 {
	//			c.c.ConnectToNSQLookupds(n.lookupdAddrs)
	//		} else {
	//			err = c.c.ConnectToNSQDs(n.addrs)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//	}
	//}
	n.running = true
	return nil
}

func (n *nsqBroker) Disconnect() error {
	n.Lock()
	defer n.Unlock()

	if !n.running {
		return nil
	}

	// stop the producers
	for _, p := range n.p {
		p.Stop()
	}

	// stop the consumers
	for _, c := range n.c {
		c.c.Stop()

		if len(n.lookupdAddrs) > 0 {
			// disconnect from all lookupd
			for _, addr := range n.lookupdAddrs {
				c.c.DisconnectFromNSQLookupd(addr)
			}
		} else {
			// disconnect from all nsq brokers
			for _, addr := range n.addrs {
				c.c.DisconnectFromNSQD(addr)
			}
		}
	}

	n.p = nil
	n.running = false
	return nil
}

func (n *nsqBroker) Publish(topic string, body []byte, opts ...broker.PublishOption) error {
	p := n.p[rand.Intn(len(n.p))]

	options := broker.PublishOptions{}
	for _, o := range opts {
		o(&options)
	}

	var (
		doneChan chan *nsq.ProducerTransaction
		delay    time.Duration
	)

	//b, err := n.opts.Codec.Marshal(message)
	//if err != nil {
	//	return err
	//}

	if doneChan != nil {
		if delay > 0 {
			return p.DeferredPublishAsync(topic, delay, body, doneChan)
		}
		return p.PublishAsync(topic, body, doneChan)
	} else {
		if delay > 0 {
			return p.DeferredPublish(topic, delay, body)
		}
		return p.Publish(topic, body)
	}
}

func (n *nsqBroker) Subscribe(topic string, handler broker.Handler, opts ...broker.SubscribeOption) error {
	options := broker.SubscribeOptions{
		AutoAck: true,
	}

	for _, o := range opts {
		o(&options)
	}

	concurrency, maxInFlight := DefaultConcurrentHandlers, DefaultConcurrentHandlers

	channel := options.Queue
	if len(channel) == 0 {
		channel = uuid.New().String() + "#ephemeral"
	}
	config := *n.config
	config.MaxInFlight = maxInFlight

	c, err := nsq.NewConsumer(topic, channel, &config)
	if err != nil {
		return err
	}

	h := nsq.HandlerFunc(func(nm *nsq.Message) error {
		if !options.AutoAck {
			nm.DisableAutoResponse()
		}
		return handler(message{
			nm:    nm,
			codec: n.opts.Codec,
		})
	})

	c.AddConcurrentHandlers(h, concurrency)

	if len(n.lookupdAddrs) > 0 {
		err = c.ConnectToNSQLookupds(n.lookupdAddrs)
	} else {
		err = c.ConnectToNSQDs(n.addrs)
	}
	if err != nil {
		return err
	}

	return nil
}

func NewBroker(opts ...broker.Option) broker.IBroker {
	options := broker.Options{
		// Default codec
		Codec: json.Marshaler{},
		// Default context
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	var addrs []string

	for _, addr := range options.Addrs {
		if len(addr) > 0 {
			addrs = append(addrs, addr)
		}
	}

	if len(addrs) == 0 {
		addrs = []string{"127.0.0.1:4150"}
	}

	n := &nsqBroker{
		addrs:  addrs,
		opts:   options,
		config: nsq.NewConfig(),
	}

	return n
}
