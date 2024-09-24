package consumers

import (
	"context"
	"log"
	"time"

	kafkago "github.com/segmentio/kafka-go"
	"lessons/kafka/pkg/message"
)

// Consumer определяет интерфейс для Kafka-консьюмера
type Consumer interface {
	Consume(ctx context.Context) error
	TopicName() string
	Close()
}

// HandlerFunc - это функция для обработки сообщений
type HandlerFunc func(ctx context.Context, msg message.Message) error

// NewConsumer создает простой Kafka-консьюмер
func NewConsumer(groupID, topic string, brokers []string, startOffset *int64, handler HandlerFunc) Consumer {
	if startOffset == nil {
		startOffset = new(int64)
		*startOffset = kafkago.FirstOffset
	}

	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: *startOffset,
		MaxWait:     10 * time.Millisecond, // сколько мы готовы ждать сообщений в одной итерации
	})
	return &consumer{
		Reader:  reader,
		handler: handler,
		topic:   topic,
	}
}

// Реализация консьюмера
type consumer struct {
	*kafkago.Reader
	handler HandlerFunc
	topic   string
}

func (c *consumer) Consume(ctx context.Context) error {
	msg, err := c.Reader.ReadMessage(ctx)
	if err != nil {
		return err
	}

	// Логика обработки сообщения через переданный хендлер
	if err := c.handler(ctx, message.NewFromRaw(msg)); err != nil {
		log.Printf("error handling message: %v", err)
		return err
	}

	// Коммит сообщения после успешной обработки
	if err := c.Reader.CommitMessages(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (c *consumer) TopicName() string {
	return c.topic
}

func (c *consumer) Close() {
	if err := c.Reader.Close(); err != nil {
		log.Printf("error closing consumer: %v", err)
	}
}
