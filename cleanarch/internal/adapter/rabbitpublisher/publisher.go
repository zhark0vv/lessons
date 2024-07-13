package rabbitpublisher

import (
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQEventPublisher struct {
	channel  *amqp.Channel
	exchange string
}

func NewRabbitMQEventPublisher(channel *amqp.Channel, exchange string) *RabbitMQEventPublisher {
	return &RabbitMQEventPublisher{
		channel:  channel,
		exchange: exchange,
	}
}

func (p *RabbitMQEventPublisher) Publish(_ context.Context, event interface{}) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		p.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf("Event published: %v", event)
	return nil
}
