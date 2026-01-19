# Kafka 两周学习方案：从零到精通

> 本学习方案参考 Apache Kafka 官方文档和 dunwu/bigdata-tutorial 教程整理
> 目标：从小白到能够满足各个场景的使用和高级特性的平衡

## 学习目标

- 理解 Kafka 的核心概念和架构设计
- 掌握生产者、消费者的开发与调优
- 熟悉 Kafka 集群搭建与运维管理
- 掌握 Kafka 高级特性（事务、Exactly-Once、流式处理）
- 能够应对各种业务场景的需求

---

## 第一周：基础入门与核心组件

### Day 1：Kafka 核心概念与快速入门

- Kafka 是什么：分布式流处理平台
- 核心概念：Producer、Broker、Consumer、Topic、Partition、Replica
- **ZooKeeper 模式 vs KRaft 模式**
- Kafka 应用场景对比
- 环境搭建（Docker 方式）

---

#### Kafka 部署模式：ZooKeeper vs KRaft

Kafka 有两种部署模式：**ZooKeeper 模式**（传统模式）和 **KRaft 模式**（Kafka Raft 模式，新增模式）。

##### ZooKeeper 模式（传统模式）

ZooKeeper 模式是 Kafka 的传统部署方式，在 Kafka 2.8 之前的版本中是唯一的选择。

**架构特点：**

```sh
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Broker 1  │     │   Broker 2  │     │   Broker 3  │
│             │     │             │     │             │
└──────┬──────┘     └──────┬──────┘     └──────┬──────┘
       │                   │                   │
       └───────────────────┼───────────────────┘
                           │
                   ┌───────┴────────┐
                   │   ZooKeeper    │
                   │   (Controller) │
                   └────────────────┘
```

**工作原理：**

1. **ZooKeeper 的作用**：
   - **Controller 选举**：ZooKeeper 负责从集群中选举出一个 Broker 作为 Controller
   - **元数据管理**：存储 Topic、Partition、Broker 等元数据信息
   - **分区 Leader 选举**：当某个分区的 Leader 故障时，Controller 通过 ZooKeeper 协议选举新的 Leader
   - **配置管理**：管理配额配置、ACL 权限等
   - **成员管理**：监控 Broker 的加入和退出

2. **Broker 角色分离**：
   - **Controller Broker**：负责管理整个集群的状态（如分区分配、Leader 选举）
   - **普通 Broker**：负责处理客户端的读写请求

**优点：**

- 成熟稳定，经过了大量生产环境验证
- 运维工具和监控体系完善
- 与旧版本兼容性好

**缺点：**

- 架构复杂，需要额外部署和维护 ZooKeeper 集群
- ZooKeeper 成为性能瓶颈（尤其是在大规模集群中）
- 运维成本高（需要维护两套系统）
- Controller 选举和元数据操作依赖 ZooKeeper，延迟较高

##### KRaft 模式（Kafka Raft 模式）

KRaft 模式是 Kafka 2.8 版本引入的重大改进，在 Kafka 3.0 版本正式可用，并在 Kafka 3.3+ 版本达到生产可用。KRaft 模式**移除了对 ZooKeeper 的依赖**，将元数据管理直接集成到 Kafka 内部。

**架构特点：**

```sh
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Broker 1  │     │   Broker 2  │     │   Broker 3  │
│ (Controller)│◄────┤  (Follower) │◄────┤  (Follower) │
│             │     │             │     │             │
└──────┬──────┘     └──────┬──────┘     └──────┬──────┘
       │                   │                   │
       └───────────────────┼───────────────────┘
                           │
                    KRaft Quorum (Raft 协议)
                    (元数据复制与选举)
```

**工作原理：**

1. **Raft 协议**：
   - 使用 Raft 一致性算法管理元数据复制和选举
   - 不再依赖外部协调服务（ZooKeeper）
   - 元数据存储在内部 Topic `__cluster_metadata` 中

2. **Quorum Controller（仲裁控制器）**：
   - 从 Broker 中选举出活跃的 Controller（通常是奇数个：3、5、7）
   - Controller 负责元数据管理、分区分配、Leader 选举等

3. **Broker 角色统一**：
   - **Controller Broker**：参与 Quorum 投票，管理元数据
   - **Broker**：只处理客户端请求，不参与元数据管理

**优点：**

- **架构简化**：不再需要部署 ZooKeeper，降低运维复杂度
- **性能提升**：元数据操作延迟降低（内部 RPC 替代 ZooKeeper 操作）
- **扩展性更好**：支持更多分区和更快的元数据变更
- **部署更简单**：只需维护 Kafka 集群
- **启动更快**：Controller 选举从数秒降至毫秒级

**缺点：**

- 相对较新，生态工具支持可能不完善
- 早期版本（2.8-3.2）有一些已知问题，建议使用 3.3+ 版本

##### 对比总结

| 特性 | ZooKeeper 模式 | KRaft 模式 |
|------|---------------|-----------|
| **Kafka 版本** | 所有版本 | 2.8+ (生产可用: 3.3+) |
| **外部依赖** | 需要 ZooKeeper 集群 | 无需外部依赖 |
| **架构复杂度** | 复杂（两套系统） | 简单（单一系统） |
| **部署难度** | 较高 | 较低 |
| **Controller 选举** | 通过 ZooKeeper | 通过 Raft 协议 |
| **元数据存储** | ZooKeeper | 内部 Topic |
| **元数据操作延迟** | 较高（ZooKeeper 网络往返） | 较低（内部 RPC） |
| **支持的分区数** | 受 ZooKeeper 性能限制 | 支持更多分区 |
| **工具成熟度** | 非常成熟 | 快速完善中 |

##### 什么时候使用哪种模式？

**使用 ZooKeeper 模式的场景：**

1. **现有生产环境**：已经在使用 ZooKeeper 模式的稳定系统
2. **工具依赖**：依赖需要 ZooKeeper 的运维工具或监控系统
3. **保守策略**：对新架构持观望态度，希望 KRaft 模式更加成熟
4. **兼容性需求**：需要与旧版本 Kafka 互操作

**使用 KRaft 模式的场景：**

1. **新部署的项目**：从零开始搭建的新 Kafka 集群
2. **简化运维**：希望减少维护的组件数量
3. **大规模部署**：需要支持更多的分区和更快的元数据操作
4. **云原生环境**：在容器化、Kubernetes 环境中部署
5. **未来规划**：KRaft 是 Kafka 未来的发展方向，官方推荐新项目使用

**迁移建议：**

- **新项目**：直接使用 KRaft 模式（推荐 Kafka 3.3+ 版本）
- **现有项目**：可以继续使用 ZooKeeper 模式，等待合适的时机迁移

---

#### 1.1 Docker 快速启动 Kafka

##### ZooKeeper 模式（传统模式）

```yaml
# docker-compose-zookeeper.yml
version: '3'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.5.0
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181

  kafka:
    image: confluentinc/cp-kafka:7.5.0
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
```

启动命令：

```bash
docker-compose -f docker-compose-zookeeper.yml up -d
```

##### KRaft 模式（推荐，Kafka 3.3+）

```yaml
# docker-compose-kraft.yml
version: '3'
services:
  kafka:
    image: confluentinc/cp-kafka:7.5.0
    hostname: kafka
    ports:
      - "9092:9092"
      - "9093:9093"
    environment:
      # KRaft 配置
      KAFKA_NODE_ID: 1
      KAFKA_PROCESS_ROLES: 'broker,controller'
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1
      KAFKA_CONTROLLER_LISTENER_NAMES: 'CONTROLLER'
      KAFKA_LISTENERS: 'PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:9093'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:9093'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT'
      KAFKA_INTER_BROKER_LISTENER_NAME: 'PLAINTEXT'
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
```

启动命令：

```bash
docker-compose -f docker-compose-kraft.yml up -d
```

##### KRaft 三节点集群（生产推荐）

```yaml
# docker-compose-kraft-cluster.yml
version: '3'
services:
  kafka1:
    image: confluentinc/cp-kafka:7.5.0
    hostname: kafka1
    ports:
      - "19092:9092"
    environment:
      KAFKA_NODE_ID: 1
      KAFKA_PROCESS_ROLES: 'broker,controller'
      KAFKA_CONTROLLER_QUORUM_VOTERS: 3
      KAFKA_CONTROLLER_LISTENER_NAMES: 'CONTROLLER'
      KAFKA_LISTENERS: 'PLAINTEXT://kafka1:9092,PLAINTEXT_HOST://localhost:19092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://kafka1:9092,PLAINTEXT_HOST://localhost:19092'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT'
      KAFKA_INTER_BROKER_LISTENER_NAME: 'PLAINTEXT'
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1

  kafka2:
    image: confluentinc/cp-kafka:7.5.0
    hostname: kafka2
    ports:
      - "19093:9092"
    environment:
      KAFKA_NODE_ID: 2
      KAFKA_PROCESS_ROLES: 'broker,controller'
      KAFKA_CONTROLLER_QUORUM_VOTERS: 3
      KAFKA_CONTROLLER_LISTENER_NAMES: 'CONTROLLER'
      KAFKA_LISTENERS: 'PLAINTEXT://kafka2:9092,PLAINTEXT_HOST://localhost:19093'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://kafka2:9092,PLAINTEXT_HOST://localhost:19093'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT'
      KAFKA_INTER_BROKER_LISTENER_NAME: 'PLAINTEXT'

  kafka3:
    image: confluentinc/cp-kafka:7.5.0
    hostname: kafka3
    ports:
      - "19094:9092"
    environment:
      KAFKA_NODE_ID: 3
      KAFKA_PROCESS_ROLES: 'broker,controller'
      KAFKA_CONTROLLER_QUORUM_VOTERS: 3
      KAFKA_CONTROLLER_LISTENER_NAMES: 'CONTROLLER'
      KAFKA_LISTENERS: 'PLAINTEXT://kafka3:9092,PLAINTEXT_HOST://localhost:19094'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://kafka3:9092,PLAINTEXT_HOST://localhost:19094'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT'
      KAFKA_INTER_BROKER_LISTENER_NAME: 'PLAINTEXT'

  kafka-ui:
    image: provectuslabs/kafka-ui:latest
    depends_on:
      - kafka1
      - kafka2
      - kafka3
    ports:
      - "8080:8080"
    environment:
      KAFKA_CLUSTERS_0_NAME: local-kraft
      KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: kafka1:9092,kafka2:9092,kafka3:9092
```

启动命令：

```bash
docker-compose -f docker-compose-kraft-cluster.yml up -d
```

#### 1.2 基础命令操作

```bash
# 创建 Topic
docker exec -it kafka-kraft /usr/bin/kafka-topics \
  --create \
  --topic quickstart-events \
  --bootstrap-server kafka:9092 \
  --partitions 3 \
  --replication-factor 1

# 列出所有 Topic
docker exec -it kafka-kraft /usr/bin/kafka-topics --list --bootstrap-server kafka:9092

# 查看 Topic 详情
docker exec -it kafka-kraft /usr/bin/kafka-topics --describe \
  --topic quickstart-events \
  --bootstrap-server kafka:9092

# 生产消息
docker exec -it kafka-kraft /usr/bin/kafka-console-producer --topic quickstart-events --bootstrap-server kafka:9092

# 消费消息
docker exec -it kafka-kraft /usr/bin/kafka-console-consumer --topic quickstart-events --from-beginning --bootstrap-server kafka:9092
```

#### 1.3 Go 客户端入门

首先安装 confluent-kafka-go 库：

```bash
go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka
```

生产者示例：

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    // 创建生产者配置
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    // 异步发送消息
    topic := "quickstart-events"
    for _, word := range []string{"key1", "key2", "key3"} {
        value := "Hello Kafka from Go!"
        p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Key:           []byte(word),
            Value:         []byte(value),
        }, nil)
    }

    // 等待消息发送完成
    p.Flush(15 * 1000)
}
```

消费者示例：

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // 创建消费者配置
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "demo-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    // 订阅 Topic
    topics := []string{"quickstart-events"}
    c.SubscribeTopics(topics, nil)

    // 设置信号处理，优雅退出
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

    // 消费消息
    run := true
    for run {
        select {
        case sig := <-sigchan:
            fmt.Printf("Caught signal %v: terminating\n", sig)
            run = false
        default:
            // 超时 100ms 轮询消息
            msg, err := c.ReadMessage(100 * 1000) // 100ms
            if err != nil {
                // 忽略超时错误
                continue
            }
            fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
        }
    }
}
```

---

### Day 2：Kafka 生产者详解

- 生产者工作原理与架构
- 核心配置参数详解
- 序列化器：String、JSON、Avro、Protobuf
- 自定义分区器

#### 2.1 生产者三种发送方式

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    // 创建生产者
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "quickstart-events"

    // 方式1：发后即忘（Fire-and-Forget）
    // 异步发送，不等待确认
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key:           []byte("key1"),
        Value:         []byte("Fire and Forget"),
    }, nil)

    // 方式2：同步发送（使用 Wait 频道）
    deliveryChan := make(chan kafka.Event, 10000)
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key:           []byte("key2"),
        Value:         []byte("Synchronous"),
    }, deliveryChan)

    // 等待消息发送完成
    e := <-deliveryChan
    m := e.(*kafka.Message)
    if m.TopicPartition.Error != nil {
        fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
    } else {
        fmt.Printf("Delivered to topic %s [%d] at offset %v\n",
            *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
    }

    // 方式3：异步发送（带回调）
    // 使用 delivery channel 接收发送结果
    for _, word := range []string{"key3", "key4", "key5"} {
        p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Key:           []byte(word),
            Value:         []byte("Asynchronous: " + word),
        }, deliveryChan)
    }

    // 处理异步发送结果
    for i := 0; i < 3; i++ {
        e := <-deliveryChan
        m := e.(*kafka.Message)
        if m.TopicPartition.Error != nil {
            fmt.Printf("Failed to deliver message: %v\n", m.TopicPartition.Error)
        } else {
            fmt.Printf("Async success - Delivered to topic %s [%d] at offset %v\n",
                *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
        }
    }

    close(deliveryChan)
    p.Flush(15 * 1000)
}
```

#### 2.2 生产者性能调优配置

```go
package main

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers":  "localhost:9092",
        "acks":               "all",                // 等待所有副本确认
        "retries":            3,                   // 重试次数
        "batch.size":         16384,              // 批量发送大小(字节)
        "linger.ms":          10,                 // 等待时间(毫秒)
        "buffer.memory":      33554432,           // 缓冲区大小(32MB)
        "compression.type":   "snappy",           // 压缩算法: none, gzip, snappy, lz4, zstd
        "max.block.ms":       60000,              // 最大阻塞时间
        "request.timeout.ms": 30000,              // 请求超时时间
        "enable.idempotence": true,               // 启用幂等性
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()
}
```

#### 2.3 自定义分区器

confluent-kafka-go 使用 librdkafka 的内置分区器，可以通过配置指定：

```go
package main

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    // 可用的分区器：
    // "random"       - 随机分区
    // "consistent"   - 一致性分区（基于 key 的 hash）
    // "consistent_random" - 一致性随机分区
    // "murmur2"      - Java 客户端默认的 murmur2 算法
    // "murmur2_random" - murmur2 或随机
    // "fnv1a"        - FNV-1a 哈希算法
    // "fnv1a_random" - FNV-1a 或随机

    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "partitioner":       "murmur2", // 使用 Java 兼容的 murmur2 分区器
        // "partitioner":       "consistent", // 或使用一致性分区器
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "test-topic"
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key:           []byte("user-123"), // 相同的 key 会被分配到同一分区
        Value:         []byte("Message content"),
    }, nil)

    p.Flush(15 * 1000)
}
```

自定义分区器（需要实现自定义逻辑）：

```go
package main

import (
    "fmt"
    "hash/fnv"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// 自定义分区逻辑
func customPartitioner(key []byte, partitionCount int32) int32 {
    if len(key) == 0 {
        // 没有 key 时使用随机分区
        return 0
    }
    // 使用 FNV 哈希算法
    h := fnv.New32a()
    h.Write(key)
    return int32(h.Sum32()) % partitionCount
}

func main() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "test-topic"
    partitionCount := int32(3) // 假设 Topic 有 3 个分区

    // 手动指定分区
    partition := customPartitioner([]byte("user-123"), partitionCount)
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
        Key:           []byte("user-123"),
        Value:         []byte("Custom partitioned message"),
    }, nil)

    fmt.Printf("Sending to partition: %d\n", partition)
    p.Flush(15 * 1000)
}
```

#### 2.4 JSON 序列化示例

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// User 用户实体
type User struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"created_at"`
}

