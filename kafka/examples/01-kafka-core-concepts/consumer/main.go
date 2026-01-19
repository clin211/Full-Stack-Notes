package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/forest/full-stack-notes/kafka/examples/01-kafka-core-concepts/config"
)

func main() {
	// 加载配置
	cfg := config.DefaultConsumerConfig()

	// 创建消费者
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg.BootstrapServers,
		"group.id":                 cfg.GroupID,
		"auto.offset.reset":        cfg.AutoOffsetReset,
		"enable.auto.commit":       cfg.EnableAutoCommit,
		"auto.commit.interval.ms":  cfg.AutoCommitIntervalMs,
		"session.timeout.ms":       cfg.SessionTimeoutMs,
		"heartbeat.interval.ms":    cfg.HeartbeatIntervalMs,
	})
	if err != nil {
		fmt.Printf("创建消费者失败: %v\n", err)
		os.Exit(1)
	}
	defer c.Close()

	fmt.Printf("Kafka 消费者已启动，连接到: %s\n", cfg.BootstrapServers)
	fmt.Printf("Topic: %s, Group: %s\n", cfg.Topic, cfg.GroupID)
	fmt.Printf("AutoOffsetReset: %s\n", cfg.AutoOffsetReset)

	// 订阅 Topic
	topics := []string{cfg.Topic}
	err = c.SubscribeTopics(topics, rebalanceCallback)
	if err != nil {
		fmt.Printf("订阅 Topic 失败: %v\n", err)
		os.Exit(1)
	}

	// 创建上下文，用于优雅退出
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 处理信号
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 使用 WaitGroup 等待所有处理完成
	var wg sync.WaitGroup

	// 消费消息
	run := true
	for run {
		select {
		case <-sigchan:
			fmt.Println("\n收到退出信号，正在停止...")
			run = false
			cancel()

		case <-ctx.Done():
			run = false

		default:
			// 读取消息，超时 100ms
			msg, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// 忽略超时错误
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}
				log.Printf("读取消息错误: %v", err)
				continue
			}

			// 异步处理消息
			wg.Add(1)
			go func(m *kafka.Message) {
				defer wg.Done()
				processMessage(ctx, m)
			}(msg)
		}
	}

	// 等待所有消息处理完成
	wg.Wait()
	fmt.Println("消费者已关闭")
}

// rebalanceCallback 处理再平衡事件
func rebalanceCallback(c *kafka.Consumer, event kafka.Event) error {
	switch e := event.(type) {
	case kafka.AssignedPartitions:
		fmt.Printf("✓ 分配分区: %v\n", e.Partitions)
		// 可以在这里执行一些初始化操作

	case kafka.RevokedPartitions:
		fmt.Printf("✗ 撤销分区: %v\n", e.Partitions)
		// 在分区被撤销前提交位移
		c.Commit()
	}
	return nil
}

// processMessage 处理消息
func processMessage(ctx context.Context, msg *kafka.Message) {
	// 检查上下文是否已取消
	select {
	case <-ctx.Done():
		return
	default:
	}

	// 打印消息基本信息
	fmt.Printf("\n--- 收到消息 ---\n")
	fmt.Printf("Topic: %s\n", *msg.TopicPartition.Topic)
	fmt.Printf("Partition: %d\n", msg.TopicPartition.Partition)
	fmt.Printf("Offset: %d\n", msg.TopicPartition.Offset)
	fmt.Printf("Key: %s\n", string(msg.Key))

	// 打印消息内容
	fmt.Printf("Value: %s\n", string(msg.Value))

	// 打印时间戳（如果有）
	if !msg.Timestamp.IsZero() {
		fmt.Printf("Timestamp: %s\n", msg.Timestamp.UTC().Format(time.RFC3339))
	}

	// 打印 Headers（如果有）
	if len(msg.Headers) > 0 {
		fmt.Printf("Headers:\n")
		for _, h := range msg.Headers {
			fmt.Printf("  %s: %s\n", h.Key, string(h.Value))
		}
	}

	// 模拟业务处理
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("--- 处理完成 ---\n\n")
}
