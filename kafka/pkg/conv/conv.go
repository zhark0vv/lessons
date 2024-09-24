package conv

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	kafkatypes "lessons/kafka/gen/go/zhark0vv/kafka/types"
	"lessons/kafka/internal/domain"
)

func ConvertLogEventToProto(event domain.LogEvent) *kafkatypes.LogEvent {
	return &kafkatypes.LogEvent{
		CreatedAt: timestamppb.New(event.CreatedAt),
		Level:     convertLogLevelToProto(event.Level),
		Message:   event.Message,
	}
}

func ConvertLogEventFromProto(event *kafkatypes.LogEvent) domain.LogEvent {
	return domain.LogEvent{
		CreatedAt: event.CreatedAt.AsTime(),
		Level:     domain.Level(event.Level),
		Message:   event.Message,
	}
}

func convertLogLevelToProto(level domain.Level) kafkatypes.LogLevel {
	switch level {
	case domain.InfoLevel:
		return kafkatypes.LogLevel_LOG_LEVEL_INFO
	case domain.ErrorLevel:
		return kafkatypes.LogLevel_LOG_LEVEL_ERROR
	case domain.DebugLevel:
		return kafkatypes.LogLevel_LOG_LEVEL_DEBUG
	case domain.WarnLevel:
		return kafkatypes.LogLevel_LOG_LEVEL_WARNING
	case domain.CritLevel:
		return kafkatypes.LogLevel_LOG_LEVEL_CRITICAL
	default:
		return kafkatypes.LogLevel_LOG_LEVEL_UNSPECIFIED
	}
}
