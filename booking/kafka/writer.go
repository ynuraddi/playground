package kafka

import (
	"github.com/segmentio/kafka-go"
)

func NewWriter(cfg Config) *kafka.Writer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      cfg.Brokers,
		RequiredAcks: 0,
	})
	return w
}
