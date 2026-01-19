# 01 Kafka 核心概念与快速入门

本目录包含《Kafka 核心概念与快速入门》文章的所有示例代码。

## 目录结构

```sh
01-kafka-core-concepts/
├── config/
│   └── config.go                 # Kafka 配置管理
├── producer/
│   ├── main.go                   # 生产者主程序
│   ├── advanced.go               # 带 Headers 的消息示例
│   └── assign_partition.go       # 指定分区发送示例
├── consumer/
│   ├── main.go                   # 消费者主程序
│   └── manual_commit.go          # 手动提交位移示例
├── docker-compose-kraft.yml      # KRaft 单节点部署
├── docker-compose-zookeeper.yml  # ZooKeeper 模式部署
└── docker-compose-kraft-cluster.yml  # KRaft 三节点集群部署
```

## 快速开始

### 1. 启动 Kafka

**开发环境（推荐）**：

```bash
docker-compose -f docker-compose-kraft.yml up -d
```

**三节点集群**：

```bash
docker-compose -f docker-compose-kraft-cluster.yml up -d
```

### 2. 创建 Topic

```bash
docker exec -it kafka-kraft /usr/bin/kafka-topics --create \
  --topic quickstart-events \
  --bootstrap-server kafka:9092 \
  --partitions 3 \
  --replication-factor 1
```

### 3. 运行示例

**启动消费者**：

```bash
cd consumer
go run main.go
```

**启动生产者**：

```bash
cd producer
go run main.go
```

## 示例说明

### 生产者示例

| 文件 | 说明 |
|------|------|
| `main.go` | 基础生产者，发送简单消息和 JSON 消息 |
| `advanced.go` | 演示如何添加消息 Headers |
| `assign_partition.go` | 演示如何指定分区发送 |

### 消费者示例

| 文件 | 说明 |
|------|------|
| `main.go` | 基础消费者，包含自动提交和再平衡处理 |
| `manual_commit.go` | 演示如何手动提交消费位移 |

## 依赖

- Go 1.25+
- Docker & Docker Compose
- Kafka 7.5.0 (Confluent Platform)

## 系统依赖

在使用 `confluent-kafka-go` 之前，需要安装 `librdkafka`：

```bash
# macOS
brew install librdkafka

# Ubuntu/Debian
sudo apt-get install -y librdkafka-dev

# CentOS/RHEL
sudo yum install -y librdkafka-devel
```

## 常用命令

```bash
# 查看 Topic 列表
docker exec -it kafka-kraft kafka-topics --list --bootstrap-server localhost:9092

# 查看 Topic 详情
docker exec -it kafka-kraft kafka-topics --describe --topic quickstart-events --bootstrap-server localhost:9092

# 查看消费者组
docker exec -it kafka-kraft kafka-consumer-groups --bootstrap-server localhost:9092 --list

# 查看消费者组详情
docker exec -it kafka-kraft kafka-consumer-groups --bootstrap-server localhost:9092 --group demo-group --describe
```
