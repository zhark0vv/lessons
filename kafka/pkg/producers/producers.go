package producer

import (
	"context"
	"crypto/tls"
	"log"

	kafkago "github.com/segmentio/kafka-go"
	"lessons/kafka/pkg/message"
)

// Producer интерфейс для Kafka продьюсера
type Producer interface {
	WriteMessage(ctx context.Context, msg message.Message) error
	Close() error
}

// producer структура для работы с Kafka
type producer struct {
	writer *kafkago.Writer
}

// NewProducer создает новый Kafka продьюсер с конфигурацией
func NewProducer(brokers []string, topic string, tlsCfg *tls.Config) Producer {
	// Конфигурация writer
	writer := &kafkago.Writer{
		Addr:  kafkago.TCP(brokers...),
		Topic: topic,
		Balancer: &kafkago.RoundRobin{
			ChunkSize: 1,
		},
	}

	if tlsCfg != nil {
		writer.Transport = &kafkago.Transport{
			TLS: tlsCfg,
		}
	}

	log.Printf("Created Kafka producer for topic %s", topic)

	return &producer{
		writer: writer,
	}
}

// WriteMessage отправляет сообщение в Kafka
func (p *producer) WriteMessage(ctx context.Context, msg message.Message) error {
	err := p.writer.WriteMessages(ctx, msg.Raw())
	if err != nil {
		log.Printf("Failed to write message: %v", err)
		return err
	}

	log.Printf("Message sent with key %s", string(msg.Key))
	return nil
}

// Close закрывает продьюсер
func (p *producer) Close() error {
	if err := p.writer.Close(); err != nil {
		log.Printf("Failed to close producer: %v", err)
		return err
	}
	return nil
}