// 序列化 User 为 JSON
func (u *User) ToJSON() ([]byte, error) {
    return json.Marshal(u)
}

// 从 JSON 反序列化 User
func UserFromJSON(data []byte) (*User, error) {
    var user User
    err := json.Unmarshal(data, &user)
    return &user, err
}

func main() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "user-topic"

    // 创建并发送用户消息
    user := &User{
        ID:        "001",
        Name:      "Alice",
        Age:       25,
        CreatedAt: time.Now(),
    }

    // 序列化为 JSON
    jsonData, err := user.ToJSON()
    if err != nil {
        panic(err)
    }

    // 发送消息
    deliveryChan := make(chan kafka.Event, 10000)
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key:           []byte(user.ID),
        Value:         jsonData,
    }, deliveryChan)

    // 等待发送结果
    e := <-deliveryChan
    m := e.(*kafka.Message)
    if m.TopicPartition.Error != nil {
        fmt.Printf("Failed to deliver message: %v\n", m.TopicPartition.Error)
    } else {
        fmt.Printf("User message delivered: %s\n", string(jsonData))
    }

    close(deliveryChan)
    p.Flush(15 * 1000)
}
```

---

### Day 3：Kafka 消费者详解

- 消费者组与再平衡机制
- 消息位移管理：自动提交 vs 手动提交
- 消费方式：poll、seek、pause/resume
- 多线程消费模型

#### 3.1 消费者组基础

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // 创建消费者配置
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "order-consumer-group", // 消费者组 ID
        "auto.offset.reset": "earliest",             // earliest/latest/none
        "enable.auto.commit": "true",                // 启用自动提交
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    // 订阅 Topic
    topics := []string{"orders-topic"}
    c.SubscribeTopics(topics, nil)

    // 设置信号处理
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

    // 消费消息
    run := true
    for run {
        select {
        case sig := <-sigchan:
            fmt.Printf("Caught signal %v: terminating\n", sig)
            run = false
        default:
            // 读取消息，超时 100ms
            msg, err := c.ReadMessage(100 * 1000)
            if err != nil {
                // 忽略超时错误
                continue
            }

            // 处理订单
            processOrder(string(msg.Value))
        }
    }
}

func processOrder(orderJSON string) {
    fmt.Printf("Processing order: %s\n", orderJSON)
}
```

#### 3.2 手动提交位移

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "log"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // 禁用自动提交
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "order-consumer-group",
        "auto.offset.reset": "earliest",
        "enable.auto.commit": "false", // 禁用自动提交
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    topics := []string{"orders-topic"}
    c.SubscribeTopics(topics, nil)

    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

    run := true
    for run {
        select {
        case sig := <-sigchan:
            fmt.Printf("Caught signal %v: terminating\n", sig)
            run = false
        default:
            msg, err := c.ReadMessage(100 * 1000)
            if err != nil {
                continue
            }

            // 处理订单
            err = processOrder(string(msg.Value))
            if err != nil {
                log.Printf("Failed to process order: %v", err)
                // 处理失败，不提交，下次重新消费
                continue
            }

            // 手动提交位移（同步提交）
            // 提交当前消息的位移
            partitions := []kafka.TopicPartition{
                {Topic: msg.TopicPartition.Topic, Partition: msg.TopicPartition.Partition},
            }
            offsets := []kafka.Offset{
                msg.TopicPartition.Offset + 1, // 提交下一条要消费的位移
            }

            // 同步提交
            err = c.CommitOffsets(partitions, offsets)
            if err != nil {
                log.Printf("Failed to commit offsets: %v", err)
            }
        }
    }
}

func processOrder(orderJSON string) error {
    fmt.Printf("Processing order: %s\n", orderJSON)
    // 处理订单逻辑...
    return nil
}
```

异步提交位移：

```go
// 异步提交（使用 CommitAsync）
c.CommitAsync(msg)
```

#### 3.3 再平衡监听器

confluent-kafka-go 使用回调函数处理再平衡事件：

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "log"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "demo-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    // 订阅 Topic 时设置 rebalance 回调
    topics := []string{"orders-topic"}

    // 使用 AssignPartitions 回调处理分区分配和撤销
    c.SubscribeTopics(topics, func(c *kafka.Consumer, event kafka.Event) error {
        switch e := event.(type) {
        case kafka.AssignedPartitions:
            // 分区分配事件
            fmt.Printf("Assigned partitions: %v\n", e.Partitions)
            // 可以在这里执行一些初始化操作

        case kafka.RevokedPartitions:
            // 分区撤销事件
            fmt.Printf("Revoked partitions: %v\n", e.Partitions)
            // 在分区被撤销前提交位移
            c.Commit()
        }
        return nil
    })

    // 消费消息
    for {
        msg, err := c.ReadMessage(100 * 1000)
        if err != nil {
            continue
        }
        fmt.Printf("Message: %s\n", string(msg.Value))
    }
}
```

#### 3.4 指定位移消费

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "demo-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    // 手动分配分区（不使用消费者组）
    topic := "orders-topic"
    partitions := []kafka.TopicPartition{
        {Topic: &topic, Partition: 0},
        {Topic: &topic, Partition: 1},
    }

    // 分配分区
    err = c.Assign(partitions)
    if err != nil {
        panic(err)
    }

    // 从指定 offset 开始消费
    offsets := []kafka.TopicPartition{
        {Topic: &topic, Partition: 0, Offset: kafka.Offset(100)},
        {Topic: &topic, Partition: 1, Offset: kafka.Offset(200)},
    }

    // 跳转到指定位置
    err = c.Seek(offsets[0], 0) // 从 offset 100 开始
    if err != nil {
        panic(err)
    }

    // 从最早的消息开始消费
    err = c.Seek(kafka.TopicPartition{Topic: &topic, Partition: 0}, kafka.OffsetBeginning)
    if err != nil {
        panic(err)
    }

    // 从最新的消息开始消费
    err = c.Seek(kafka.TopicPartition{Topic: &topic, Partition: 0}, kafka.OffsetEnd)
    if err != nil {
        panic(err)
    }

    // 消费消息
    for {
        msg, err := c.ReadMessage(100 * 1000)
        if err != nil {
            continue
        }
        fmt.Printf("Message: %s\n", string(msg.Value))
    }
}
```

#### 3.5 暂停和恢复分区

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "demo-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    topics := []string{"orders-topic"}
    c.SubscribeTopics(topics, nil)

    pausedPartitions := make(map[kafka.TopicPartition]bool)

    for {
        msg, err := c.ReadMessage(100 * 1000)
        if err != nil {
            continue
        }

        // 需要暂停时
        if needToPause() {
            // 获取当前分配的分区
            assigned, err := c.Assignment()
            if err != nil {
                panic(err)
            }

            // 暂停分区
            err = c.Pause(assigned)
            if err != nil {
                panic(err)
            }

            // 记录暂停的分区
            for _, p := range assigned {
                pausedPartitions[p] = true
            }
        }

        // 处理未暂停的消息
        fmt.Printf("Message: %s\n", string(msg.Value))

        // 需要恢复时
        if shouldResume() && len(pausedPartitions) > 0 {
            var partitions []kafka.TopicPartition
            for p := range pausedPartitions {
                partitions = append(partitions, p)
            }

            // 恢复分区
            err = c.Resume(partitions)
            if err != nil {
                panic(err)
            }

            pausedPartitions = make(map[kafka.TopicPartition]bool)
        }
    }
}

func needToPause() bool {
    // 自定义逻辑判断是否需要暂停
    return false
}

func shouldResume() bool {
    // 自定义逻辑判断是否需要恢复
    return false
}
```

---

### Day 4：Kafka 消费者高级特性

**学习内容**

- 消费者组协调器
- 静态成员
- 分区分配策略
- 消费者多线程模型

**可执行示例**

#### 4.1 分区分配策略

```go
package main

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    // 分区分配策略配置
    // 可用策略: "range", "roundrobin", "sticky", "cooperative-sticky"

    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers":        "localhost:9092",
        "group.id":                 "demo-group",
        "partition.assignment.strategy": "range",   // Range 分配策略（默认）
        // "partition.assignment.strategy": "roundrobin",  // RoundRobin 分配策略
        // "partition.assignment.strategy": "sticky",      // Sticky 分配策略（减少再平衡时的分区移动）
        // "partition.assignment.strategy": "cooperative-sticky", // Cooperative Sticky（渐进式再平衡）
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()
}
```

#### 4.2 静态成员配置

```go
package main

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    // 配置静态成员
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers":  "localhost:9092",
        "group.id":           "demo-group",
        "group.instance.id":  "consumer-instance-1", // 静态成员 ID
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    // 静态成员的优势：
    // - 再平衡时优先分配回之前的分区
    // - 减少不必要的分区重新分配
    // - 降低再平衡开销
}
```

#### 4.3 单线程单消费者模型

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "single-thread-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    topics := []string{"orders-topic"}
    c.SubscribeTopics(topics, nil)

    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

    run := true
    for run {
        select {
        case <-sigchan:
            run = false
        default:
            msg, err := c.ReadMessage(100 * 1000)
            if err != nil {
                continue
            }
            processRecord(msg)
        }
    }
}

func processRecord(msg *kafka.Message) {
    fmt.Printf("Processing: %s\n", string(msg.Value))
}
```

#### 4.4 多线程消费者模型（每个线程独立消费者）

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "sync"
)

func main() {
    threadCount := 3
    var wg sync.WaitGroup

    for i := 0; i < threadCount; i++ {
        wg.Add(1)
        go func(threadID int) {
            defer wg.Done()

            c, err := kafka.NewConsumer(&kafka.ConfigMap{
                "bootstrap.servers": "localhost:9092",
                "group.id":          "multi-thread-group",
                "auto.offset.reset": "earliest",
            })
            if err != nil {
                fmt.Printf("Consumer %d error: %v\n", threadID, err)
                return
            }
            defer c.Close()

            topics := []string{"orders-topic"}
            c.SubscribeTopics(topics, nil)

            for {
                msg, err := c.ReadMessage(100 * 1000)
                if err != nil {
                    continue
                }
                fmt.Printf("Thread %d processing: %s\n", threadID, string(msg.Value))
            }
        }(i)
    }

    wg.Wait()
}
```

#### 4.5 单消费者多线程处理模型

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "sync"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "single-consumer-multi-worker-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    topics := []string{"orders-topic"}
    c.SubscribeTopics(topics, nil)

    // 创建工作线程池
    workerCount := 10
    workerChan := chan *kafka.Message{}
    var wg sync.WaitGroup

    // 启动工作线程
    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for msg := range workerChan {
                processRecordAsync(msg, workerID)
            }
        }(i)
    }

    // 主消费循环
    for {
        msg, err := c.ReadMessage(100 * 1000)
        if err != nil {
            continue
        }
        // 将消息发送到工作线程处理
        workerChan <- msg
    }
}

func processRecordAsync(msg *kafka.Message, workerID int) {
    fmt.Printf("Worker %d processing: %s\n", workerID, string(msg.Value))
}
```

---

### Day 5：Kafka 高级特性 - 幂等性与事务

- 幂等生产者原理
- 事务生产者配置
- Exactly-Once 语义
- 事务消费者配置

#### 5.1 幂等生产者

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers":  "localhost:9092",
        "enable.idempotence": true, // 启用幂等性
        // 以下配置由幂等性自动设置
        // "acks": "all",
        // "max.in.flight": 5,
        // "retries": 2147483647,
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "idempotent-topic"
    deliveryChan := make(chan kafka.Event, 10000)

    // 幂等性保证：即使重试，也只会写入一条消息
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("key-%d", i)
        value := fmt.Sprintf("value-%d", i)

        p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Key:           []byte(key),
            Value:         []byte(value),
        }, deliveryChan)
    }

    // 等待所有消息发送完成
    for i := 0; i < 10; i++ {
        <-deliveryChan
    }

    close(deliveryChan)
    p.Flush(15 * 1000)
}
```

#### 5.2 事务生产者

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "transactional.id":  "my-transactional-producer-1", // 必须配置 transactional.id
        "enable.idempotence": true,
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    // 初始化事务
    err = p.InitTransactions(10000) // 10秒超时
    if err != nil {
        panic(err)
    }

    // 开始事务
    err = p.BeginTransaction()
    if err != nil {
        panic(err)
    }

    topic1 := "orders-topic"
    topic2 := "inventory-topic"
    topic3 := "notifications-topic"

    // 在事务中发送多条消息（可以是多个 Topic）
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic1, Partition: kafka.PartitionAny},
        Key:           []byte("order-1"),
        Value:         []byte("Order data 1"),
    }, nil)

    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic2, Partition: kafka.PartitionAny},
        Key:           []byte("inventory-1"),
        Value:         []byte("Inventory update 1"),
    }, nil)

    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic3, Partition: kafka.PartitionAny},
        Key:           []byte("notify-1"),
        Value:         []byte("Notification 1"),
    }, nil)

    // 提交事务
    err = p.CommitTransaction(10000)
    if err != nil {
        fmt.Printf("Failed to commit transaction: %v\n", err)
        // 中止事务
        p.AbortTransaction()
        return
    }

    fmt.Println("Transaction committed successfully")

    // 等待所有消息发送
    p.Flush(15 * 1000)
}
```

#### 5.3 跨多个 Topic 的事务示例

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Order struct {
    ID     string  `json:"id"`
    UserID string  `json:"user_id"`
    Amount float64 `json:"amount"`
}

func ProcessOrderWithTransaction(order Order) error {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "transactional.id":  "order-service-" + order.ID,
        "enable.idempotence": true,
    })
    if err != nil {
        return err
    }
    defer p.Close()

    // 初始化事务
    err = p.InitTransactions(10000)
    if err != nil {
        return err
    }

    // 开始事务
    err = p.BeginTransaction()
    if err != nil {
        return err
    }

    // 发送到多个 Topic
    orderData, _ := json.Marshal(order)

    topics := []string{"orders-topic", "inventory-topic", "notifications-topic"}
    for _, topic := range topics {
        p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Key:           []byte(order.ID),
            Value:         orderData,
        }, nil)
    }

    // 提交事务（所有消息要么全部成功，要么全部失败）
    err = p.CommitTransaction(10000)
    if err != nil {
        p.AbortTransaction()
        return err
    }

    p.Flush(15 * 1000)
    return nil
}
```

#### 5.4 事务消费者

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "transactional-consumer-group",
        "auto.offset.reset": "earliest",
        // 设置隔离级别为 read_committed
        // 只能读取到已提交事务的消息
        "isolation.level": "read_committed",
        // 如果设置为 read_uncommitted（默认）
        // 可以读取到所有消息，包括未提交事务的消息
        // "isolation.level": "read_uncommitted",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    topics := []string{"orders-topic"}
    c.SubscribeTopics(topics, nil)

    for {
        msg, err := c.ReadMessage(100 * 1000)
        if err != nil {
            continue
        }
        fmt.Printf("Consumed: %s\n", string(msg.Value))
    }
}
```

#### 5.5 消费-生产-消费模式（Exactly-Once）

```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type StreamProcessor struct {
    consumer *kafka.Consumer
    producer *kafka.Producer
}

func NewStreamProcessor() (*StreamProcessor, error) {
    // 消费者配置
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "stream-processor",
        "isolation.level":   "read_committed",
        "enable.auto.commit": "false",
    })
    if err != nil {
        return nil, err
    }

    // 生产者配置
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "transactional.id":  "stream-processor-1",
    })
    if err != nil {
        c.Close()
        return nil, err
    }

    // 初始化事务
    err = p.InitTransactions(10000)
    if err != nil {
        c.Close()
        p.Close()
        return nil, err
    }

    return &StreamProcessor{consumer: c, producer: p}, nil
}

