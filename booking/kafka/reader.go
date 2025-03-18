package kafka

import "github.com/segmentio/kafka-go"

func NewReader(cfg Config, topic, groupId string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return r
}
