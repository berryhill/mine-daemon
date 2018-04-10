package services
//
//import (
//	"fmt"
//	"log"
//	"os"
//	"os/signal"
//
//	"github.com/Shopify/sarama"
//)
//
//type Command struct {
//	Action 		string		`json:"name"`
//}
//
//func StartCommander(brokers []string, id string) {
//
//	config := sarama.NewConfig()
//	config.Consumer.Return.Errors = true
//
//	master, err := sarama.NewConsumer(brokers, config)
//	if err != nil {
//		panic(err)
//	}
//
//	defer func() {
//		if err := master.Close(); err != nil {
//			panic(err)
//		}
//	}()
//
//	topic := "commands"
//	consumer, err := master.ConsumePartition(
//		topic, 0, sarama.OffsetNewest)
//	if err != nil {
//		panic(err)
//	}
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, os.Interrupt)
//
//	msgCount := 0
//
//	doneCh := make(chan struct{})
//	go func() {
//		for {
//			select {
//			case err := <-consumer.Errors():
//				log.Printf("Error parsing kafka message", err)
//			case msg := <-consumer.Messages():
//				if string(msg.Value) != id {
//					break
//				}
//				msgCount++
//				fmt.Println(string(msg.Key), string(msg.Value))
//			case <-signals:
//				fmt.Println("Interrupt is detected")
//				doneCh <- struct{}{}
//			}
//		}
//	}()
//
//	<-doneCh
//	fmt.Println("Processed", msgCount, "messages")
//}
