package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)
func main() {
	producer, err := nsq.NewProducer("111.62.107.178:4150", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewProducer", err)
		panic(err)
	}
	var message string
	fmt.Scanln(&message)
	if err := producer.Publish("test", []byte(fmt.Sprintf(message))); err != nil {
		fmt.Println("Publish", err)
		panic(err)
	}
}
