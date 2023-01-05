package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/zjllib/go-micro/codec"
	"time"
)

type message struct {
	nm    *nsq.Message
	codec codec.Marshaler
}

func (m message) Unmarshal(obj interface{}) error {
	return m.codec.Unmarshal(m.nm.Body, obj)
}

func (m message) Finish() error {
	m.nm.Finish()
	return nil
}

func (m message) Requeue() error {
	m.nm.Requeue(time.Second)
	return nil
}

func (m message) Body() []byte {
	return m.nm.Body
}

func (m message) UUID() string {
	return fmt.Sprintf("%s", m.nm.ID)
}
