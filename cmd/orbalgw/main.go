package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kabbah/orbalgw/internal/gateway"
	"github.com/Kabbah/orbalgw/internal/transport"
)

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	tp, err := transport.NewUDPTransport(":11037")
	if err != nil {
		log.Panicln(err)
	}

	gw := gateway.New(tp)
	if err = gw.Start(); err != nil {
		log.Panicln(err)
	}

	<-sigChan

	if err = gw.Stop(); err != nil {
		log.Panicln(err)
	}
}
