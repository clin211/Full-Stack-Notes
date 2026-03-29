// 01-hello-world: SSE 快速入门示例
// 这是 go-sse 的最简单用法演示
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tmaxmax/go-sse"
)

func main() {
	// 创建 SSE 服务器（零值即可用）
	server := &sse.Server{}

	// 启动一个 goroutine 定期发送消息
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		msg := &sse.Message{}
		msg.AppendData("Hello, SSE!")
		msg.ID = sse.ID("msg-1")

		for range ticker.C {
			if err := server.Publish(msg); err != nil {
				log.Printf("发布失败: %v", err)
			}
		}
	}()

	// 启动 HTTP 服务器
	log.Println("SSE 服务器启动在 :8080")
	log.Println("测试命令: curl -N http://localhost:8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