func (sp *StreamProcessor) Process() error {
    inputTopic := "input-topic"
    outputTopic := "output-topic"

    sp.consumer.SubscribeTopics([]string{inputTopic}, nil)

    for {
        msg, err := sp.consumer.ReadMessage(100 * 1000)
        if err != nil {
            continue
        }

        // 开始事务
        err = sp.producer.BeginTransaction()
        if err != nil {
            return err
        }

        // 处理消息
        processedValue := transform(string(msg.Value))

        // 发送处理后的消息
        sp.producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
            Key:           msg.Key,
            Value:         []byte(processedValue),
        }, nil)

        // 提交消费位移到事务
        offsets := []kafka.TopicPartition{
            {Topic: msg.TopicPartition.Topic, Partition: msg.TopicPartition.Partition, Offset: msg.TopicPartition.Offset + 1},
        }
        err = sp.producer.SendOffsetsToTransaction(
            offsets,
            sp.consumer.GetMetadata(),
            10000,
        )
        if err != nil {
            sp.producer.AbortTransaction()
            return err
        }

        // 提交事务
        err = sp.producer.CommitTransaction(10000)
        if err != nil {
            sp.producer.AbortTransaction()
            return err
        }
    }
}

func transform(value string) string {
    // 数据转换逻辑
    return strings.ToUpper(value)
}
```

---

### Day 6：Kafka Connect 简介

- Kafka Connect 概念与架构
- Source Connector 和 Sink Connector
- 运行模式：Standalone 和 Distributed
- 常用 Connector 实战

#### 6.1 FileStream Source Connector

`connect-file-source.properties`:

```properties
name=local-file-source
connector.class=FileStreamSourceConnector
tasks.max=1
file=/tmp/test.txt
topic=connect-test
```

#### 6.2 FileStream Sink Connector

`connect-file-sink.properties`:

```properties
name=local-file-sink
connector.class=FileStreamSinkConnector
tasks.max=1
file=/tmp/test.sink.txt
topics=connect-test
```

#### 6.3 启动 Standalone Worker

```bash
# 配置文件
connect-standalone.properties:
    bootstrap.servers=localhost:9092
    key.converter=org.apache.kafka.connect.json.JsonConverter
    value.converter=org.apache.kafka.connect.json.JsonConverter
    offset.storage.file.filename=/tmp/connect.offsets
    offset.flush.interval.ms=10000

# 启动命令
connect-standalone.sh connect-standalone.properties \
    connect-file-source.properties \
    connect-file-sink.properties
```

#### 6.4 JDBC Source Connector（从数据库读取）

```json
{
  "name": "jdbc-source-connector",
  "config": {
    "connector.class": "io.confluent.connect.jdbc.JdbcSourceConnector",
    "database.hostname": "localhost",
    "database.port": "3306",
    "database.user": "user",
    "database.password": "password",
    "database.url": "jdbc:mysql://localhost:3306/mydb",
    "table.whitelist": "users,orders",
    "mode": "timestamp+incrementing",
    "timestamp.column.name": "updated_at",
    "incrementing.column.name": "id",
    "topic.prefix": "mysql-",
    "poll.interval.ms": 1000
  }
}
```

#### 6.5 JDBC Sink Connector（写入数据库）

```json
{
  "name": "jdbc-sink-connector",
  "config": {
    "connector.class": "io.confluent.connect.jdbc.JdbcSinkConnector",
    "database.hostname": "localhost",
    "database.port": "3306",
    "database.user": "user",
    "database.password": "password",
    "database.url": "jdbc:mysql://localhost:3306/mydb",
    "topics": "orders-topic",
    "auto.create": "true",
    "auto.evolve": "true",
    "insert.mode": "upsert",
    "primary.key.mode": "record_key",
    "primary.key.fields": "order_id"
  }
}
```

#### 6.6 Elasticsearch Sink Connector

```json
{
  "name": "es-sink-connector",
  "config": {
    "connector.class": "io.confluent.connect.elasticsearch.ElasticsearchSinkConnector",
    "tasks.max": "1",
    "topics": "logs-topic",
    "connection.url": "http://localhost:9200",
    "type.name": "_doc",
    "key.ignore": "true",
    "schema.ignore": "true"
  }
}
```

---

### Day 7：Kafka Streams 入门

- Kafka Streams 概念与架构
- Stream 和 Table 的区别
- KStream 和 KTable
- WordCount 示例

#### 7.1 Kafka Streams Hello World

**重要说明**: Go 生态中没有官方的 Kafka Streams API 等价物。Java 的 Kafka Streams 提供了高级流处理 DSL（如 KStream、KTable、窗口、聚合等），而 Go 的 Kafka 客户端库主要提供低级别的消费者和生产者 API。

在 Go 中实现流处理，有以下几种方案：

**方案1: 使用 confluent-kafka-go 手动实现流处理**

```go
package main

import (
	"fmt"
	"strings"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"regexp"
	"sync"
)

type WordCountProcessor struct {
	consumer    *kafka.Consumer
	producer    *kafka.Producer
	wordCounts  map[string]int64
	mu          sync.Mutex
}

func NewWordCountProcessor() (*WordCountProcessor, error) {
	// 创建消费者
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "wordcount-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	// 创建生产者
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		c.Close()
		return nil, err
	}

	return &WordCountProcessor{
		consumer:   c,
		producer:   p,
		wordCounts: make(map[string]int64),
	}, nil
}

func (wcp *WordCountProcessor) Process() error {
	// 订阅输入 Topic
	err := wcp.consumer.SubscribeTopics([]string{"text-lines-topic"}, nil)
	if err != nil {
		return err
	}

	outputTopic := "word-count-output"

	fmt.Println("WordCount application started...")

	for {
		msg, err := wcp.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		// 处理文本行
		textLine := string(msg.Value)
		words := wcp.splitWords(textLine)

		// 更新单词计数
		wcp.mu.Lock()
		for _, word := range words {
			wcp.wordCounts[word]++
		}

		// 输出结果到 Kafka
		for word, count := range wcp.wordCounts {
			result := fmt.Sprintf("%s:%d", word, count)
			wcp.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
				Key:           []byte(word),
				Value:         []byte(result),
			}, nil)
		}
		wcp.mu.Unlock()

		wcp.producer.Flush(15 * 1000)
	}
}

func (wcp *WordCountProcessor) splitWords(text string) []string {
	// 转换为小写并拆分单词
	text = strings.ToLower(text)
	re := regexp.MustCompile(`\W+`)
	words := re.Split(text, -1)

	// 过滤空字符串
	var result []string
	for _, word := range words {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}

func (wcp *WordCountProcessor) Close() {
	wcp.consumer.Close()
	wcp.producer.Close()
}

func main() {
	processor, err := NewWordCountProcessor()
	if err != nil {
		panic(err)
	}
	defer processor.Close()

	// 处理流数据
	processor.Process()
}
```

**方案2: 使用 Go 流处理框架（推荐用于生产环境）**

Go 生态中有一些第三方流处理框架可以替代 Kafka Streams：

1. **goka** (<https://github.com/lovoo/goka>) - 一个强大的 Kafka 流处理框架
2. **kafka-go** (<https://github.com/segmentio/kafka-go>) - 更低级别的 Kafka 客户端
3. **Tempest** (<https://github.com/amirgamil/tempest>) - 类似 Kafka Streams 的 Go 实现

使用 goka 的示例：

```go
// 首先安装 goka
// go get github.com/lovoo/goka

package main

import (
	"fmt"
	"log"
	"github.com/lovoo/goka"
)

func main() {
	// 定义处理器
	g := goka.NewProcessor(nil,
		goka.DefineGroup("wordcount-processor"),
		goka.Input("text-lines-topic", new(goka.StringCodec), wordCount),
		goka.Output("word-count-output", new(goka.StringCodec)),
		goka.Persist(new(goka.StringCodec)),
	)

	// 运行处理器
	if err := g.Run(); err != nil {
		log.Fatalf("Error running processor: %v", err)
	}
}

func wordCount(ctx goka.Context, msg interface{}) {
	var wordCounts map[string]int
	if val := ctx.Value(); val != nil {
		wordCounts = val.(map[string]int)
	} else {
		wordCounts = make(map[string]int)
	}

	// 处理消息
	text := msg.(string)
	// 这里简化处理，实际需要拆分单词
	// wordCounts[word]++

	ctx.Emit("word-count-output", ctx.Key(), wordCounts)
	ctx.SetValue(wordCounts)
}
```

**总结**: 对于简单的流处理任务，可以使用 confluent-kafka-go 手动实现。对于复杂的流处理需求（窗口、聚合、连接等），建议使用 goka 或其他专门的流处理框架。

#### 7.2 流式处理拓扑

**使用 confluent-kafka-go 实现流式处理拓扑**:

```go
package main

import (
	"fmt"
	"strings"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type StreamProcessor struct {
	consumer       *kafka.Consumer
	producer       *kafka.Producer
	importantTopic string
	normalTopic    string
}

func NewStreamProcessor() (*StreamProcessor, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "stream-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		c.Close()
		return nil, err
	}

	return &StreamProcessor{
		consumer:       c,
		producer:       p,
		importantTopic: "important-topic",
		normalTopic:    "normal-topic",
	}, nil
}

func (sp *StreamProcessor) Process() error {
	err := sp.consumer.SubscribeTopics([]string{"input-topic"}, nil)
	if err != nil {
		return err
	}

	fmt.Println("Stream processing started...")

	for {
		msg, err := sp.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		// 转换操作
		transformedValue := sp.transformMessage(string(msg.Value))

		// 分支处理
		var targetTopic string
		if strings.HasPrefix(transformedValue, "IMPORTANT") {
			targetTopic = sp.importantTopic
		} else {
			targetTopic = sp.normalTopic
		}

		// 发送到对应的 Topic
		sp.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &targetTopic, Partition: kafka.PartitionAny},
			Key:           msg.Key,
			Value:         []byte(transformedValue),
		}, nil)

		sp.producer.Flush(15 * 1000)
	}
}

func (sp *StreamProcessor) transformMessage(value string) string {
	// 过滤空值
	if value == "" || len(value) == 0 {
		return ""
	}

	// 转换为大写
	return strings.ToUpper(value)
}

func (sp *StreamProcessor) Close() {
	sp.consumer.Close()
	sp.producer.Close()
}

func main() {
	processor, err := NewStreamProcessor()
	if err != nil {
		panic(err)
	}
	defer processor.Close()

	processor.Process()
}
```

**实现聚合操作**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type AggregateProcessor struct {
	consumer    *kafka.Consumer
	producer    *kafka.Producer
	aggregates  map[string]int64
	mu          sync.Mutex
}

func NewAggregateProcessor() (*AggregateProcessor, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "aggregate-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		c.Close()
		return nil, err
	}

	return &AggregateProcessor{
		consumer:   c,
		producer:   p,
		aggregates: make(map[string]int64),
	}, nil
}

func (ap *AggregateProcessor) Process() error {
	err := ap.consumer.SubscribeTopics([]string{"input-topic"}, nil)
	if err != nil {
		return err
	}

	outputTopic := "aggregated-output"

	for {
		msg, err := ap.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		// 聚合计数
		key := string(msg.Key)
		ap.mu.Lock()
		ap.aggregates[key]++
		count := ap.aggregates[key]
		ap.mu.Unlock()

		// 输出聚合结果
		result := map[string]interface{}{
			"key":   key,
			"count": count,
		}
		resultJSON, _ := json.Marshal(result)

		ap.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
			Key:           msg.Key,
			Value:         resultJSON,
		}, nil)

		ap.producer.Flush(15 * 1000)
	}
}

func (ap *AggregateProcessor) Close() {
	ap.consumer.Close()
	ap.producer.Close()
}
```

**实现 Join 操作**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type JoinProcessor struct {
	leftConsumer   *kafka.Consumer
	rightConsumer  *kafka.Consumer
	producer       *kafka.Producer
	rightTableData map[string]string
	mu             sync.RWMutex
}

func NewJoinProcessor() (*JoinProcessor, error) {
	leftC, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "join-left-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	rightC, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "join-right-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		leftC.Close()
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		leftC.Close()
		rightC.Close()
		return nil, err
	}

	return &JoinProcessor{
		leftConsumer:   leftC,
		rightConsumer:  rightC,
		producer:       p,
		rightTableData: make(map[string]string),
	}, nil
}

func (jp *JoinProcessor) Process() error {
	// 启动右流消费者（构建表）
	go func() {
		jp.rightConsumer.SubscribeTopics([]string{"right-topic"}, nil)
		for {
			msg, err := jp.rightConsumer.ReadMessage(100 * 1000)
			if err != nil {
				continue
			}

			// 更新右表数据
			key := string(msg.Key)
			value := string(msg.Value)
			jp.mu.Lock()
			jp.rightTableData[key] = value
			jp.mu.Unlock()
		}
	}()

	// 处理左流并执行 Join
	jp.leftConsumer.SubscribeTopics([]string{"left-topic"}, nil)
	outputTopic := "joined-output"

	for {
		msg, err := jp.leftConsumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		key := string(msg.Key)
		leftValue := string(msg.Value)

		// 从右表查找匹配的值
		jp.mu.RLock()
		rightValue, exists := jp.rightTableData[key]
		jp.mu.RUnlock()

		if exists {
			// 执行 Join
			joinedValue := leftValue + "-" + rightValue

			result := map[string]string{
				"key":   key,
				"left":  leftValue,
				"right": rightValue,
				"joined": joinedValue,
			}
			resultJSON, _ := json.Marshal(result)

			jp.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
				Key:           msg.Key,
				Value:         resultJSON,
			}, nil)

			jp.producer.Flush(15 * 1000)
		}
	}
}

func (jp *JoinProcessor) Close() {
	jp.leftConsumer.Close()
	jp.rightConsumer.Close()
	jp.producer.Close()
}
```

#### 7.3 窗口操作

**使用 confluent-kafka-go 实现窗口操作**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type WindowedCounter struct {
	count      int64
	windowStart time.Time
	windowEnd   time.Time
}

type WindowedProcessor struct {
	consumer       *kafka.Consumer
	producer       *kafka.Producer
	tumblingWindow map[string]*WindowedCounter
	hoppingWindow  map[string]*WindowedCounter
	mu             sync.RWMutex
	windowSize     time.Duration
	advanceBy      time.Duration
}

func NewWindowedProcessor() (*WindowedProcessor, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "windowed-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		c.Close()
		return nil, err
	}

	return &WindowedProcessor{
		consumer:       c,
		producer:       p,
		tumblingWindow: make(map[string]*WindowedCounter),
		hoppingWindow:  make(map[string]*WindowedCounter),
		windowSize:     5 * time.Minute,
		advanceBy:      1 * time.Minute,
	}, nil
}

func (wp *WindowedProcessor) Process() error {
	err := wp.consumer.SubscribeTopics([]string{"events-topic"}, nil)
	if err != nil {
		return err
	}

	outputTopic := "windowed-output"

	// 启动窗口清理和输出 goroutine
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			wp.emitWindowedResults(outputTopic)
		}
	}()

	fmt.Println("Windowed processor started...")

	for {
		msg, err := wp.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		key := string(msg.Key)
		now := time.Now()

		// 滚动窗口计数
		wp.mu.Lock()
		updateTumblingWindow(wp.tumblingWindow, key, now, wp.windowSize)
		updateHoppingWindow(wp.hoppingWindow, key, now, wp.windowSize, wp.advanceBy)
		wp.mu.Unlock()
	}
}

func updateTumblingWindow(windows map[string]*WindowedCounter, key string, now time.Time, windowSize time.Duration) {
	windowStart := now.Truncate(windowSize)
	windowEnd := windowStart.Add(windowSize)

	// 检查是否需要创建新窗口
	if wc, exists := windows[key]; !exists || now.After(wc.windowEnd) {
		windows[key] = &WindowedCounter{
			count:      1,
			windowStart: windowStart,
			windowEnd:   windowEnd,
		}
	} else {
		windows[key].count++
	}
}

func updateHoppingWindow(windows map[string]*WindowedCounter, key string, now time.Time, windowSize time.Duration, advanceBy time.Duration) {
	windowStart := now.Truncate(advanceBy)
	windowEnd := windowStart.Add(windowSize)

	windowKey := fmt.Sprintf("%s-%d", key, windowStart.Unix())

	// 检查是否需要创建新窗口
	if wc, exists := windows[windowKey]; !exists || now.After(wc.windowEnd) {
		windows[windowKey] = &WindowedCounter{
			count:      1,
			windowStart: windowStart,
			windowEnd:   windowEnd,
		}
	} else {
		windows[windowKey].count++
	}
}

func (wp *WindowedProcessor) emitWindowedResults(outputTopic string) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	now := time.Now()

	// 输出滚动窗口结果
	for key, wc := range wp.tumblingWindow {
		if now.After(wc.windowEnd) {
			result := map[string]interface{}{
				"type":         "tumbling",
				"key":          key,
				"count":        wc.count,
				"window_start": wc.windowStart.Unix(),
				"window_end":   wc.windowEnd.Unix(),
			}
			resultJSON, _ := json.Marshal(result)

			wp.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
				Key:           []byte(key),
				Value:         resultJSON,
			}, nil)

			// 清理过期窗口
			delete(wp.tumblingWindow, key)
		}
	}

	// 输出滑动窗口结果
	for key, wc := range wp.hoppingWindow {
		if now.After(wc.windowEnd) {
			result := map[string]interface{}{
				"type":         "hopping",
				"key":          key,
				"count":        wc.count,
				"window_start": wc.windowStart.Unix(),
				"window_end":   wc.windowEnd.Unix(),
			}
			resultJSON, _ := json.Marshal(result)

			wp.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
				Key:           []byte(key),
				Value:         resultJSON,
			}, nil)

			// 清理过期窗口
			delete(wp.hoppingWindow, key)
		}
	}

	wp.producer.Flush(15 * 1000)
}

func (wp *WindowedProcessor) Close() {
	wp.consumer.Close()
	wp.producer.Close()
}
```

