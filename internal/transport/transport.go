package transport

import "net"

// Datagram represents a chunk of data tagged with the address of its source
// (in the case of a received datagram) or its destination (in the case of a
// sent datagram).
type Datagram struct {
	Addr net.Addr
	Data []byte
}

// DatagramTransport defines the interface that any datagram protocol must
// implement in order to interact with the gateway.
type DatagramTransport interface {
	// Start starts the transport, enabling it to send and receive datagrams.
	Start() error

	// Stop stops the transport, making it unable to send and receive datagrams.
	Stop() error

	// SendDatagram is used to send some data to a specific address.
	SendDatagram(dgram Datagram) error

	// SetReceiveDatagramFunc sets the callback that is called when a datagram
	// is received from the network.
	SetReceiveDatagramFunc(func(dgram Datagram))
}
