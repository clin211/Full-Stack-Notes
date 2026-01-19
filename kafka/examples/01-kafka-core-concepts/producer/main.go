package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/forest/full-stack-notes/kafka/examples/01-kafka-core-concepts/config"
)

func main() {
	// 加载配置
	cfg := config.DefaultProducerConfig()

	// 创建生产者
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.BootstrapServers,
		"acks":               cfg.Acks,
		"compression.type":   cfg.CompressionType,
		"linger.ms":          cfg.LingerMs,
		"batch.size":         cfg.BatchSize,
		"retries":            cfg.RetryTimes,
		"request.timeout.ms": int(cfg.Timeout.Milliseconds()),
		"enable.idempotence": true, // 启用幂等性
	})
	if err != nil {
		fmt.Printf("创建生产者失败: %v\n", err)
		os.Exit(1)
	}
	defer p.Close()

	fmt.Printf("Kafka 生产者已启动，连接到: %s\n", cfg.BootstrapServers)
	fmt.Printf("Topic: %s\n", cfg.Topic)
	fmt.Printf("Acks: %s, Compression: %s\n", cfg.Acks, cfg.CompressionType)

	// 创建 delivery channel 用于接收发送结果
	deliveryChan := make(chan kafka.Event, 10000)

	// 处理信号，优雅退出
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 示例 1：发送简单消息
	go func() {
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("Hello Kafka! Message #%d from Go", i)

			err := p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &cfg.Topic,
					Partition: kafka.PartitionAny,
				},
				Key:   []byte(key),
				Value: []byte(value),
			}, deliveryChan)

			if err != nil {
				fmt.Printf("发送失败: %v\n", err)
			}
		}
	}()

	// 示例 2：发送 JSON 消息
	go func() {
		user := map[string]interface{}{
			"id":        "user-123",
			"name":      "张三",
			"email":     "zhangsan@example.com",
			"age":       25,
			"action":    "login",
			"timestamp": "2024-01-15T10:30:00Z",
		}

		// 简单的 JSON 序列化（生产环境建议使用 encoding/json）
		jsonStr := fmt.Sprintf(
			`{"id":"%s","name":"%s","email":"%s","age":%d,"action":"%s","timestamp":"%s"}`,
			user["id"], user["name"], user["email"], user["age"],
			user["action"], user["timestamp"],
		)

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &cfg.Topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(user["id"].(string)),
			Value: []byte(jsonStr),
		}, deliveryChan)
	}()

	// 处理发送结果
	run := true
	remaining := 11
	for run && remaining > 0 {
		select {
		case <-sigchan:
			fmt.Println("\n收到退出信号，正在停止...")
			run = false

		case ev := <-deliveryChan:
			switch e := ev.(type) {
			case *kafka.Message:
				if e.TopicPartition.Error != nil {
					fmt.Printf("发送失败 [%s]: %v\n",
						e.TopicPartition, e.TopicPartition.Error)
				} else {
					fmt.Printf("发送成功: Topic=%s, Partition=%d, Offset=%d, Key=%s\n",
						*e.TopicPartition.Topic,
						e.TopicPartition.Partition,
						e.TopicPartition.Offset,
						string(e.Key))
				}
				remaining--
			}
		}
	}

	// 等待所有消息发送完成
	remaining = p.Flush(15 * 1000)
	if remaining > 0 {
		fmt.Printf("还有 %d 条消息未发送完成\n", remaining)
	}

	close(deliveryChan)
	fmt.Println("生产者已关闭")
}