**实现会话窗口**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type SessionWindow struct {
	firstEventTime time.Time
	lastEventTime  time.Time
	count          int64
}

type SessionWindowProcessor struct {
	consumer          *kafka.Consumer
	producer          *kafka.Producer
	sessionWindows    map[string]*SessionWindow
	mu                sync.RWMutex
	sessionTimeout    time.Duration
}

func NewSessionWindowProcessor() (*SessionWindowProcessor, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "session-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		c.Close()
		return nil, err
	}

	return &SessionWindowProcessor{
		consumer:       c,
		producer:       p,
		sessionWindows: make(map[string]*SessionWindow),
		sessionTimeout: 10 * time.Minute,
	}, nil
}

func (swp *SessionWindowProcessor) Process() error {
	err := swp.consumer.SubscribeTopics([]string{"events-topic"}, nil)
	if err != nil {
		return err
	}

	outputTopic := "session-window-output"

	// 启动会话清理和输出 goroutine
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			swp.emitExpiredSessions(outputTopic)
		}
	}()

	fmt.Println("Session window processor started...")

	for {
		msg, err := swp.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		key := string(msg.Key)
		now := time.Now()

		// 更新会话窗口
		swp.mu.Lock()
		if session, exists := swp.sessionWindows[key]; exists {
			// 检查是否在同一会话中
			if now.Sub(session.lastEventTime) <= swp.sessionTimeout {
				// 延长会话
				session.lastEventTime = now
				session.count++
			} else {
				// 旧会话过期，创建新会话
				swp.emitSession(key, session, outputTopic)
				swp.sessionWindows[key] = &SessionWindow{
					firstEventTime: now,
					lastEventTime:  now,
					count:          1,
				}
			}
		} else {
			// 创建新会话
			swp.sessionWindows[key] = &SessionWindow{
				firstEventTime: now,
				lastEventTime:  now,
				count:          1,
			}
		}
		swp.mu.Unlock()
	}
}

func (swp *SessionWindowProcessor) emitSession(key string, session *SessionWindow, outputTopic string) {
	result := map[string]interface{}{
		"type":              "session",
		"key":               key,
		"count":             session.count,
		"session_start":     session.firstEventTime.Unix(),
		"session_end":       session.lastEventTime.Unix(),
		"session_duration":  session.lastEventTime.Sub(session.firstEventTime).Seconds(),
	}
	resultJSON, _ := json.Marshal(result)

	swp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
		Key:           []byte(key),
		Value:         resultJSON,
	}, nil)

	swp.producer.Flush(15 * 1000)
}

func (swp *SessionWindowProcessor) emitExpiredSessions(outputTopic string) {
	swp.mu.Lock()
	defer swp.mu.Unlock()

	now := time.Now()

	for key, session := range swp.sessionWindows {
		if now.Sub(session.lastEventTime) > swp.sessionTimeout {
			// 会话超时，输出结果
			swp.emitSession(key, session, outputTopic)
			delete(swp.sessionWindows, key)
		}
	}
}

func (swp *SessionWindowProcessor) Close() {
	swp.consumer.Close()
	swp.producer.Close()
}
```

#### 7.4 状态存储

**使用 confluent-kafka-go 实现状态存储**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// StateStore 状态存储接口
type StateStore interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Delete(key string)
	Flush() error
	Close() error
}

// InMemoryStateStore 内存状态存储实现
type InMemoryStateStore struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewInMemoryStateStore() *InMemoryStateStore {
	return &InMemoryStateStore{
		data: make(map[string]interface{}),
	}
}

func (s *InMemoryStateStore) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.data[key]
	return val, exists
}

func (s *InMemoryStateStore) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *InMemoryStateStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func (s *InMemoryStateStore) Flush() error {
	return nil
}

func (s *InMemoryStateStore) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string]interface{})
	return nil
}

// FileStateStore 文件持久化状态存储实现
type FileStateStore struct {
	data      map[string]int64
	filePath  string
	mu        sync.RWMutex
}

func NewFileStateStore(filePath string) (*FileStateStore, error) {
	store := &FileStateStore{
		data:     make(map[string]int64),
		filePath: filePath,
	}

	// 从文件加载状态
	if err := store.load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *FileStateStore) load() error {
	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 文件不存在，首次运行
		}
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&s.data)
}

func (s *FileStateStore) save() error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(s.data)
}

func (s *FileStateStore) Increment(key string) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key]++
	count := s.data[key]

	// 持久化到文件
	if err := s.save(); err != nil {
		return 0, err
	}

	return count, nil
}

func (s *FileStateStore) Get(key string) (int64, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.data[key]
	return val, exists
}

func (s *FileStateStore) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.save()
}

func (s *FileStateStore) Close() error {
	return s.Flush()
}

// StatefulProcessor 有状态处理器
type StatefulProcessor struct {
	consumer    *kafka.Consumer
	producer    *kafka.Producer
	countStore  *FileStateStore
	httpServer  *HttpQueryServer
}

func NewStatefulProcessor() (*StatefulProcessor, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "stateful-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		c.Close()
		return nil, err
	}

	// 创建状态存储
	countStore, err := NewFileStateStore("/tmp/kafka-state-store.json")
	if err != nil {
		c.Close()
		p.Close()
		return nil, err
	}

	// 创建 HTTP 查询服务器
	httpServer := NewHttpQueryServer(countStore)

	return &StatefulProcessor{
		consumer:   c,
		producer:   p,
		countStore: countStore,
		httpServer: httpServer,
	}, nil
}

func (sp *StatefulProcessor) Process() error {
	err := sp.consumer.SubscribeTopics([]string{"input-topic"}, nil)
	if err != nil {
		return err
	}

	outputTopic := "count-output"

	// 启动 HTTP 查询服务器
	go sp.httpServer.Start(":8080")

	fmt.Println("Stateful processor started...")
	fmt.Println("HTTP query server started on :8080")

	for {
		msg, err := sp.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		key := string(msg.Key)

		// 使用状态存储进行聚合
		count, err := sp.countStore.Increment(key)
		if err != nil {
			fmt.Printf("Error incrementing count: %v\n", err)
			continue
		}

		// 输出结果
		result := map[string]interface{}{
			"key":   key,
			"count": count,
		}
		resultJSON, _ := json.Marshal(result)

		sp.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &outputTopic, Partition: kafka.PartitionAny},
			Key:           msg.Key,
			Value:         resultJSON,
		}, nil)

		sp.producer.Flush(15 * 1000)
	}
}

func (sp *StatefulProcessor) Close() {
	sp.consumer.Close()
	sp.producer.Close()
	sp.countStore.Close()
	sp.httpServer.Stop()
}

// HttpQueryServer HTTP 查询服务器（交互式查询）
type HttpQueryServer struct {
	store *FileStateStore
}

func NewHttpQueryServer(store *FileStateStore) *HttpQueryServer {
	return &HttpQueryServer{store: store}
}

func (s *HttpQueryServer) Start(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/query/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[len("/query/"):]
		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		count, exists := s.store.Get(key)
		if !exists {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		result := map[string]interface{}{
			"key":   key,
			"count": count,
		}
		json.NewEncoder(w).Encode(result)
	})

	mux.HandleFunc("/query/all", func(w http.ResponseWriter, r *http.Request) {
		// 获取所有状态（需要修改 StateStore 以支持此功能）
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Use /query/{key} to query specific keys",
		})
	})

	fmt.Printf("HTTP query server listening on %s\n", addr)
	return http.ListenAndServe(addr, mux)
}

func (s *HttpQueryServer) Stop() {
	// HTTP 服务器会自动关闭
}

func main() {
	processor, err := NewStatefulProcessor()
	if err != nil {
		panic(err)
	}
	defer processor.Close()

	// 处理消息
	processor.Process()
}
```

**使用示例**:

```bash
# 启动处理器
go run stateful_processor.go

# 查询状态
curl http://localhost:8080/query/key-1
# 返回: {"key":"key-1","count":42}
```

---

## 第二周：集群运维与企业级应用

### Day 8：Kafka 集群部署

- Kafka 集群架构
- 多节点 Broker 部署
- ZooKeeper vs KRaft 模式
- 集群配置管理

#### 8.1 三节点 Kafka 集群（ZooKeeper 模式）

```yaml
# docker-compose-cluster.yml
version: '3'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.5.0
    hostname: zookeeper
    container_name: zookeeper
    ports:
      - "2181:2181"
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  kafka-broker1:
    image: confluentinc/cp-kafka:7.5.0
    hostname: kafka-broker1
    container_name: kafka-broker1
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-broker1:9092,PLAINTEXT_HOST://localhost:9092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 3
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 3
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 2

  kafka-broker2:
    image: confluentinc/cp-kafka:7.5.0
    hostname: kafka-broker2
    container_name: kafka-broker2
    depends_on:
      - zookeeper
    ports:
      - "9093:9093"
    environment:
      KAFKA_BROKER_ID: 2
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-broker2:9093,PLAINTEXT_HOST://localhost:9093
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 3

  kafka-broker3:
    image: confluentinc/cp-kafka:7.5.0
    hostname: kafka-broker3
    container_name: kafka-broker3
    depends_on:
      - zookeeper
    ports:
      - "9094:9094"
    environment:
      KAFKA_BROKER_ID: 3
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-broker3:9094,PLAINTEXT_HOST://localhost:9094
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 3

  kafka-ui:
    image: provectuslabs/kafka-ui:latest
    container_name: kafka-ui
    depends_on:
      - kafka-broker1
      - kafka-broker2
      - kafka-broker3
    ports:
      - "8080:8080"
    environment:
      KAFKA_CLUSTERS_0_NAME: local
      KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: kafka-broker1:9092,kafka-broker2:9093,kafka-broker3:9094
```

#### 8.2 创建高可用 Topic

```bash
# 创建 3 副本、6 分区的 Topic
docker exec -it kafka-kraft /usr/bin/kafka-topics \
  --create \
  --bootstrap-server kafka:9092 \
  --topic high-availability-topic \
  --partitions 6 \
  --replication-factor 3

# 查看 Topic 详情
docker exec -it kafka-kraft /usr/bin/kafka-topics \
  --describe \
  --bootstrap-server kafka:9092 \
  --topic high-availability-topic
```

#### 8.3 常用集群管理命令

```bash
# 查看所有 Broker
kafka-broker-api-versions.sh --bootstrap-server localhost:9092

# 查看所有 Topic
kafka-topics.sh --list --bootstrap-server localhost:9092

# 修改 Topic 分区数（只能增加，不能减少）
kafka-topics.sh --bootstrap-server localhost:9092 \
  --topic high-availability-topic \
  --alter --partitions 12

# 删除 Topic
kafka-topics.sh --bootstrap-server localhost:9092 \
  --topic old-topic \
  --delete

# 查看 Topic 消费情况
kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --group my-consumer-group \
  --describe
```

---

### Day 9：Kafka 可靠性与副本机制

- 副本机制：Leader、Follower、ISR
- 不完全 Leader 选举
- 最小同步副本数
- 数据可靠性保障

#### 9.1 创建高可靠 Topic

```bash
# 创建高可靠 Topic
docker exec -it kafka-kraft /usr/bin/kafka-topics \
  --create \
  --bootstrap-server kafka:9092 \
  --topic reliable-topic \
  --partitions 6 \
  --replication-factor 3 \
  --config min.insync.replicas=2 \
  --config unclean.leader.election.enable=false \
  --config min.cleanable.dirty.ratio=0.5
```

#### 9.2 Topic 级别配置

```bash
# 设置 Topic 级别的配置
kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type topics \
  --entity-name reliable-topic \
  --alter --add-config min.insync.replicas=2

# 查看 Topic 配置
kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type topics \
  --entity-name reliable-topic \
  --describe

# 删除 Topic 配置
kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type topics \
  --entity-name reliable-topic \
  --alter --delete-config min.insync.replicas
```

#### 9.3 Broker 配置

```properties
# server.properties

# 副本配置
num.replica.fetchers=4                    # 副本拉取线程数
replica.fetch.max.bytes=1048576           # 副本拉取最大字节数
replica.fetch.wait.max.ms=500             # 副本拉取最大等待时间

# ISR 配置
min.insync.replicas=2                     # 最小同步副本数
unclean.leader.election.enable=false      # 是否允许非 ISR 副本成为 Leader

# 日志配置
log.retention.hours=168                   # 日志保留时间（7天）
log.segment.bytes=1073741824              # 日志段大小（1GB）
log.retention.check.interval.ms=300000    # 日志保留检查间隔
```

#### 9.4 检查副本状态

```bash
# 查看 Topic 副本分布
kafka-topics.sh --describe --bootstrap-server localhost:9092 --topic reliable-topic

# 输出示例：
# Topic: reliable-topic PartitionCount: 6 ReplicationFactor: 3
# Topic: reliable-topic Partition: 0 Leader: 1 Replicas: 1,2,3 Isr: 1,2,3
# Topic: reliable-topic Partition: 1 Leader: 2 Replicas: 2,3,1 Isr: 2,3,1

# 查看下线副本
kafka-topics.sh --describe --bootstrap-server localhost:9092 | grep UnderReplicated

# 查看 ISR 扩展/缩小的分区
kafka-topics.sh --describe --bootstrap-server localhost:9092 | grep -i isr
```

#### 9.5 生产者高可靠配置

```go
package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092,localhost:9093,localhost:9094",
		"acks":                      "all",           // 等待所有 ISR 副本确认
		"retries":                   2147483647,      // 无限重试
		"max.in.flight.requests.per.connection": 5,
		"enable.idempotence":        "true",         // 启用幂等性
		"delivery.timeout.ms":       120000,         // 交付超时时间
		"retry.backoff.ms":          100,            // 重试退避时间
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()
}
```

---

### Day 10：Kafka 性能调优

- JVM 调优
- 操作系统调优
- Broker 参数调优
- 生产者/消费者性能调优

#### 10.1 Broker JVM 配置

```bash
# kafka-server-start.sh
export KAFKA_HEAP_OPTS="-Xms4g -Xmx4g"
export KAFKA_JVM_PERFORMANCE_OPTS="-server -XX:+UseG1GC -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35 -XX:+ExplicitGCInvokesConcurrent -Djava.awt.headless=true"
```

#### 10.2 操作系统调优

```bash
# /etc/sysctl.conf
# 文件描述符
fs.file-max=100000

# TCP 连接队列
net.core.somaxconn=4096
net.ipv4.tcp_max_syn_backlog=8192

# TCP 缓冲区
net.core.rmem_default=262144
net.core.rmem_max=16777216
net.core.wmem_default=262144
net.core.wmem_max=16777216
net.ipv4.tcp_rmem=4096 87380 16777216
net.ipv4.tcp_wmem=4096 65536 16777216

# 虚拟内存
vm.swappiness=1
vm.dirty_ratio=10
vm.dirty_background_ratio=5

# 应用配置
ulimit -n 100000
```

#### 10.3 Broker 性能配置

