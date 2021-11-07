package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Handler struct{}

func (m *Handler) HandleMessage(msg *nsq.Message) (err error) {
	addr := msg.NSQDAddress
	message := string(msg.Body)
	fmt.Println(addr, message)
	return
}
func NewConsumers(t string, c string, addr string) error {
	conf := nsq.NewConfig()
	nc, err := nsq.NewConsumer(t, c, conf)
	if err != nil {
		fmt.Println("create consumer failed err ", err)
		return err
	}
	consumer := &Handler{}
	nc.AddHandler(consumer)
	// 连接nsqlookupd
	if err:= nc.ConnectToNSQLookupd(addr);err!=nil{
		fmt.Println("connect nsqlookupd failed ", err)
		return err
	}
	return nil
}
func main() {
	// 这是nsqlookupd的地址
	addr := "111.62.107.178:4161"
	err := NewConsumers("test", "channel-aa", addr)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}
	select {}
}