// Package transport is an interface for synchronous connection based communication
package transport

import (
	"time"
)

type TransportProtocol string

const (
	TCP  TransportProtocol = "tcp"
	KCP  TransportProtocol = "kcp"
	UDP  TransportProtocol = "udp"
	WS   TransportProtocol = "websocket"
	HTTP TransportProtocol = "http"
	QUIC TransportProtocol = "quic"
	RPC  TransportProtocol = "rpc"
)

// ITransport is an interface which is used for communication between
// services. It uses connection based socket send/recv semantics and
// has various implementations; http, grpc, quic.
type ITransport interface {
	Server() IServer
	Client() IClient
}

type Message struct {
	Header map[string]string
	Body   []byte
}

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

var (
	DefaultDialTimeout = time.Second * 5
)