```properties
# server.properties

# 网络线程配置
num.network.threads=8             # 网络处理线程数
num.io.threads=16                 # I/O 线程数

# Socket 缓冲区配置
socket.send.buffer.bytes=102400
socket.receive.buffer.bytes=102400
socket.request.max.bytes=104857600

# 日志配置
log.dirs=/data/kafka-logs-1,/data/kafka-logs-2
log.flush.interval.messages=10000
log.flush.interval.ms=1000
log.roll.ms=604800000             # 7天滚动一次

# 副本拉取配置
replica.fetch.max.bytes=1048576
replica.fetch.wait.max.ms=500
num.replica.fetchers=4

# 配置后台线程
background.threads=10

# 配置监控
metric.reporters=com.example.KafkaMetricsReporter
```

#### 10.4 生产者性能调优

```go
package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":               "1",

		// 批量配置
		"batch.size":    32768,      // 批量大小（32KB）
		"linger.ms":     20,         // 等待时间（20ms）

		// 缓冲区配置
		"buffer.memory":  67108864,   // 缓冲区大小（64MB）

		// 压缩配置
		"compression.type": "lz4",   // 使用 lz4 压缩

		// 超时配置
		"request.timeout.ms": 30000,
		"delivery.timeout.ms": 120000,
		"max.block.ms":       60000,
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()
}
```

#### 10.5 消费者性能调优

```go
package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "performance-group",

		// 拉取配置
		"fetch.min.bytes":          1024,      // 最小拉取字节数
		"fetch.max.wait.ms":        500,       // 最大等待时间
		"fetch.max.bytes":          52428800,  // 单次拉取最大字节数（50MB）
		"max.partition.fetch.bytes": 1048576,   // 单分区最大拉取字节数（1MB）

		// 会话配置
		"session.timeout.ms":  30000,   // 会话超时
		"heartbeat.interval.ms": 10000, // 心跳间隔
		"max.poll.interval.ms": 300000, // 最大 poll 间隔

		// Poll 配置
		"max.poll.records": 1000, // 单次 poll 最大记录数
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()
}
```

#### 10.6 性能测试工具

```bash
# 生产者性能测试
kafka-producer-perf-test.sh \
  --topic performance-topic \
  --num-records 1000000 \
  --record-size 1024 \
  --throughput 10000 \
  --producer-props bootstrap.servers=localhost:9092

# 消费者性能测试
kafka-consumer-perf-test.sh \
  --topic performance-topic \
  --messages 1000000 \
  --threads 1 \
  --bootstrap-server localhost:9092

# 输出分析结果
# 1000000 records sent, 99989.2 records/sec (97.65 MB/sec)
```

---

### Day 11：Kafka 安全机制

- SSL/TLS 加密通信
- SASL 认证：PLAIN、SCRAM
- ACL 授权管理
- 生产环境安全最佳实践

#### 11.1 生成 SSL 证书

```bash
#!/bin/bash
# generate-ssl-certs.sh

# 1. 生成 CA 密钥和证书
openssl req -new -x509 -keyout ca-key -out ca-cert -days 365 \
  -subj "/CN=Kafka-CA"

# 2. 为每个 Broker 生成密钥库
for i in 1 2 3; do
  keytool -keystore kafka-server-$i.keystore.jks \
    -alias localhost \
    -validity 365 \
    -genkey \
    -keyalg RSA \
    -dname "CN=kafka-broker$i" \
    -storepass kafka123 \
    -keypass kafka123

  # 导出证书
  keytool -keystore kafka-server-$i.keystore.jks \
    -alias localhost \
    -certreq -file cert-file-$i \
    -storepass kafka123

  # 使用 CA 签名
  openssl x509 -req -CA ca-cert -CAkey ca-key \
    -in cert-file-$i -out cert-signed-$i \
    -days 365 -CAcreateserial \
    -passin pass:kafka123

  # 导入 CA 证书
  keytool -keystore kafka-server-$i.keystore.jks \
    -alias CARoot \
    -import -file ca-cert \
    -storepass kafka123 \
    -noprompt

  # 导入签名证书
  keytool -keystore kafka-server-$i.keystore.jks \
    -alias localhost \
    -import -file cert-signed-$i \
    -storepass kafka123 \
    -noprompt

  # 创建信任库
  keytool -keystore kafka-server-$i.truststore.jks \
    -alias CARoot \
    -import -file ca-cert \
    -storepass kafka123 \
    -noprompt
done

# 为客户端生成证书
keytool -keystore kafka-client.keystore.jks \
  -alias localhost \
  -validity 365 \
  -genkey \
  -keyalg RSA \
  -dname "CN=kafka-client" \
  -storepass kafka123 \
  -keypass kafka123

# 导出并签名客户端证书
keytool -keystore kafka-client.keystore.jks \
  -alias localhost \
  -certreq -file client-cert-file \
  -storepass kafka123

openssl x509 -req -CA ca-cert -CAkey ca-key \
  -in client-cert-file -out client-cert-signed \
  -days 365 -CAcreateserial \
  -passin pass:kafka123

keytool -keystore kafka-client.keystore.jks \
  -alias CARoot \
  -import -file ca-cert \
  -storepass kafka123 \
  -noprompt

keytool -keystore kafka-client.keystore.jks \
  -alias localhost \
  -import -file client-cert-signed \
  -storepass kafka123 \
  -noprompt

keytool -keystore kafka-client.truststore.jks \
  -alias CARoot \
  -import -file ca-cert \
  -storepass kafka123 \
  -noprompt
```

#### 11.2 Broker SSL 配置

```properties
# server.properties

# SSL 监听器配置
listeners=PLAINTEXT://:9092,SSL://:9093
advertised.listeners=PLAINTEXT://localhost:9092,SSL://localhost:9093
inter.broker.listener.name=PLAINTEXT
listener.security.protocol.map=PLAINTEXT:PLAINTEXT,SSL:SSL

# SSL 密钥库配置
ssl.keystore.location=/path/to/kafka-server-1.keystore.jks
ssl.keystore.password=kafka123
ssl.key.password=kafka123

# SSL 信任库配置
ssl.truststore.location=/path/to/kafka-server-1.truststore.jks
ssl.truststore.password=kafka123

# SSL 客户端认证
ssl.client.auth=required

# 启用的 SSL 密码套件
ssl.enabled.protocols=TLSv1.2,TLSv1.3
ssl.keystore.type=JKS
ssl.truststore.type=JKS
```

#### 11.3 SASL/PLAIN 认证

**JAAS 配置文件** (`kafka_server_jaas.conf`):

```conf
KafkaServer {
    org.apache.kafka.common.security.plain.PlainLoginModule required
    username="admin"
    password="admin-secret"
    user_admin="admin-secret"
    user_alice="alice-secret"
    user_bob="bob-secret";
};
```

**Broker 配置** (`server.properties`):

```properties
# 启用 SASL
listeners=SASL_PLAINTEXT://:9092
security.inter.broker.protocol=SASL_PLAINTEXT
sasl.mechanism.inter.broker.protocol=PLAIN
sasl.enabled.mechanisms=PLAIN

# JAAS 配置路径
authorizer.class.name=kafka.security.authorizer.AclAuthorizer
allow.everyone.if.no.acl.found=false

# 添加 JAAS 配置到启动参数
# KAFKA_OPTS="-Djava.security.auth.login.config=/path/to/kafka_server_jaas.conf"
```

**生产者配置**:

```go
package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanism":    "PLAIN",
		"sasl.username":     "alice",
		"sasl.password":     "alice-secret",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()
}
```

#### 11.4 SASL/SCRAM 认证

```bash
# 创建 SCRAM 用户
kafka-configs.sh --bootstrap-server localhost:9092 \
  --alter --add-config 'SCRAM-SHA-256=[password=alice-secret],SCRAM-SHA-512=[password=alice-secret]' \
  --entity-type users \
  --entity-name alice
```

**Broker 配置**:

```properties
sasl.enabled.mechanisms=SCRAM-SHA-256,SCRAM-SHA-512
sasl.mechanism.inter.broker.protocol=SCRAM-SHA-256
```

**客户端配置**:

```go
package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// 生产者
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanism":    "SCRAM-SHA-256",
		"sasl.username":     "alice",
		"sasl.password":     "alice-secret",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// 消费者
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "demo-group",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanism":    "SCRAM-SHA-256",
		"sasl.username":     "alice",
		"sasl.password":     "alice-secret",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()
}
```

#### 11.5 ACL 授权

```bash
# 添加生产者 ACL（允许 alice 向 test-topic 写入）
kafka-acls.sh --bootstrap-server localhost:9092 \
  --add --allow-principal User:alice \
  --operation Write \
  --topic test-topic

# 添加消费者 ACL（允许 bob 从 test-topic 读取）
kafka-acls.sh --bootstrap-server localhost:9092 \
  --add --allow-principal User:bob \
  --operation Read \
  --topic test-topic

# 添加消费者组 ACL
kafka-acls.sh --bootstrap-server localhost:9092 \
  --add --allow-principal User:bob \
  --operation Read \
  --group test-group

# 查看 ACL
kafka-acls.sh --bootstrap-server localhost:9092 --list

# 删除 ACL
kafka-acls.sh --bootstrap-server localhost:9092 \
  --remove --allow-principal User:alice \
  --operation Write \
  --topic test-topic
```

---

### Day 12：Kafka 监控与运维

- Kafka 核心监控指标
- JMX 监控配置
- 常用监控工具
- 日志分析与故障排查

#### 12.1 启用 JMX 监控

```bash
# 启动 Kafka 时启用 JMX
export KAFKA_JMX_OPTS="-Dcom.sun.management.jmxremote \
  -Dcom.sun.management.jmxremote.authenticate=false \
  -Dcom.sun.management.jmxremote.ssl=false \
  -Dcom.sun.management.jmxremote.port=9999"

kafka-server-start.sh config/server.properties
```

#### 12.2 使用 JConsole 监控

```bash
# 启动 JConsole
jconsole localhost:9999

# 关键 MBean 属性：
# - kafka.server:type=BrokerTopicMetrics,name=BytesInPerSec
# - kafka.server:type=BrokerTopicMetrics,name=BytesOutPerSec
# - kafka.server:type=ReplicaManager,name=UnderReplicatedPartitions
# - kafka.controller:type=KafkaController,name=ActiveControllerCount
```

#### 12.3 Kafka-UI 部署

```yaml
# docker-compose.yml 添加
services:
  kafka-ui:
    image: provectuslabs/kafka-ui:latest
    container_name: kafka-ui
    ports:
      - "8080:8080"
    environment:
      - KAFKA_CLUSTERS_0_NAME=local-cluster
      - KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka-broker1:9092,kafka-broker2:9093,kafka-broker3:9094
      - KAFKA_CLUSTERS_0_ZOOKEEPER=zookeeper:2181
```

#### 12.4 Prometheus + Grafana 监控

**Kafka JMX Exporter 配置** (`kafka-jmx.yml`):

```yaml
rules:
  - pattern: 'kafka.server<type=(.+), name=(.+)PerSec\d+><>Count'
    name: kafka_server_$1_$2_total
    type: COUNTER
  - pattern: 'kafka.server<type=(.+), name=(.+)PerSec\d+><>OneMinuteRate'
    name: kafka_server_$1_$2_rate_1m
    type: GAUGE
```

**Prometheus 配置**:

```yaml
scrape_configs:
  - job_name: 'kafka'
    static_configs:
      - targets: ['localhost:9999']
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_match
      - source_labels: [__param_match]
        target_label: __address__
        replacement: 'jmx-exporter:5556'
```

#### 12.5 关键监控指标

```go
package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// 创建 AdminClient
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}
	defer adminClient.Close()

	// 获取所有 Topic
	topics, err := adminClient.GetMetadata(nil, true, 5000)
	if err != nil {
		fmt.Printf("Failed to get metadata: %v\n", err)
		return
	}

	fmt.Println("Topics:")
	for _, topic := range topics.Topics {
		fmt.Printf("  - %s (partitions: %d)\n", topic.Topic, len(topic.Partitions))
	}

	// 获取 Consumer Group 列表
	groups, err := adminClient.ListConsumerGroups(nil)
	if err != nil {
		fmt.Printf("Failed to list consumer groups: %v\n", err)
		return
	}

	fmt.Println("\nConsumer Groups:")
	for _, group := range groups.Valid {
		fmt.Printf("  - %s\n", group)
	}
}
```

#### 12.6 故障排查脚本

```bash
#!/bin/bash
# kafka-health-check.sh

BOOTSTRAP_SERVER="localhost:9092"

echo "=== Kafka Health Check ==="

# 1. 检查 Broker 连接
echo "1. Checking Broker connection..."
kafka-broker-api-versions.sh --bootstrap-server $BOOTSTRAP_SERVER > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "   ✓ Broker is accessible"
else
    echo "   ✗ Broker is not accessible"
    exit 1
fi

# 2. 检查 Controller
echo "2. Checking Controller..."
kafka-metadata-shell.sh --snapshot /tmp/kafka-logs/__cluster_metadata-0/00000000000000000000.log --quiet --batch controller > /dev/null 2>&1

# 3. 检查 Under Replicated 分区
echo "3. Checking Under Replicated Partitions..."
UNDER_REPLICATED=$(kafka-topics.sh --describe --bootstrap-server $BOOTSTRAP_SERVER | grep "UnderReplicated" | wc -l)
if [ "$UNDER_REPLICATED" -eq 0 ]; then
    echo "   ✓ No under replicated partitions"
else
    echo "   ✗ Found $UNDER_REPLICATED under replicated partitions"
fi

# 4. 检查消费组 Lag
echo "4. Checking Consumer Group Lag..."
kafka-consumer-groups.sh --bootstrap-server $BOOTSTRAP_SERVER --describe --all-groups | \
    awk -F',' 'NR>1 {lag+=$5} END {print "   Total Lag: " lag}'

# 5. 检查磁盘使用
echo "5. Checking Disk Usage..."
df -h /tmp/kafka-logs | tail -1 | awk '{print "   Disk Usage: " $5 " (Used: " $3 ", Available: " $4 ")"}'

echo "=== Health Check Complete ==="
```

---

### Day 13：Kafka 企业级实战场景

- 日志收集系统（ELK + Kafka）
- 事件驱动架构
- 流式数据处理管道
- 微服务异步通信

#### 13.1 日志收集系统

**日志生产者（应用端）**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// LogEntry 日志条目
type LogEntry struct {
	Level       string    `json:"level"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
	Hostname    string    `json:"hostname"`
	Application string    `json:"application"`
}

func (le *LogEntry) ToJSON() ([]byte, error) {
	return json.Marshal(le)
}

// LogProducer 日志生产者
type LogProducer struct {
	producer    *kafka.Producer
	topic       string
	hostname    string
	application string
}

func NewLogProducer(brokers string, topic string) (*LogProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		return nil, err
	}

	hostname, _ := os.Hostname()
	application := os.Getenv("APP_NAME")
	if application == "" {
		application = "my-app"
	}

	return &LogProducer{
		producer:    p,
		topic:       topic,
		hostname:    hostname,
		application: application,
	}, nil
}

func (lp *LogProducer) SendLog(level string, message string) error {
	logEntry := LogEntry{
		Level:       level,
		Message:     message,
		Timestamp:   time.Now(),
		Hostname:    lp.hostname,
		Application: lp.application,
	}

	logJSON, err := logEntry.ToJSON()
	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event, 10000)
	err = lp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &lp.topic, Partition: kafka.PartitionAny},
		Key:           []byte(level),
		Value:         logJSON,
	}, deliveryChan)

	if err != nil {
		return err
	}

	// 等待发送结果
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	close(deliveryChan)
	return nil
}

func (lp *LogProducer) Close() {
	lp.producer.Flush(15 * 1000)
	lp.producer.Close()
}

// 使用示例
func businessMethod(lp *LogProducer) {
	// 业务逻辑
	err := lp.SendLog("INFO", "Operation completed successfully")
	if err != nil {
		fmt.Printf("Failed to send log: %v\n", err)
	}

	// 模拟错误
	err = someOperation()
	if err != nil {
		lp.SendLog("ERROR", "Operation failed: "+err.Error())
	}
}

func someOperation() error {
	// 模拟业务操作
	return nil
}

func main() {
	producer, err := NewLogProducer("localhost:9092", "application-logs")
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// 发送日志
	err = producer.SendLog("INFO", "Application started")
	if err != nil {
		fmt.Printf("Failed to send log: %v\n", err)
	}

	// 执行业务方法
	businessMethod(producer)
}
```

**Logstash 配置** (`logstash.conf`):

```conf
input {
  kafka {
    bootstrap_servers => "localhost:9092"
    topics => ["application-logs"]
    group_id => "logstash-consumer"
    codec => json
  }
}

filter {
  # 解析 JSON
  json {
    source => "message"
  }

  # 添加时间戳
  date {
    match => ["timestamp", "ISO8601"]
    target => "@timestamp"
  }

  # 根据日志级别设置标签
  if [level] == "ERROR" {
    mutate { add_tag => ["error"] }
  }
}

output {
  elasticsearch {
    hosts => ["localhost:9200"]
    index => "app-logs-%{+YYYY.MM.dd}"
  }

  # 同时输出到标准输出用于调试
  stdout { codec => rubydebug }
}
```

#### 13.2 事件驱动架构

**订单事件定义**:

```go
package main

import (
	"encoding/json"
	"time"
)

// OrderEvent 订单事件
type OrderEvent struct {
	EventType string                 `json:"event_type"`
	OrderID   string                 `json:"order_id"`
	UserID    string                 `json:"user_id"`
	Amount    float64                `json:"amount"`
	Payload   map[string]interface{} `json:"payload"`
	Timestamp int64                  `json:"timestamp"`
}

// Order 订单
type Order struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Items       []OrderItem `json:"items"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     float64 `json:"price"`
}

