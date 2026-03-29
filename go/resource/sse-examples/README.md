# go-sse 示例代码

这是《SSE 详解：从原理到实战，彻底掌握 Server-Sent Events》文章的配套示例代码。

## 目录结构

```
sse-examples/
├── 01-hello-world/          # 快速入门：Hello World
├── 02-event-replay/         # 事件重放功能
├── 03-notification-system/  # 完整的实时通知系统
├── 04-client/               # SSE 客户端示例
└── README.md               # 本文件
```

## 环境要求

- Go 1.21 或更高版本

## 安装依赖

```bash
go get github.com/tmaxmax/go-sse@latest
```

## 运行示例

### 1. Hello World

最简单的 SSE 服务器：

```bash
cd 01-hello-world
go run main.go
```

在另一个终端测试：

```bash
curl -N http://localhost:8080
```

### 2. 事件重放

演示如何实现事件历史记录和重放：

```bash
cd 02-event-replay
go run main.go
```

测试方式：

```bash
# 终端1：第一次连接
curl -N http://localhost:8081

# 终端2：等待几秒后新连接，将收到历史事件
curl -N http://localhost:8081
```

### 3. 实时通知系统

完整的用户通知系统，支持：

- 全局广播通知
- 用户专属通知
- 事件历史重放
- 认证与授权

```bash
cd 03-notification-system
go run main.go
```

访问 Web 界面：

```
http://localhost:8082/
```

使用 curl 测试：

```bash
# 连接接收通知
curl -N 'http://localhost:8082/events?user_id=alice'

# 发送个人通知
curl -X POST http://localhost:8082/notify \
  -H 'Content-Type: application/json' \
  -d '{"user_id":"alice","title":"新消息","content":"你好，Alice！"}'
```

### 4. SSE 客户端

演示如何使用 go-sse 客户端连接和接收事件：

```bash
# 先启动通知系统
cd 03-notification-system && go run main.go

# 在另一个终端运行客户端
cd 04-client
go run main.go
```

## 相关链接

- [go-sse GitHub](https://github.com/tmaxmax/go-sse)
- [go-sse 文档](https://pkg.go.dev/github.com/tmaxmax/go-sse)
- [SSE 官方规范](https://html.spec.whatwg.org/multipage/server-sent-events.html)
- [MDN SSE 文档](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
