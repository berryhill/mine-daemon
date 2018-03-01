package services

import (
	"log"
	"time"

	"github.com/berryhill/websocket"
	"encoding/json"
)

type Ping struct {
	Id 			string			`json:"id"`
	Service 	string			`json:"service"`
	Time 		time.Time		`json:"time"`
	Message 	string 			`json:"message"`
}

func NewPing(id string) *Ping {
	p := new(Ping)
	p.Id = id
	p.Service = "Health"
	p.Time = time.Now()
	p.Message = "Hey friend"

	return p
}

func StartPing(c *websocket.Conn, ticker *time.Ticker, id string) {
	go func () {
		for t := range ticker.C {
			if t.String() != "" {}

			ping := NewPing(id)
			pingJson, err := json.Marshal(ping)
			if err != nil {
				log.Println("error marshalling ping json", err)
			}

			err = c.WriteMessage(websocket.TextMessage, pingJson)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}()
}
