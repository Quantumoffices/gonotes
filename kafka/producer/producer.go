package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

//上报挖矿阶段
type UpMiningStageReq struct {
	Stage int `json:"stage"` //当前挖矿阶段
}

//想要更高的吞吐量就设置：异步、ack=0；
//想要不丢失消息数据就选：同步、ack=-1策略
//https://blog.csdn.net/moonpure/article/details/86541480
func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	//Kafka保证消息被安全生产，有三个选项分别是0,1,-1
	//通过request.required.acks属性进行配置：
	//0代表：不进行消息接收是否成功的确认(默认值)；
	//1代表：当Leader副本接收成功后，返回接收成功确认信息；
	//-1代表：当Leader和Follower副本都接收成功后，返回接收成功确认信息；
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//Kafka消息发送分同步(sync)、异步(async)两种方式
	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for i := 1; i < 20; i++ {
		msg := &sarama.ProducerMessage{
			Topic:     "update.mining.stage",
			Partition: int32(1),
			Key:       sarama.StringEncoder("key"),
		}
		//var value string
		req := UpMiningStageReq{Stage: i}
		bytes, _ := json.Marshal(req)
		fmt.Println(string(bytes))
		msg.Value = sarama.ByteEncoder(bytes)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Send Message Fail!")
		}
		fmt.Printf("Partion = %d, offset = %d\n", partition, offset)

		time.Sleep(time.Second * 10)
	}

}
