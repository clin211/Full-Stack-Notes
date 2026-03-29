// 04-client: SSE 客户端示例
// 演示如何使用 go-sse 客户端连接和接收事件
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tmaxmax/go-sse"
)

func main() {
	// 创建带取消上下文的请求
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 示例1：连接到 Hello World 服务器
	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", http.NoBody)

	// 示例2：连接到通知系统服务器
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8082/events?user_id=alice", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	// 创建连接
	conn := sse.NewConnection(req)

	// 方式1：订阅所有消息
	conn.SubscribeMessages(func(event sse.Event) {
		fmt.Printf("[消息] %s\n", event.Data)
	})

	// 方式2：订阅特定类型的事件
	conn.SubscribeEvent("info", func(event sse.Event) {
		fmt.Printf("[信息] %s\n", event.Data)
	})

	conn.SubscribeEvent("warning", func(event sse.Event) {
		fmt.Printf("[警告] %s\n", event.Data)
	})

	conn.SubscribeEvent("success", func(event sse.Event) {
		fmt.Printf("[成功] %s\n", event.Data)
	})

	conn.SubscribeEvent("error", func(event sse.Event) {
		fmt.Printf("[错误] %s\n", event.Data)
	})

	// 方式3：订阅所有事件（包括有类型的）
	conn.SubscribeToAll(func(event sse.Event) {
		fmt.Printf("[全部] 类型=%s, ID=%s, 数据=%s\n", event.Type, event.LastEventID, event.Data)
	})

	// 处理中断信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// 启动连接
	go func() {
		if err := conn.Connect(); err != nil {
			log.Printf("连接错误: %v", err)
		}
	}()

	log.Println("客户端已启动，按 Ctrl+C 退出...")
	<-sigCh
	fmt.Println("\n客户端关闭")
}

// Example: 自定义客户端配置（带重连策略）
func exampleCustomClient() {
	// 自定义客户端配置
	client := &sse.Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Backoff: sse.Backoff{
			InitialInterval: 500 * time.Millisecond, // 初始重连延迟
			Multiplier:      1.5,                     // 延迟倍增
			Jitter:          0.5,                     // 随机抖动
			MaxInterval:     30 * time.Second,        // 最大延迟
			MaxElapsedTime:  5 * time.Minute,         // 总重连时长
			MaxRetries:      0,                       // 无限重试（0 表示无限）
		},
		OnRetry: func(err error, delay time.Duration) {
			log.Printf("连接失败，%v 后重试: %v", delay, err)
		},
	}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8082/events?user_id=alice", http.NoBody)
	conn := client.NewConnection(req)

	conn.SubscribeToAll(func(event sse.Event) {
		log.Printf("收到事件: %s", event.Data)
	})

	if err := conn.Connect(); err != nil {
		log.Printf("最终连接失败: %v", err)
	}
}
