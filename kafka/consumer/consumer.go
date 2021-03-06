package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"os"
	"strings"
	"time"
)

var (
	topics = "update.mining.stage"
)

func main() {
	groupID := "group-1"
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	//创建消费者
	c, err := cluster.NewConsumer(strings.Split("127.0.0.1:9092", ","), groupID, strings.Split(topics, ","), config)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	defer c.Close()
	//subscriptions := c.Subscriptions()
	//fmt.Println(subscriptions)

	go func(c *cluster.Consumer) {
		errors := c.Errors()
		noti := c.Notifications()
		for {
			select {
			case err := <-errors:
				fmt.Println("err ", err)
			case <-noti:
				fmt.Printf("notify %+v \n", noti)
			}
		}
	}(c)

	for msg := range c.Messages() {
		fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}
