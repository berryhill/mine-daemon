package services

import (
	"fmt"
	"log"

	"github.com/hpcloud/tail"
)

func StartLogs() {

	logs := make(chan string)

	t, err := tail.TailFile("./logs.txt", tail.Config{Follow: true})
	if err != nil {
		log.Println("Error tailing log", err)
	}

	go readLog(t, logs)
	go handleLog(logs)
}

func readLog(t *tail.Tail, logs chan string) {
	for line := range t.Lines {
		logs<- line.Text
	}
}

func handleLog(logs chan string) {
	for log := range logs {
		fmt.Println(log)
	}
}
