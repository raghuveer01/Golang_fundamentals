package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	broker = "localhost:9093"
	group  = "my-group"
	topics = "myTopic"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := topics
	for i := 0; i < 100; i++ {
		dataSent := []byte(strconv.Itoa(i))
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          dataSent,
		}, nil)
		time.Sleep(time.Second)
		fmt.Println(dataSent)
	}

	// Wait for message deliveries before shutting down
	p.Flush(5 * 1000)
}
