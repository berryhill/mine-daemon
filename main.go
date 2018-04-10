package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/berryhill/mine-daemon/services"
)

var addr = flag.String(
	"addr",
	"amqp://guest:guest@localhost:5672/",
	"http service address",
	)

var id = "1"

func main () {

	flag.Parse()

	fmt.Println(*addr)

	services.StartPing(*addr, id)
	services.StartLogs(*addr, id)

	fmt.Println("Up and running")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	doneCh := make(chan struct{})
	go func() {
		select {
		case <-signals:
			doneCh <- struct{}{}
		}
	}()
	<-doneCh
}
