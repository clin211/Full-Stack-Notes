# Kafka Go 示例

本目录包含 Kafka 的 Go 语言示例代码，使用统一的 Go Module 管理。

## 环境要求

- Go 1.23+
- Kafka 2.0+ (本地运行需要)

## 安装依赖

```bash
cd kafka/examples
go mod tidy
```

## 示例列表

### 01-simple-producer
简单同步生产者示例，演示基本的消息发送功能。

```bash
cd 01-simple-producer
go run main.go
```

### 02-simple-consumer
简单消费者示例，演示基本的消息消费功能。

```bash
cd 02-simple-consumer
go run main.go
```

### 03-advanced-producer
高级生产者示例，包含完整的生产者配置、批量发送、压缩等特性。

```bash
cd 03-advanced-producer
go run main.go
```

### 04-consumer-group
消费者组示例，演示消费者组的完整实现，包括重平衡、错误处理等。

```bash
cd 04-consumer-group
go run main.go
```

### 05-async-producer
异步生产者示例，演示异步消息发送和响应处理。

```bash
cd 05-async-producer
go run main.go
```

## 快速开始

1. 启动本地 Kafka (需要 Docker):
```bash
docker run -d \
  --name kafka \
  -p 9092:9092 \
  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
  confluentinc/cp-kafka:latest
```

2. 创建测试主题:
```bash
docker exec -it kafka kafka-topics --create --topic test-topic --bootstrap-server kafka:9092 --partitions 3 --replication-factor 1
```

3. 运行示例:
```bash
# 终端 1 - 启动消费者
cd 02-simple-consumer
go run main.go

# 终端 2 - 启动生产者
cd 01-simple-producer
go run main.go
```

## 使用的库

- [IBM Sarama](https://github.com/IBM/sarama) - Go 语言的 Kafka 客户端库

## 目录结构

```
examples/
├── go.mod                           # 统一的 Go Module
├── 01-simple-producer/
│   └── main.go                      # 简单生产者
├── 02-simple-consumer/
│   └── main.go                      # 简单消费者
├── 03-advanced-producer/
│   └── main.go                      # 高级生产者
├── 04-consumer-group/
│   └── main.go                      # 消费者组
└── 05-async-producer/
    └── main.go                      # 异步生产者
```
