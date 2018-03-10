package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/berryhill/mine-daemon/services"
)

var Addr = flag.String(
	"addr", "localhost:5050", "http service address")
var id = "100"

func main () {

	brokers := []string{"35.193.166.194:9092"}
	go services.StartPing(brokers, id)

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
