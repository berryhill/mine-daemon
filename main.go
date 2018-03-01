package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/berryhill/mine-daemon/services"

	"github.com/berryhill/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var id = "1234"

func main () {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("error dialing websocket: ", err)
	}
	defer c.Close()

	t := time.NewTicker(time.Second)
	defer t.Stop()
	go services.StartPing(c, t, id)

	go services.StartLogs()

	fmt.Println("Up and running")
	for {}
}
