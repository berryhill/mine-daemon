package services

import (
	"time"
)

type Ping struct {
	Id 			string			`json:"id"`
	Service 	string			`json:"service"`
	Time 		time.Time		`json:"time"`
}

func NewPing(id string) *Ping {

	p := new(Ping)
	p.Id = id
	p.Service = "ping"
	p.Time = time.Now()

	return p
}

func StartPing(hub *Hub, ticker *time.Ticker, id string) {

	go func () {
		for t := range ticker.C {
			if t.String() != "" {}
			ping := NewPing(id)
			hub.Ping<- ping
		}
	}()
}
