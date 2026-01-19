package config

import "time"

// Config Kafka 配置
type Config struct {
	// BootstrapServers Kafka 集群地址
	BootstrapServers string

	// Topic 默认 Topic
	Topic string

	// GroupID 消费者组 ID
	GroupID string
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		BootstrapServers: "localhost:9092",
		Topic:           "quickstart-events",
		GroupID:         "demo-group",
	}
}

// ProducerConfig 生产者配置
type ProducerConfig struct {
	*Config

	// Acks 确认级别: 0, 1, all
	Acks string

	// CompressionType 压缩类型: none, gzip, snappy, lz4, zstd
	CompressionType string

	// LingerMs 等待时间（毫秒）
	LingerMs int

	// BatchSize 批量大小（字节）
	BatchSize int

	// RetryTimes 重试次数
	RetryTimes int

	// Timeout 超时时间
	Timeout time.Duration
}

// DefaultProducerConfig 返回默认生产者配置
func DefaultProducerConfig() *ProducerConfig {
	return &ProducerConfig{
		Config:          DefaultConfig(),
		Acks:            "all",
		CompressionType: "snappy",
		LingerMs:        10,
		BatchSize:       16384,
		RetryTimes:      3,
		Timeout:         30 * time.Second,
	}
}

// ConsumerConfig 消费者配置
type ConsumerConfig struct {
	*Config

	// AutoOffsetReset 位移重置策略: earliest, latest, none
	AutoOffsetReset string

	// EnableAutoCommit 是否自动提交位移
	EnableAutoCommit bool

	// AutoCommitIntervalMs 自动提交间隔（毫秒）
	AutoCommitIntervalMs int

	// SessionTimeoutMs 会话超时（毫秒）
	SessionTimeoutMs int

	// HeartbeatIntervalMs 心跳间隔（毫秒）
	HeartbeatIntervalMs int
}

// DefaultConsumerConfig 返回默认消费者配置
func DefaultConsumerConfig() *ConsumerConfig {
	return &ConsumerConfig{
		Config:               DefaultConfig(),
		AutoOffsetReset:      "earliest",
		EnableAutoCommit:     true,
		AutoCommitIntervalMs: 5000,
		SessionTimeoutMs:     30000,
		HeartbeatIntervalMs:  3000,
	}
}
