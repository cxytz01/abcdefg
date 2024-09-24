package utils

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer

func NewKafkaWriter(ctx context.Context, cf Kafka) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(cf.BrokerList),
		Topic:                  cf.Topic,
		RequiredAcks:           kafka.RequiredAcks(cf.Ack),
		Balancer:               &kafka.Hash{},
		WriteTimeout:           2 * time.Second,
		AllowAutoTopicCreation: false,
	}
}
