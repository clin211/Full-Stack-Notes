// 02-event-replay: 事件重放功能演示
// 展示如何使用 Replayer 实现事件历史记录和重放
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tmaxmax/go-sse"
)

// DefaultLogger 简单的日志实现
type DefaultLogger struct{}

func (l *DefaultLogger) Log(ctx context.Context, level sse.LogLevel, msg string, data map[string]any) {
	log.Printf("[%d] %s: %v", level, msg, data)
}

func main() {
	// 使用 ValidReplayer：重放未过期的事件
	// 参数1：事件有效期 5 分钟
	// 参数2：是否自动生成 ID（false 表示需要手动设置 ID）
	replayer, err := sse.NewValidReplayer(5*time.Minute, false)
	if err != nil {
		log.Fatal(err)
	}

	joe := &sse.Joe{
		Replayer: replayer,
	}

	server := &sse.Server{
		Provider: joe,
		Logger:   &DefaultLogger{},
	}

	// 定期发送带 ID 的事件
	go func() {
		counter := 0
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			counter++
			msg := &sse.Message{}
			msg.ID = sse.ID(fmt.Sprintf("event-%d", counter))
			msg.AppendData(fmt.Sprintf("事件 #%d - %s", counter, time.Now().Format(time.RFC3339)))

			server.Publish(msg)
			log.Printf("已发布事件 #%d", counter)
		}
	}()

	log.Println("SSE 服务器启动在 :8081（支持事件重放）")
	log.Println("新连接将自动收到最近 5 分钟内的所有历史事件")
	log.Fatal(http.ListenAndServe(":8081", server))
}
