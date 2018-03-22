package services

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Ping struct {
	Time 		time.Time		`json:"time"`
}

func NewPing() *Ping {

	p := new(Ping)
	p.Time = time.Now()

	return p
}

func StartPing(url, id string) {

	go func() {
		conn, err := amqp.Dial(url)
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"ping",
			false,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "Failed to declare a queue")

		for {
			ping := NewPing()
			message := NewMessage(id, "ping", ping)
			body, _ := json.Marshal(message)
			err = ch.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			log.Printf(" [x] Sent %s", body)
			failOnError(err, "Failed to publish a message")
			time.Sleep(time.Second * 2)
		}
	}()
}
