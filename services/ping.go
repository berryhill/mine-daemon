package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

type Ping struct {
	Time 		time.Time		`json:"time"`
}

func NewPing() *Ping {

	p := new(Ping)
	p.Time = time.Now()

	return p
}

func StartPing(brokers []string, id string) {

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, errors int
	doneCh := make(chan struct{})

	go func() {
		for {

			time.Sleep(4 * time.Second)

			ping := NewPing()
			pingPayload, err := json.Marshal(ping)
			if err != nil {
				log.Printf("Error marshalling json", err)
			}

			msg := &sarama.ProducerMessage{
				Topic: "ping",
				Key:   sarama.ByteEncoder(id),
				Value: sarama.ByteEncoder(pingPayload),
			}
			select {
			case producer.Input() <- msg:
				enqueued++
			case err := <-producer.Errors():
				errors++
				fmt.Println("Failed to produce message:", err)
			case <-signals:
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
