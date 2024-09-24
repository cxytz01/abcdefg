package utils

import (
	"context"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var KafkaReader *kafka.Reader

func NewKafkaReader(ctx context.Context, cf Kafka) *kafka.Reader {
	brokers := strings.Split(cf.BootstrapServer, ",")

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          cf.Topic,
		CommitInterval: 3 * time.Second,
		GroupID:        cf.ConsumerGroup,
	})
}
