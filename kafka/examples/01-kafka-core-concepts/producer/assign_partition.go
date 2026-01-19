package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// ProduceToSpecificPartition 演示如何指定分区发送消息
func ProduceToSpecificPartition() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		fmt.Printf("创建生产者失败: %v\n", err)
		return
	}
	defer p.Close()

	topic := "quickstart-events"

	// 方式 1：直接指定分区号
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: 0, // 指定分区 0
		},
		Key:   []byte("explicit-partition"),
		Value: []byte("This message goes to partition 0"),
	}, nil)
	if err != nil {
		fmt.Printf("发送失败: %v\n", err)
		return
	}

	// 方式 2：让 Kafka 自动选择（基于 Key 的哈希）
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny, // 自动选择
		},
		Key:   []byte("user-123"), // 相同 Key 总是分配到同一分区
		Value: []byte("User 123 data"),
	}, nil)
	if err != nil {
		fmt.Printf("发送失败: %v\n", err)
		return
	}

	// 等待消息发送完成
	p.Flush(15 * 1000)
	fmt.Println("指定分区的消息已发送")
}
