package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
	"lessons/kafka/internal/domain"
	internalProducers "lessons/kafka/internal/producers"
	"lessons/kafka/pkg/consumers"
)

type KafkaApp struct {
	c []consumers.Consumer
	p *internalProducers.LogProducer
}

func NewKafkaApp(p *internalProducers.LogProducer, consumers ...consumers.Consumer) *KafkaApp {
	return &KafkaApp{c: consumers, p: p}
}

func (a *KafkaApp) Run(ctx context.Context) {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		a.testProduce(a.p)
		return nil
	})

	for _, c := range a.c {
		consumer := c

		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					consumer.Close()
					return ctx.Err()
				default:
					err := consumer.Consume(ctx)
					if err != nil {
						return fmt.Errorf("failed to consume message: %w", err)
					}
				}
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.Panic(err)
	}
}

func (a *KafkaApp) testProduce(p *internalProducers.LogProducer) {
	i := 0
	logLevel := domain.InfoLevel

	for {
		i++
		if i%2 == 0 {
			logLevel = domain.ErrorLevel
		}

		if i%3 == 0 {
			logLevel = domain.DebugLevel
		}

		if i%5 == 0 {
			logLevel = domain.CritLevel
		}

		err := p.ProduceLogMessage(
			context.Background(),
			domain.LogEvent{
				CreatedAt: time.Now(),
				Level:     logLevel,
				Message:   fmt.Sprintf("Hello, world! #%d", i),
			},
		)

		if err != nil {
			panic(err)
		}
	}
}
