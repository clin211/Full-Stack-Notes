package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// ConsumeWithManualCommit 演示如何手动提交消费位移
func ConsumeWithManualCommit() {
	// 禁用自动提交
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "manual-commit-group",
		"auto.offset.reset": "earliest",
		"enable.auto.commit": "false", // 禁用自动提交
	})
	if err != nil {
		fmt.Printf("创建消费者失败: %v\n", err)
		return
	}
	defer c.Close()

	err = c.SubscribeTopics([]string{"quickstart-events"}, nil)
	if err != nil {
		fmt.Printf("订阅 Topic 失败: %v\n", err)
		return
	}

	fmt.Println("消费者已启动，使用手动提交模式")
	fmt.Println("消息 'error' 将不会提交位移，下次会重新消费")

	for {
		msg, err := c.ReadMessage(100 * time.Millisecond)
		if err != nil {
			// 忽略超时错误
			if err.(kafka.Error).Code() == kafka.ErrTimedOut {
				continue
			}
			fmt.Printf("读取消息错误: %v\n", err)
			continue
		}

		// 处理消息
		fmt.Printf("处理消息: %s\n", string(msg.Value))

		// 模拟处理失败
		if string(msg.Value) == "error" {
			fmt.Println("处理失败，不提交位移")
			continue // 不提交，下次重新消费
		}

		// 处理成功，提交位移
		// 方式 1：同步提交当前消息的位移
		_, err = c.CommitMessage(msg)
		if err != nil {
			fmt.Printf("提交失败: %v\n", err)
		} else {
			fmt.Printf("位移已提交: Partition=%d, Offset=%d\n",
				msg.TopicPartition.Partition, msg.TopicPartition.Offset)
		}
	}
}
