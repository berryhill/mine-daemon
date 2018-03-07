package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/berryhill/mine-daemon/services"

	"github.com/gorilla/websocket"
)

var Addr = flag.String(
	"addr", "localhost:8080", "http service address")
var id = "1234"

func main () {

	u := url.URL{Scheme: "ws", Host: *Addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("error dialing websocket: ", err)
	}
	defer c.Close()

	t := time.NewTicker(time.Second)
	defer t.Stop()

	hub := services.NewHub(c)
	hub.Run()
	go services.StartPing(hub, t, id)
	go services.StartLogs()

	fmt.Println("Up and running")
	for { /* do nothing */ }
}
