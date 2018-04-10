package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hpcloud/tail"
	"github.com/streadway/amqp"
)

func StartLogs(url, id string) {

	lineChan := make(chan string)

	go func() {
		t, err := tail.TailFile(
			"./logs.txt", tail.Config{Follow: true})
		if err != nil {
			log.Println("Error tailing log", err)
		}
		for line := range t.Lines {
			fmt.Println(line.Text)
			lineChan<- line.Text
		}
	}()

	go func() {
		conn, err := amqp.Dial(url)
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"logs",
			false,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "Failed to declare a queue")

		for {
			message := NewMessage(id, "logs",  <-lineChan)
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
		}
	}()
}