// 工厂方法：创建订单创建事件
func OrderEventCreated(order Order) OrderEvent {
	payload := map[string]interface{}{
		"items": order.Items,
	}

	return OrderEvent{
		EventType: "ORDER_CREATED",
		OrderID:   order.ID,
		UserID:    order.UserID,
		Amount:    order.TotalAmount,
		Payload:   payload,
		Timestamp: time.Now().UnixMilli(),
	}
}

// 工厂方法：创建订单取消事件
func OrderEventCancelled(order Order) OrderEvent {
	return OrderEvent{
		EventType: "ORDER_CANCELLED",
		OrderID:   order.ID,
		UserID:    order.UserID,
		Amount:    order.TotalAmount,
		Payload:   map[string]interface{}{},
		Timestamp: time.Now().UnixMilli(),
	}
}

// 工厂方法：创建订单发货事件
func OrderEventShipped(order Order) OrderEvent {
	return OrderEvent{
		EventType: "ORDER_SHIPPED",
		OrderID:   order.ID,
		UserID:    order.UserID,
		Amount:    order.TotalAmount,
		Payload:   map[string]interface{}{},
		Timestamp: time.Now().UnixMilli(),
	}
}

func (oe OrderEvent) ToJSON() ([]byte, error) {
	return json.Marshal(oe)
}
```

**订单服务（事件发布者）**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"time"
)

type OrderService struct {
	producer       *kafka.Producer
	topic          string
	orderRepository OrderRepository
}

func NewOrderService(brokers string, topic string) (*OrderService, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"acks":              "all",
		"enable.idempotence": true,
	})
	if err != nil {
		return nil, err
	}

	return &OrderService{
		producer:       p,
		topic:          topic,
		orderRepository: NewMockOrderRepository(),
	}, nil
}

func (os *OrderService) CreateOrder(userID string, items []OrderItem) (*Order, error) {
	// 1. 创建订单
	order := &Order{
		ID:          uuid.New().String(),
		UserID:      userID,
		Items:       items,
		TotalAmount: calculateTotal(items),
		Status:      "CREATED",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 保存到数据库（模拟）
	err := os.orderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	// 2. 发布订单创建事件
	event := OrderEventCreated(*order)
	eventJSON, _ := json.Marshal(event)

	err = os.publishEvent(order.ID, eventJSON)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderService) CancelOrder(orderID string) (*Order, error) {
	// 查询订单
	order, err := os.orderRepository.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	// 更新状态
	order.Status = "CANCELLED"
	order.UpdatedAt = time.Now()

	err = os.orderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	// 发布订单取消事件
	event := OrderEventCancelled(*order)
	eventJSON, _ := json.Marshal(event)

	err = os.publishEvent(order.ID, eventJSON)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderService) ShipOrder(orderID string) (*Order, error) {
	order, err := os.orderRepository.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = "SHIPPED"
	order.UpdatedAt = time.Now()

	err = os.orderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	// 发布订单发货事件
	event := OrderEventShipped(*order)
	eventJSON, _ := json.Marshal(event)

	err = os.publishEvent(order.ID, eventJSON)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderService) publishEvent(key string, value []byte) error {
	deliveryChan := make(chan kafka.Event, 10000)

	err := os.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &os.topic, Partition: kafka.PartitionAny},
		Key:           []byte(key),
		Value:         value,
	}, deliveryChan)

	if err != nil {
		return err
	}

	// 等待发送结果
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	close(deliveryChan)
	return nil
}

func (os *OrderService) Close() {
	os.producer.Flush(15 * 1000)
	os.producer.Close()
}

func calculateTotal(items []OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

// MockOrderRepository 模拟订单仓库
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id string) (*Order, error)
}

type MockOrderRepository struct {
	orders map[string]*Order
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{
		orders: make(map[string]*Order),
	}
}

func (r *MockOrderRepository) Save(order *Order) error {
	r.orders[order.ID] = order
	return nil
}

func (r *MockOrderRepository) FindByID(id string) (*Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return nil, fmt.Errorf("order not found: %s", id)
	}
	return order, nil
}
```

**库存服务（事件订阅者）**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type InventoryService struct {
	consumer           *kafka.Consumer
	inventoryRepository InventoryRepository
}

func NewInventoryService(brokers string, groupID string) (*InventoryService, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &InventoryService{
		consumer:           c,
		inventoryRepository: NewMockInventoryRepository(),
	}, nil
}

func (is *InventoryService) Start() error {
	err := is.consumer.SubscribeTopics([]string{"order-events"}, nil)
	if err != nil {
		return err
	}

	fmt.Println("Inventory service started...")

	for {
		msg, err := is.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		var event OrderEvent
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("Failed to unmarshal event: %v\n", err)
			continue
		}

		err = is.handleOrderEvent(event)
		if err != nil {
			log.Printf("Failed to handle order event: %v\n", err)
			// 在实际应用中，可以将失败的消息发送到死信队列
		}
	}
}

func (is *InventoryService) handleOrderEvent(event OrderEvent) error {
	switch event.EventType {
	case "ORDER_CREATED":
		return is.deductInventory(event)
	case "ORDER_CANCELLED":
		return is.restoreInventory(event)
	default:
		fmt.Printf("Unknown event type: %s\n", event.EventType)
		return nil
	}
}

func (is *InventoryService) deductInventory(event OrderEvent) error {
	fmt.Printf("Deducting inventory for order: %s\n", event.OrderID)

	items, ok := event.Payload["items"].([]interface{})
	if !ok {
		return fmt.Errorf("invalid items payload")
	}

	for _, itemInterface := range items {
		itemMap, ok := itemInterface.(map[string]interface{})
		if !ok {
			continue
		}

		productID, _ := itemMap["product_id"].(string)
		quantity, _ := itemMap["quantity"].(float64)

		// 查询库存
		inventory, err := is.inventoryRepository.FindByProductID(productID)
		if err != nil {
			return fmt.Errorf("inventory not found for product: %s", productID)
		}

		// 检查库存是否充足
		if inventory.Quantity < int(quantity) {
			return fmt.Errorf("insufficient inventory for product: %s", productID)
		}

		// 扣减库存
		inventory.Quantity -= int(quantity)
		err = is.inventoryRepository.Save(inventory)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Inventory deducted for order: %s\n", event.OrderID)
	return nil
}

func (is *InventoryService) restoreInventory(event OrderEvent) error {
	fmt.Printf("Restoring inventory for cancelled order: %s\n", event.OrderID)
	// 实现库存恢复逻辑
	return nil
}

func (is *InventoryService) Close() {
	is.consumer.Close()
}

// Inventory 库存
type Inventory struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

// InventoryRepository 库存仓库接口
type InventoryRepository interface {
	Save(inventory *Inventory) error
	FindByProductID(productID string) (*Inventory, error)
}

// MockInventoryRepository 模拟库存仓库
type MockInventoryRepository struct {
	inventories map[string]*Inventory
}

func NewMockInventoryRepository() *MockInventoryRepository {
	repo := &MockInventoryRepository{
		inventories: make(map[string]*Inventory),
	}

	// 初始化一些库存
	repo.inventories["product-1"] = &Inventory{ProductID: "product-1", Quantity: 100}
	repo.inventories["product-2"] = &Inventory{ProductID: "product-2", Quantity: 50}

	return repo
}

func (r *MockInventoryRepository) Save(inventory *Inventory) error {
	r.inventories[inventory.ProductID] = inventory
	return nil
}

func (r *MockInventoryRepository) FindByProductID(productID string) (*Inventory, error) {
	inventory, exists := r.inventories[productID]
	if !exists {
		return nil, fmt.Errorf("inventory not found")
	}
	return inventory, nil
}
```

**通知服务（事件订阅者）**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type NotificationService struct {
	consumer    *kafka.Consumer
	emailService EmailService
	smsService   SmsService
}

func NewNotificationService(brokers string, groupID string) (*NotificationService, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &NotificationService{
		consumer:     c,
		emailService: NewMockEmailService(),
		smsService:   NewMockSmsService(),
	}, nil
}

func (ns *NotificationService) Start() error {
	err := ns.consumer.SubscribeTopics([]string{"order-events"}, nil)
	if err != nil {
		return err
	}

	fmt.Println("Notification service started...")

	for {
		msg, err := ns.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		var event OrderEvent
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			fmt.Printf("Failed to unmarshal event: %v\n", err)
			continue
		}

		// 处理事件（不抛出异常，避免重试导致重复通知）
		ns.handleOrderEvent(event)
	}
}

func (ns *NotificationService) handleOrderEvent(event OrderEvent) {
	switch event.EventType {
	case "ORDER_CREATED":
		ns.sendOrderConfirmation(event)
	case "ORDER_SHIPPED":
		ns.sendShippingNotification(event)
	case "ORDER_COMPLETED":
		ns.sendCompletionNotification(event)
	default:
		fmt.Printf("Ignoring event type: %s\n", event.EventType)
	}
}

func (ns *NotificationService) sendOrderConfirmation(event OrderEvent) {
	fmt.Printf("Sending order confirmation for: %s\n", event.OrderID)

	// 发送邮件
	email := Email{
		To:       getUserEmail(event.UserID),
		Subject:  "Order Confirmation - " + event.OrderID,
		Template: "order-confirmation",
		Variables: map[string]interface{}{
			"orderId": event.OrderID,
			"amount":  event.Amount,
		},
	}

	err := ns.emailService.Send(email)
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
	}

	// 发送短信
	user := getUserByID(event.UserID)
	if user.SmsEnabled {
		err = ns.smsService.Send(user.PhoneNumber,
			"Your order "+event.OrderID+" has been received.")
		if err != nil {
			fmt.Printf("Failed to send SMS: %v\n", err)
		}
	}

	fmt.Printf("Order confirmation sent for: %s\n", event.OrderID)
}

func (ns *NotificationService) sendShippingNotification(event OrderEvent) {
	fmt.Printf("Sending shipping notification for: %s\n", event.OrderID)
	// 实现发货通知逻辑
}

func (ns *NotificationService) sendCompletionNotification(event OrderEvent) {
	fmt.Printf("Sending completion notification for: %s\n", event.OrderID)
	// 实现完成通知逻辑
}

func (ns *NotificationService) Close() {
	ns.consumer.Close()
}

// Email 邮件
type Email struct {
	To       string                 `json:"to"`
	Subject  string                 `json:"subject"`
	Template string                 `json:"template"`
	Variables map[string]interface{} `json:"variables"`
}

// EmailService 邮件服务接口
type EmailService interface {
	Send(email Email) error
}

// MockEmailService 模拟邮件服务
type MockEmailService struct{}

func NewMockEmailService() *MockEmailService {
	return &MockEmailService{}
}

func (s *MockEmailService) Send(email Email) error {
	fmt.Printf("Email sent to %s: %s\n", email.To, email.Subject)
	return nil
}

// SmsService 短信服务接口
type SmsService interface {
	Send(phoneNumber string, message string) error
}

// MockSmsService 模拟短信服务
type MockSmsService struct{}

func NewMockSmsService() *MockSmsService {
	return &MockSmsService{}
}

func (s *MockSmsService) Send(phoneNumber string, message string) error {
	fmt.Printf("SMS sent to %s: %s\n", phoneNumber, message)
	return nil
}

// User 用户
type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	SmsEnabled bool   `json:"sms_enabled"`
}

func getUserEmail(userID string) string {
	// 模拟查询用户邮箱
	return "user@example.com"
}

func getUserByID(userID string) User {
	// 模拟查询用户
	return User{
		ID:          userID,
		Email:       "user@example.com",
		PhoneNumber: "+1234567890",
		SmsEnabled:  true,
	}
}
```

#### 13.3 流式数据处理管道

**实时数据处理**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type StreamProcessor struct {
	consumer     *kafka.Consumer
	producer     *kafka.Producer
	userTopic    string
	systemTopic  string
	otherTopic   string
	statsTopic   string
	eventCounts  map[string]int64
	windowStart  time.Time
	windowSize   time.Duration
}

type Event struct {
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Data      string `json:"data"`
}

func NewStreamProcessor(brokers string) (*StreamProcessor, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          "data-pipeline-processor",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		c.Close()
		return nil, err
	}

	return &StreamProcessor{
		consumer:    c,
		producer:    p,
		userTopic:   "user-events",
		systemTopic: "system-events",
		otherTopic:  "other-events",
		statsTopic:  "event-statistics",
		eventCounts: make(map[string]int64),
		windowStart: time.Now(),
		windowSize:  5 * time.Minute,
	}, nil
}

func (sp *StreamProcessor) Process(inputTopic string) error {
	err := sp.consumer.SubscribeTopics([]string{inputTopic}, nil)
	if err != nil {
		return err
	}

	fmt.Println("Stream processing started...")

	ticker := time.NewTicker(sp.windowSize)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 定期输出统计结果
			sp.publishStatistics()
			// 重置计数器
			sp.eventCounts = make(map[string]int64)
			sp.windowStart = time.Now()

		default:
			msg, err := sp.consumer.ReadMessage(100 * 1000)
			if err != nil {
				continue
			}

			// 数据清洗和转换
			cleanedEvent, ok := sp.cleanAndTransform(msg.Value)
			if !ok {
				continue
			}

			// 分支处理
			var targetTopic string
			if strings.HasPrefix(cleanedEvent.Type, "user.") {
				targetTopic = sp.userTopic
			} else if strings.HasPrefix(cleanedEvent.Type, "system.") {
				targetTopic = sp.systemTopic
			} else {
				targetTopic = sp.otherTopic
			}

			// 发送到目标 Topic
			sp.produceEvent(targetTopic, msg.Key, cleanedEvent)

			// 更新统计计数
			sp.eventCounts[cleanedEvent.Type]++
		}
	}
}

func (sp *StreamProcessor) cleanAndTransform(data []byte) (*Event, bool) {
	var event Event
	err := json.Unmarshal(data, &event)
	if err != nil {
		log.Printf("Failed to parse event: %v\n", err)
		return nil, false
	}

	// 数据清洗：过滤无效数据
	if event.Timestamp == 0 || event.Type == "" {
		return nil, false
	}

	// 数据转换（可以在这里添加更多转换逻辑）
	return &event, true
}

func (sp *StreamProcessor) produceEvent(topic string, key []byte, event *Event) {
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v\n", err)
		return
	}

	err = sp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:           key,
		Value:         data,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce event: %v\n", err)
	}
}

func (sp *StreamProcessor) publishStatistics() {
	stats := map[string]interface{}{
		"window_start": sp.windowStart.Unix(),
		"window_end":   time.Now().Unix(),
		"event_counts": sp.eventCounts,
	}

	data, err := json.Marshal(stats)
	if err != nil {
		log.Printf("Failed to marshal statistics: %v\n", err)
		return
	}

	err = sp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &sp.statsTopic, Partition: kafka.PartitionAny},
		Value:         data,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce statistics: %v\n", err)
	}

	sp.producer.Flush(15 * 1000)
}

