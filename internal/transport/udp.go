package transport

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

// UDPTransport is the UDP server implementation of DatagramTransport.
type UDPTransport struct {
	receiveDatagram func(dgram Datagram)
	sendChan        chan Datagram

	laddr *net.UDPAddr
	sock  *net.UDPConn

	runningMu sync.RWMutex
	running   bool
	wg        sync.WaitGroup
}

// NewUDPTransport creates a new instance of UDPTransport. The UDP server does
// not start listening until Start() is called.
// The address parameter is passed directly onto net.ResolveUDPAddr.
func NewUDPTransport(address string) (*UDPTransport, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	return &UDPTransport{laddr: addr}, err
}

func (tp *UDPTransport) isRunning() bool {
	tp.runningMu.RLock()
	defer tp.runningMu.RUnlock()
	return tp.running
}

// Start implements DatagramTransport.Start.
func (tp *UDPTransport) Start() error {
	tp.runningMu.Lock()
	defer tp.runningMu.Unlock()

	if tp.running {
		return errors.New("transport: attempt to start a transport that is already running")
	}
	sock, err := net.ListenUDP("udp", tp.laddr)
	if err != nil {
		return err
	}
	tp.sock = sock
	tp.sendChan = make(chan Datagram)
	tp.running = true

	tp.wg.Add(2)
	go tp.startReceive()
	go tp.startSend()
	return nil
}

func (tp *UDPTransport) startReceive() {
	buf := make([]byte, 2048)
	for {
		n, addr, err := tp.sock.ReadFromUDP(buf)
		if n > 0 {
			data := make([]byte, n)
			copy(data, buf[:n])
			go tp.processDatagram(Datagram{addr, data})
		}
		if err != nil {
			if !tp.isRunning() {
				break
			}
			// TODO: unexpected read error, what should we do?
			fmt.Println(err)
		}
	}
	tp.wg.Done()
}

func (tp *UDPTransport) startSend() {
	for dgram := range tp.sendChan {
		if err := tp.sendDatagram(dgram); err != nil {
			// TODO: unexpected write error, what should we do?
			fmt.Println(err)
		}
	}
	tp.wg.Done()
}

func (tp *UDPTransport) processDatagram(dgram Datagram) {
	if cb := tp.receiveDatagram; cb != nil {
		cb(dgram)
	}
}

func (tp *UDPTransport) sendDatagram(dgram Datagram) error {
	addr, err := net.ResolveUDPAddr(dgram.Addr.Network(), dgram.Addr.String())
	if err != nil {
		return err
	}
	_, err = tp.sock.WriteToUDP(dgram.Data, addr)
	return err
}

// Stop implements DatagramTransport.Stop.
func (tp *UDPTransport) Stop() error {
	if err := tp.close(); err != nil {
		return err
	}
	tp.wg.Wait()
	return nil
}

func (tp *UDPTransport) close() error {
	tp.runningMu.Lock()
	defer tp.runningMu.Unlock()

	if !tp.running {
		return errors.New("transport: attempt to stop a transport that is not running")
	}
	tp.running = false

	close(tp.sendChan)
	err := tp.sock.Close()
	return err
}

// SendDatagram implements DatagramTransport.SendDatagram.
func (tp *UDPTransport) SendDatagram(dgram Datagram) error {
	tp.sendChan <- dgram
	return nil
}

// SetReceiveDatagramFunc implements DatagramTransport.SetReceiveDatagramFunc.
func (tp *UDPTransport) SetReceiveDatagramFunc(f func(dgram Datagram)) {
	tp.receiveDatagram = f
}
