package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// ProduceWithHeaders 演示如何发送带 Headers 的消息
func ProduceWithHeaders() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		fmt.Printf("创建生产者失败: %v\n", err)
		return
	}
	defer p.Close()

	topic := "quickstart-events"

	headers := []kafka.Header{
		{Key: "content-type", Value: []byte("application/json")},
		{Key: "producer-id", Value: []byte("producer-1")},
		{Key: "trace-id", Value: []byte("abc-123-def")},
		{Key: "timestamp", Value: []byte(time.Now().Format(time.RFC3339))},
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:     []byte("message-with-headers"),
		Value:   []byte(`{"status":"success","data":"test"}`),
		Headers: headers,
	}, nil)

	if err != nil {
		fmt.Printf("发送失败: %v\n", err)
		return
	}

	// 等待消息发送完成
	p.Flush(15 * 1000)
	fmt.Println("带 Headers 的消息已发送")
}
