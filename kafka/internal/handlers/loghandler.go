package handlers

import (
	"context"
	"fmt"
	"log"

	kafkatypes "lessons/kafka/gen/go/zhark0vv/kafka/types"
	"lessons/kafka/pkg/message"
)

type LogHandler struct{}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

func (l *LogHandler) Handle(_ context.Context, msg message.Message) error {
	logMsg := kafkatypes.LogEvent{}

	err := msg.AsProto(&logMsg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	log.Printf("Date: %s, Level: %s, Message: %s",
		logMsg.CreatedAt.AsTime(), logMsg.Level, logMsg.Message)

	return nil
}