func (sp *StreamProcessor) Close() {
	sp.consumer.Close()
	sp.producer.Close()
}

func main() {
	processor, err := NewStreamProcessor("localhost:9092")
	if err != nil {
		panic(err)
	}
	defer processor.Close()

	processor.Process("raw-events")
}
```

---

### Day 14：Kafka 常见问题与最佳实践

- 消息丢失问题
- 消息重复消费问题
- 消息顺序保证
- 消息积压处理
- 生产环境最佳实践

#### 14.1 防止消息丢失

**生产端配置**:

```go
package main

import (
	"fmt"
	"log"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// 创建高可靠生产者
func CreateReliableProducer(brokers string) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        brokers,
		"acks":                     "all",                         // 等待所有 ISR 确认
		"retries":                  2147483647,                   // 最大重试次数
		"max.in.flight.requests.per.connection": 5,
		"enable.idempotence":       true,                         // 启用幂等性
		"delivery.timeout.ms":      120000,                       // 交付超时 120秒
		"request.timeout.ms":       30000,                        // 请求超时 30秒
		"retry.backoff.ms":         100,                          // 重试退避 100ms
	})

	if err != nil {
		return nil, err
	}

	return p, nil
}

// 同步发送确保可靠
func SendReliably(p *kafka.Producer, topic string, key string, value string) error {
	deliveryChan := make(chan kafka.Event, 10000)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:           []byte(key),
		Value:         []byte(value),
	}, deliveryChan)

	if err != nil {
		return err
	}

	// 等待发送结果，设置 30 秒超时
	select {
	case e := <-deliveryChan:
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %w", m.TopicPartition.Error)
		}
		log.Printf("Message sent successfully: partition=%d, offset=%d\n",
			m.TopicPartition.Partition, m.TopicPartition.Offset)
	case <-time.After(30 * time.Second):
		return fmt.Errorf("delivery timeout after 30 seconds")
	}

	close(deliveryChan)
	return nil
}
```

**消费端配置**:

```go
package main

import (
	"fmt"
	"log"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// 创建高可靠消费者
func CreateReliableConsumer(brokers string, groupID string) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
		"enable.auto.commit": "false", // 禁用自动提交
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}

// 处理消息并手动提交
func ProcessMessagesReliably(c *kafka.Consumer, topics []string) error {
	err := c.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	for {
		msg, err := c.ReadMessage(100 * 1000)
		if err != nil {
			// 忽略超时错误
			continue
		}

		// 处理消息
		err = processMessage(msg)
		if err != nil {
			log.Printf("Failed to process message: %v\n", err)
			// 处理失败，不提交，下次重新消费
			continue
		}

		// 消息处理成功后，提交位移
		// 提交当前消息的位移
		partitions := []kafka.TopicPartition{
			{Topic: msg.TopicPartition.Topic, Partition: msg.TopicPartition.Partition},
		}
		offsets := []kafka.Offset{
			msg.TopicPartition.Offset + 1, // 提交下一条要消费的位移
		}

		err = c.CommitOffsets(partitions, offsets)
		if err != nil {
			log.Printf("Failed to commit offsets: %v\n", err)
		}
	}
}

func processMessage(msg *kafka.Message) error {
	// 实现消息处理逻辑
	fmt.Printf("Processing message: %s\n", string(msg.Value))
	return nil
}
```

#### 14.2 防止消息重复消费（幂等处理）

```go
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// IdempotentMessageProcessor 幂等消息处理器
type IdempotentMessageProcessor struct {
	processedMessages map[string]time.Time
	mu               sync.RWMutex
	ttl              time.Duration // 消息 ID 保留时间
}

func NewIdempotentMessageProcessor(ttl time.Duration) *IdempotentMessageProcessor {
	imp := &IdempotentMessageProcessor{
		processedMessages: make(map[string]time.Time),
		ttl:              ttl,
	}

	// 启动清理 goroutine
	go imp.cleanupExpiredMessages()

	return imp
}

func (imp *IdempotentMessageProcessor) ProcessMessage(msg *kafka.Message) error {
	messageID := string(msg.Key)

	// 检查是否已处理
	if imp.IsProcessed(messageID) {
		log.Printf("Duplicate message detected: %s\n", messageID)
		return nil // 跳过重复消息
	}

	// 标记为已处理
	imp.MarkAsProcessed(messageID)

	// 处理业务逻辑
	err := imp.processBusinessLogic(msg)
	if err != nil {
		// 处理失败，删除标记以便重试
		imp.RemoveProcessedMark(messageID)
		return err
	}

	return nil
}

func (imp *IdempotentMessageProcessor) IsProcessed(messageID string) bool {
	imp.mu.RLock()
	defer imp.mu.RUnlock()

	_, exists := imp.processedMessages[messageID]
	return exists
}

func (imp *IdempotentMessageProcessor) MarkAsProcessed(messageID string) {
	imp.mu.Lock()
	defer imp.mu.Unlock()

	imp.processedMessages[messageID] = time.Now()
}

func (imp *IdempotentMessageProcessor) RemoveProcessedMark(messageID string) {
	imp.mu.Lock()
	defer imp.mu.Unlock()

	delete(imp.processedMessages, messageID)
}

func (imp *IdempotentMessageProcessor) processBusinessLogic(msg *kafka.Message) error {
	// 实现业务逻辑
	fmt.Printf("Processing message: %s\n", string(msg.Value))
	return nil
}

// 定期清理过期的消息 ID
func (imp *IdempotentMessageProcessor) cleanupExpiredMessages() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		imp.mu.Lock()
		now := time.Now()

		for messageID, timestamp := range imp.processedMessages {
			if now.Sub(timestamp) > imp.ttl {
				delete(imp.processedMessages, messageID)
			}
		}

		imp.mu.Unlock()
	}
}

// 使用 Redis 存储已处理消息 ID（生产环境推荐）
type RedisIdempotentProcessor struct {
	redisClient *RedisClient
	keyPrefix   string
	ttl         time.Duration
}

func (rip *RedisIdempotentProcessor) ProcessMessage(msg *kafka.Message) error {
	messageID := string(msg.Key)
	key := rip.keyPrefix + messageID

	// 使用 SETNX 原子操作检查并设置
	alreadyProcessed, err := rip.redisClient.SetNX(key, "1", rip.ttl)
	if err != nil {
		return err
	}

	if !alreadyProcessed {
		log.Printf("Duplicate message detected: %s\n", messageID)
		return nil
	}

	// 处理业务逻辑
	err = rip.processBusinessLogic(msg)
	if err != nil {
		// 处理失败，删除标记以便重试
		rip.redisClient.Del(key)
		return err
	}

	return nil
}

func (rip *RedisIdempotentProcessor) processBusinessLogic(msg *kafka.Message) error {
	fmt.Printf("Processing message: %s\n", string(msg.Value))
	return nil
}

// 数据库层面幂等（使用唯一键）
type OrderRepository interface {
	Save(order Order) error
	FindByID(id string) (*Order, error)
}

type Order struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (r *OrderRepositoryImpl) SaveOrder(order Order) (*Order, error) {
	// 尝试保存订单
	err := r.Save(order)
	if err != nil {
		// 检查是否是唯一键冲突错误
		if isDuplicateKeyError(err) {
			log.Printf("Order already exists: %s\n", order.ID)
			// 查询现有订单
			return r.FindByID(order.ID)
		}
		return nil, err
	}

	return &order, nil
}

func isDuplicateKeyError(err error) bool {
	// 检查错误类型，判断是否是唯一键冲突
	// 实现取决于使用的数据库驱动
	return false
}
```

#### 14.3 保证消息顺序

**方案1：单分区 + 单线程**:

```go
package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// 创建单分区 Topic
// kafka-topics.sh --create --topic ordered-events --partitions 1 --replication-factor 3

// 单线程消费保证顺序
func ConsumeSequentially(brokers string, groupID string) error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
		"max.poll.records":  1, // 每次只拉取一条
	})
	if err != nil {
		return err
	}
	defer c.Close()

	topics := []string{"ordered-events"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		// 处理消息
		processMessage(msg)

		// 处理完立即提交
		c.Commit()
	}
}

func processMessage(msg *kafka.Message) {
	fmt.Printf("Processing message: %s\n", string(msg.Value))
}
```

**方案2：按 Key 分区 + 顺序处理**:

```go
package main

import (
	"fmt"
	"sync"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// OrderedConsumer 按分区分组的有序消费者
type OrderedConsumer struct {
	consumer *kafka.Consumer
	workers  int
}

func NewOrderedConsumer(brokers string, groupID string, workers int) (*OrderedConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &OrderedConsumer{
		consumer: c,
		workers:  workers,
	}, nil
}

func (oc *OrderedConsumer) Start(topics []string) error {
	err := oc.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	// 创建分区处理器池
	partitionProcessors := make(map[int32]*PartitionProcessor)
	var mu sync.Mutex

	for {
		msg, err := oc.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		partition := msg.TopicPartition.Partition

		// 为每个分区创建专用的处理器
		mu.Lock()
		processor, exists := partitionProcessors[partition]
		if !exists {
			processor = NewPartitionProcessor(partition)
			partitionProcessors[partition] = processor
		}
		mu.Unlock()

		// 将消息发送到分区专用的处理器
		processor.Submit(msg)
	}
}

// PartitionProcessor 分区专用处理器
type PartitionProcessor struct {
	partition int32
	msgChan   chan *kafka.Message
	wg        sync.WaitGroup
}

func NewPartitionProcessor(partition int32) *PartitionProcessor {
	pp := &PartitionProcessor{
		partition: partition,
		msgChan:   make(chan *kafka.Message, 1000),
	}

	pp.wg.Add(1)
	go pp.process()

	return pp
}

func (pp *PartitionProcessor) Submit(msg *kafka.Message) {
	pp.msgChan <- msg
}

func (pp *PartitionProcessor) process() {
	defer pp.wg.Done()

	for msg := range pp.msgChan {
		// 同步处理保证同一分区内消息顺序
		pp.processSequentially(msg)
	}
}

func (pp *PartitionProcessor) processSequentially(msg *kafka.Message) {
	fmt.Printf("Partition %d: Processing message %s\n",
		pp.partition, string(msg.Value))
	// 处理消息逻辑
}
```

**生产者：按业务 key 分区**:

```go
package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func SendOrderedMessage(p *kafka.Producer, topic string, orderID string, orderData string) error {
	// 使用 orderID 作为 key，确保相同订单的消息进入同一分区
	deliveryChan := make(chan kafka.Event, 10000)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:           []byte(orderID), // 相同的 key 会被分配到同一分区
		Value:         []byte(orderData),
	}, deliveryChan)

	if err != nil {
		return err
	}

	// 等待发送结果
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	close(deliveryChan)
	fmt.Printf("Message sent to partition %d\n", m.TopicPartition.Partition)
	return nil
}
```

#### 14.4 处理消息积压

**方案1：增加消费者实例**:

```bash
# 扩容消费组（同一 group.id）
# 原有消费者 + 新增消费者 = 总消费者数 <= 分区数
# 例如：Topic 有 6 个分区，可以启动最多 6 个消费者实例
```

**方案2：提高消费并行度**:

```go
package main

import (
	"fmt"
	"sync"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// HighThroughputConsumer 高吞吐量消费者
type HighThroughputConsumer struct {
	consumer *kafka.Consumer
	workers  int
}

func NewHighThroughputConsumer(brokers string, groupID string, workers int) (*HighThroughputConsumer, error) {
	// 增加拉取配置
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        brokers,
		"group.id":                 groupID,
		"auto.offset.reset":        "earliest",
		"max.poll.records":         1000,                  // 增加拉取量
		"max.partition.fetch.bytes": 10485760,             // 10MB 单分区最大拉取
		"fetch.min.bytes":          1024,                  // 最小拉取字节数
		"fetch.max.wait.ms":        500,                   // 最大等待时间
	})
	if err != nil {
		return nil, err
	}

	return &HighThroughputConsumer{
		consumer: c,
		workers:  workers,
	}, nil
}

func (htc *HighThroughputConsumer) Start(topics []string) error {
	err := htc.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	// 创建工作线程池
	workerChan := make(chan *kafka.Message, htc.workers*10)
	var wg sync.WaitGroup

	// 启动工作线程
	for i := 0; i < htc.workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for msg := range workerChan {
				processMessageAsync(msg, workerID)
			}
		}(i)
	}

	// 主消费循环
	for {
		msg, err := htc.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		// 将消息发送到工作线程处理
		workerChan <- msg
	}
}

func processMessageAsync(msg *kafka.Message, workerID int) {
	fmt.Printf("Worker %d processing message: %s\n", workerID, string(msg.Value))
	// 处理消息逻辑
}
```

**方案3：临时扩容方案（积压回放）**:

```go
package main

import (
	"fmt"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// CatchupConsumer 追赶消费者，用于处理积压
type CatchupConsumer struct {
	consumer       *kafka.Consumer
	originalGroup  string
	catchupGroup   string
}

func NewCatchupConsumer(brokers string, originalGroup string, catchupGroup string) (*CatchupConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          catchupGroup, // 使用新的消费组
		"auto.offset.reset": "earliest",   // 从最早的消息开始
	})
	if err != nil {
		return nil, err
	}

	return &CatchupConsumer{
		consumer:      c,
		originalGroup: originalGroup,
		catchupGroup:  catchupGroup,
	}, nil
}

func (cc *CatchupConsumer) Start(topics []string) error {
	err := cc.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Catchup consumer started for group: %s\n", cc.catchupGroup)

	startTime := time.Now()
	messageCount := 0

	for {
		msg, err := cc.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		// 快速处理消息（简化逻辑，加快处理速度）
		processMessageFast(msg)
		messageCount++

		// 定期输出进度
		if messageCount%1000 == 0 {
			elapsed := time.Since(startTime).Seconds()
			rate := float64(messageCount) / elapsed
			fmt.Printf("Processed %d messages, rate: %.2f msg/s\n", messageCount, rate)
		}

		// 检查是否追赶完成
		if cc.isCaughtUp() {
			fmt.Println("Catchup completed!")
			break
		}
	}

	return nil
}

func (cc *CatchupConsumer) isCaughtUp() bool {
	// 检查消费位移是否接近最新位移
	// 实现取决于具体需求
	return false
}

func processMessageFast(msg *kafka.Message) {
	// 简化处理逻辑，只做必要的操作
	// 快速写入数据库或缓存
}

func (cc *CatchupConsumer) Close() {
	cc.consumer.Close()
}
```

#### 14.5 生产环境最佳实践

**Topic 设计**:

```bash
# 1. 合理设置分区数
# 分区数 = 预期最大消费者数 * (1.5 ~ 2)
# 例如：预期最多 10 个消费者，则分区数设置为 15-20

# 2. 合理设置副本数
# 生产环境：3 副本
# 测试环境：1-2 副本

# 3. Topic 命名规范
# - 使用小写字母和连字符
# - 包含环境前缀：prod-orders, dev-orders
# - 包含业务含义：user-events, order-events
```

**生产者最佳实践**:

```go
package main

import (
	"fmt"
	"log"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// 创建生产环境生产者
func CreateProductionProducer(brokers string) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        brokers,
		"acks":                     "all",                          // 等待所有副本确认
		"retries":                  3,                              // 重试次数
		"retry.backoff.ms":         100,                            // 重试退避
		"batch.size":               32768,                          // 批量大小 32KB
		"linger.ms":                10,                             // 等待时间 10ms
		"buffer.memory":            67108864,                       // 缓冲区 64MB
		"compression.type":         "snappy",                       // 压缩算法
		"max.block.ms":             60000,                          // 最大阻塞时间
		"request.timeout.ms":       30000,                          // 请求超时
		"delivery.timeout.ms":      120000,                         // 交付超时
		"enable.idempotence":       true,                           // 启用幂等性
	})

	if err != nil {
		return nil, err
	}

	return p, nil
}

// 发送消息（使用合适的 key）
func SendMessage(p *kafka.Producer, topic string, userID string, eventData []byte) error {
	// 使用 userID 作为 key，相同用户的消息会进入同一分区
	deliveryChan := make(chan kafka.Event, 10000)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:           []byte(userID), // 相同用户进入同一分区
		Value:         eventData,
	}, deliveryChan)

	if err != nil {
		return err
	}

	// 监控发送结果
	go func() {
		e := <-deliveryChan
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			log.Printf("Delivery failed for key %s: %v\n", userID, m.TopicPartition.Error)
			// 记录失败指标
			recordSendFailure(userID)
		} else {
			log.Printf("Delivered to partition %d at offset %d\n",
				m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
		close(deliveryChan)
	}()

	return nil
}

