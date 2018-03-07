package services

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Conn 			*websocket.Conn
	Ping			chan *Ping
	Pong			chan *Ping
	LastPong 		*Ping
}

func NewHub(conn *websocket.Conn) *Hub {

	h := new(Hub)
	h.Conn = conn
	h.Ping = make(chan *Ping)
	h.Pong = make(chan *Ping)

	return h
}

func (h *Hub) Run() {

	go func () {
		for {
			_, message, err := h.Conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			pong := new(Ping)
			err = json.Unmarshal(message, pong)
			if err != nil {
				log.Printf("Error parsing pong", err)
			}
			h.Pong<- pong
			log.Printf("recv: %s", message)
		}
	}()

	go func() {
		for {
			select {
			case p := <-h.Ping:
				ping, err := json.Marshal(p)
				if err != nil {
					log.Println("error marshalling ping json", err)
				}
				err = h.Conn.WriteMessage(websocket.TextMessage, ping)
				if err != nil {
					log.Println("write:", err)
					return
				}
			case p := <-h.Pong:
				h.LastPong = p
			}
		}
	}()
}