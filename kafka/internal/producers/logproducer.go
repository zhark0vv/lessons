package producers

import (
	"context"
	"fmt"

	"lessons/kafka/internal/domain"
	"lessons/kafka/pkg/conv"
	"lessons/kafka/pkg/message"
	producer "lessons/kafka/pkg/producers"
)

type LogProducer struct {
	p producer.Producer
}

func NewLogProducer(p producer.Producer) *LogProducer {
	return &LogProducer{p: p}
}

func (lp *LogProducer) ProduceLogMessage(ctx context.Context, ev domain.LogEvent) error {
	proto := conv.ConvertLogEventToProto(ev)

	msg, err := message.FromProto(proto)
	if err != nil {
		return fmt.Errorf("failed to convert log event to message: %w", err)
	}

	return lp.p.WriteMessage(ctx, *msg)
}
