package domain

import (
	"time"
)

type Level string

const (
	InfoLevel  Level = "INFO"
	ErrorLevel Level = "ERROR"
	DebugLevel Level = "DEBUG"
	WarnLevel  Level = "WARN"
	CritLevel  Level = "CRIT"
)

func (l Level) Ptr() *Level {
	return &l
}

type LogEvent struct {
	CreatedAt time.Time
	Level     Level
	Message   string
}