func recordSendFailure(key string) {
	// 实现失败指标记录
	// 可以发送到监控系统
}
```

**消费者最佳实践**:

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// 创建生产环境消费者
func CreateProductionConsumer(brokers string, groupID string) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    brokers,
		"group.id":             groupID,
		"auto.offset.reset":    "earliest",
		"session.timeout.ms":   30000,   // 会话超时 30秒
		"heartbeat.interval.ms": 10000,   // 心跳间隔 10秒
		"max.poll.interval.ms":  300000,  // 最大 poll 间隔 5分钟
		"enable.auto.commit":    "false",  // 手动提交
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}

// 优雅关闭消费者
func ConsumeWithGracefulShutdown(c *kafka.Consumer, topics []string) error {
	err := c.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	// 设置信号处理
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 主消费循环
	run := true
	for run {
		select {
		case <-sigchan:
			fmt.Println("Received shutdown signal, closing consumer...")
			run = false
		default:
			msg, err := c.ReadMessage(100 * 1000)
			if err != nil {
				continue
			}

			// 处理消息
			err = processMessage(msg)
			if err != nil {
				fmt.Printf("Failed to process message: %v\n", err)
				continue
			}

			// 手动提交位移
			partitions := []kafka.TopicPartition{
				{Topic: msg.TopicPartition.Topic, Partition: msg.TopicPartition.Partition},
			}
			offsets := []kafka.Offset{
				msg.TopicPartition.Offset + 1,
			}
			c.CommitOffsets(partitions, offsets)
		}
	}

	// 优雅关闭
	fmt.Println("Closing consumer...")
	c.Close()
	fmt.Println("Consumer closed gracefully")

	return nil
}

// 处理再平衡
func ConsumeWithRebalanceListener(c *kafka.Consumer, topics []string) error {
	// 订阅 Topic 时设置 rebalance 回调
	c.SubscribeTopics(topics, func(c *kafka.Consumer, event kafka.Event) error {
		switch e := event.(type) {
		case kafka.AssignedPartitions:
			// 分区分配事件
			fmt.Printf("Assigned partitions: %v\n", e.Partitions)
			// 可以在这里执行一些初始化操作

		case kafka.RevokedPartitions:
			// 分区撤销事件
			fmt.Printf("Revoked partitions: %v\n", e.Partitions)
			// 在分区被撤销前提交位移
			c.Commit()
		}
		return nil
	})

	for {
		msg, err := c.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		processMessage(msg)
	}
}

// 监控消费延迟
func MonitorConsumerLag(brokers string, groupID string) {
	// 使用 AdminClient 获取消费组信息
	// 定期检查 consumer lag
	// 实现取决于具体需求
	fmt.Printf("Monitoring lag for group: %s\n", groupID)
}
```

---

### Day 15：综合项目实战

项目架构：

```sh
订单服务 → Kafka Topic (order-events)
                          ↓
        ┌─────────────────┼─────────────────┐
        ↓                 ↓                 ↓
    库存服务          通知服务         数据分析服务
   (Consumer)        (Consumer)        (Consumer + Streams)
```

#### 15.1 项目结构

```sh
kafka-project/
├── cmd/
│   ├── order-service/          # 订单服务（生产者）
│   │   └── main.go
│   ├── inventory-service/      # 库存服务（消费者）
│   │   └── main.go
│   ├── notification-service/   # 通知服务（消费者）
│   │   └── main.go
│   └── analytics-service/      # 数据分析服务（消费者 + 流处理）
│       └── main.go
├── internal/
│   ├── model/                  # 数据模型
│   │   ├── order.go
│   │   └── order_event.go
│   ├── repository/             # 数据访问层
│   │   └── order_repository.go
│   └── service/                # 业务逻辑层
│       ├── order_service.go
│       ├── inventory_service.go
│       └── notification_service.go
├── pkg/
│   └── kafka/                  # Kafka 工具包
│       ├── producer.go
│       └── consumer.go
├── go.mod
├── go.sum
├── Makefile
└── docker-compose.yml
```

#### 15.2 订单服务（生产者）

**internal/model/order.go**:

```go
package model

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID          string      `json:"id"`
	UserID      string      `json:"user_id"`
	Items       []OrderItem `json:"items"`
	TotalAmount float64     `json:"total_amount"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func (o Order) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Order) CalculateTotal() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
```

**internal/service/order_service.go**:

```go
package service

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"kafka-project/internal/model"
	"kafka-project/internal/repository"
	"time"
)

type OrderService struct {
	orderRepository repository.OrderRepository
	producer       *kafka.Producer
	topic          string
}

func NewOrderService(
	repo repository.OrderRepository,
	producer *kafka.Producer,
	topic string,
) *OrderService {
	return &OrderService{
		orderRepository: repo,
		producer:       producer,
		topic:          topic,
	}
}

func (s *OrderService) CreateOrder(userID string, items []model.OrderItem) (*model.Order, error) {
	// 1. 创建订单
	order := &model.Order{
		ID:        uuid.New().String(),
		UserID:    userID,
		Items:     items,
		Status:    "CREATED",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	order.TotalAmount = order.CalculateTotal()

	// 保存订单
	err := s.orderRepository.Save(order)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	// 2. 发布订单创建事件
	event := model.OrderEvent{
		EventType:  "ORDER_CREATED",
		OrderID:    order.ID,
		UserID:     order.UserID,
		Items:      order.Items,
		TotalAmount: order.TotalAmount,
		Timestamp:  time.Now().UnixMilli(),
	}

	err = s.publishEvent(order.ID, event)
	if err != nil {
		return nil, fmt.Errorf("failed to publish event: %w", err)
	}

	return order, nil
}

func (s *OrderService) ShipOrder(orderID string) (*model.Order, error) {
	// 查询订单
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	// 更新状态
	order.Status = "SHIPPED"
	order.UpdatedAt = time.Now()

	err = s.orderRepository.Save(order)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	// 发布订单发货事件
	event := model.OrderEvent{
		EventType: "ORDER_SHIPPED",
		OrderID:   order.ID,
		UserID:    order.UserID,
		Timestamp: time.Now().UnixMilli(),
	}

	err = s.publishEvent(order.ID, event)
	if err != nil {
		return nil, fmt.Errorf("failed to publish event: %w", err)
	}

	return order, nil
}

func (s *OrderService) CompleteOrder(orderID string) (*model.Order, error) {
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	order.Status = "COMPLETED"
	order.UpdatedAt = time.Now()

	err = s.orderRepository.Save(order)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	// 发布订单完成事件
	event := model.OrderEvent{
		EventType:   "ORDER_COMPLETED",
		OrderID:     order.ID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Timestamp:   time.Now().UnixMilli(),
	}

	err = s.publishEvent(order.ID, event)
	if err != nil {
		return nil, fmt.Errorf("failed to publish event: %w", err)
	}

	return order, nil
}

func (s *OrderService) publishEvent(key string, event model.OrderEvent) error {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event, 10000)

	err = s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.topic, Partition: kafka.PartitionAny},
		Key:           []byte(key),
		Value:         eventJSON,
	}, deliveryChan)

	if err != nil {
		return err
	}

	// 等待发送结果
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	close(deliveryChan)
	return nil
}
```

**cmd/order-service/main.go**:

```go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"kafka-project/internal/repository"
	"kafka-project/internal/service"
)

func main() {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	// 创建 Kafka 生产者
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"acks":               "all",
		"enable.idempotence": true,
		"transactional.id":   "order-service-1",
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	// 初始化事务
	err = producer.InitTransactions(10000)
	if err != nil {
		log.Fatalf("Failed to init transactions: %v", err)
	}

	// 创建服务
	orderRepo := repository.NewInMemoryOrderRepository()
	orderService := service.NewOrderService(orderRepo, producer, "order-events")

	log.Println("Order service started...")

	// 等待信号
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	log.Println("Shutting down order service...")
}
```

#### 15.3 库存服务（消费者）

**cmd/inventory-service/main.go**:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type InventoryConsumer struct {
	consumer         *kafka.Consumer
	inventoryService *InventoryService
	processedEvents  map[string]bool
	mu               sync.RWMutex
}

func NewInventoryConsumer(brokers string, groupID string) (*InventoryConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
		"enable.auto.commit": "false",
	})
	if err != nil {
		return nil, err
	}

	return &InventoryConsumer{
		consumer:         c,
		inventoryService: NewInventoryService(),
		processedEvents:  make(map[string]bool),
	}, nil
}

func (ic *InventoryConsumer) Start(topics []string) error {
	err := ic.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	log.Println("Inventory service started...")

	for {
		msg, err := ic.consumer.ReadMessage(100 * 1000)
		if err != nil {
			continue
		}

		var event OrderEvent
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("Failed to unmarshal event: %v\n", err)
			continue
		}

		// 处理事件
		err = ic.handleOrderEvent(event)
		if err != nil {
			log.Printf("Failed to handle order event: %v\n", err)
			// 不提交位移，允许重试
			continue
		}

		// 手动提交位移
		ic.consumer.Commit()
	}
}

func (ic *InventoryConsumer) handleOrderEvent(event OrderEvent) error {
	// 幂等性检查
	eventKey := event.EventType + ":" + event.OrderID

	ic.mu.RLock()
	_, alreadyProcessed := ic.processedEvents[eventKey]
	ic.mu.RUnlock()

	if alreadyProcessed {
		log.Printf("Duplicate event: %s\n", eventKey)
		return nil
	}

	// 标记为已处理
	ic.mu.Lock()
	ic.processedEvents[eventKey] = true
	ic.mu.Unlock()

	// 处理事件
	switch event.EventType {
	case "ORDER_CREATED":
		return ic.inventoryService.DeductInventory(event)
	case "ORDER_CANCELLED":
		return ic.inventoryService.RestoreInventory(event)
	default:
		log.Printf("Unknown event type: %s\n", event.EventType)
		return nil
	}
}

func (ic *InventoryConsumer) Close() {
	ic.consumer.Close()
}

type InventoryService struct {
	// 库存数据
	Inventory map[string]int
	mu        sync.RWMutex
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		Inventory: map[string]int{
			"product-1": 100,
			"product-2": 50,
		},
	}
}

func (is *InventoryService) DeductInventory(event OrderEvent) error {
	log.Printf("Deducting inventory for order: %s\n", event.OrderID)

	is.mu.Lock()
	defer is.mu.Unlock()

	for _, item := range event.Items {
		currentStock, exists := is.Inventory[item.ProductID]
		if !exists {
			return fmt.Errorf("product not found: %s", item.ProductID)
		}

		if currentStock < item.Quantity {
			return fmt.Errorf("insufficient inventory for product: %s", item.ProductID)
		}

		is.Inventory[item.ProductID] = currentStock - item.Quantity
	}

	log.Printf("Inventory deducted successfully for order: %s\n", event.OrderID)
	return nil
}

func (is *InventoryService) RestoreInventory(event OrderEvent) error {
	log.Printf("Restoring inventory for order: %s\n", event.OrderID)
	// 实现库存恢复逻辑
	return nil
}

func main() {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	consumer, err := NewInventoryConsumer(brokers, "inventory-service-group")
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	// 处理信号
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 启动消费
	go func() {
		err := consumer.Start([]string{"order-events"})
		if err != nil {
			log.Printf("Consumer error: %v\n", err)
		}
	}()

	<-sigchan
	log.Println("Shutting down inventory service...")
}
```

#### 15.6 Docker Compose 配置

```yaml
version: '3.8'

services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.5.0
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181

  kafka:
    image: confluentinc/cp-kafka:7.5.0
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092,PLAINTEXT_HOST://kafka:9093
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1

  order-service:
    build:
      context: .
      dockerfile: cmd/order-service/Dockerfile
    depends_on:
      - kafka
    environment:
      KAFKA_BROKERS: kafka:9093
    restart: unless-stopped

  inventory-service:
    build:
      context: .
      dockerfile: cmd/inventory-service/Dockerfile
    depends_on:
      - kafka
    environment:
      KAFKA_BROKERS: kafka:9093
    restart: unless-stopped

  notification-service:
    build:
      context: .
      dockerfile: cmd/notification-service/Dockerfile
    depends_on:
      - kafka
    environment:
      KAFKA_BROKERS: kafka:9093
    restart: unless-stopped

  analytics-service:
    build:
      context: .
      dockerfile: cmd/analytics-service/Dockerfile
    depends_on:
      - kafka
    environment:
      KAFKA_BROKERS: kafka:9093
    restart: unless-stopped

  kafka-ui:
    image: provectuslabs/kafka-ui:latest
    depends_on:
      - kafka
    ports:
      - "8080:8080"
    environment:
      KAFKA_CLUSTERS_0_NAME: local
      KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: kafka:9093
```

**Makefile**:

```makefile
.PHONY: build run test clean docker-build docker-up docker-down

# 构建所有服务
build:
	@echo "Building all services..."
	@go build -o bin/order-service ./cmd/order-service
	@go build -o bin/inventory-service ./cmd/inventory-service
	@go build -o bin/notification-service ./cmd/notification-service
	@go build -o bin/analytics-service ./cmd/analytics-service

# 运行所有服务（本地）
run: build
	@echo "Running all services..."
	@./bin/order-service &
	@./bin/inventory-service &
	@./bin/notification-service &
	@./bin/analytics-service &

# 运行测试
test:
	@go test -v ./...

# 清理
clean:
	@rm -rf bin/

# Docker 构建
docker-build:
	@docker-compose build

# Docker 启动
docker-up:
	@docker-compose up -d

# Docker 停止
docker-down:
	@docker-compose down

# 查看日志
logs:
	@docker-compose logs -f
```

---

## 学习检查清单

### 第一周检查清单

- [ ] Day 1: 理解 Kafka 核心概念，成功搭建本地环境
- [ ] Day 2: 掌握生产者的三种发送方式及配置
- [ ] Day 3: 理解消费者组，掌握手动提交位移
- [ ] Day 4: 理解再平衡机制，实现再平衡监听器
- [ ] Day 5: 理解幂等性和事务，实现 Exactly-Once
- [ ] Day 6: 了解 Kafka Connect 基础
- [ ] Day 7: 掌握 Kafka Streams 基础操作

### 第二周检查清单

- [ ] Day 8: 成功部署多节点 Kafka 集群
- [ ] Day 9: 理解副本机制和 ISR，配置高可靠 Topic
- [ ] Day 10: 掌握生产者和消费者性能调优
- [ ] Day 11: 配置 Kafka SSL/SASL 安全认证
- [ ] Day 12: 使用监控工具监控 Kafka 集群
- [ ] Day 13: 实现至少一个企业级场景
- [ ] Day 14: 掌握常见问题的解决方案
- [ ] Day 15: 完成综合项目实战

---

## 学习资源

### 官方资源

- [Apache Kafka 官方文档](https://kafka.apache.org/documentation/)
- [Confluent Kafka 文档](https://docs.confluent.io/)
- [Kafka GitHub](https://github.com/apache/kafka)

### 推荐书籍

- 《Kafka 权威指南》- Neha Narkhede 等
- 《深入理解 Kafka：核心设计与实践原理》- 朱忠华
- 《Kafka 技术内幕》- 郭俊

### 在线教程

- [dunwu/bigdata-tutorial Kafka 教程](https://dunwu.github.io/bigdata-tutorial/kafka/)
- [Kafka 中文文档](https://kafka.apachecn.org/)
- [Confluent Kafka 课程](https://developer.confluent.io/)

### 实用工具

- Kafka UI: <https://github.com/provectus/kafka-ui>
- Kafdrop: <https://github.com/obsidiandynamics/kafdrop>
- Kafka-UI: <https://github.com/sourcelaborg/kafka-ui>

---

## 学习建议

1. **每天投入时间**: 建议 2-3 小时
2. **边学边练**: 每个示例代码都要亲自运行
3. **记录笔记**: 整理遇到的问题和解决方案
4. **循序渐进**: 不要跳过基础内容直接学高级特性
5. **实战为主**: 理论学习后立即动手实践

祝你学习顺利！
