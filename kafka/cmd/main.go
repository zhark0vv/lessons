package main

import (
	"context"

	"lessons/kafka/internal/app"
	"lessons/kafka/internal/handlers"
	internalProducers "lessons/kafka/internal/producers"
	pkgConsumers "lessons/kafka/pkg/consumers"
	pkgProducers "lessons/kafka/pkg/producers"
)

func main() {
	ctx := context.Background()

	logHandler := handlers.NewLogHandler()

	logConsumer := pkgConsumers.NewConsumer(
		"my-group",
		"logs_events",
		[]string{"localhost:9092"},
		nil,
		logHandler.Handle,
	)

	logProducer := internalProducers.NewLogProducer(
		pkgProducers.NewProducer(
			[]string{"kafka:9092"},
			"logs_events",
			nil,
		))

	kafkaApp := app.NewKafkaApp(logProducer, logConsumer)
	kafkaApp.Run(ctx)
}
