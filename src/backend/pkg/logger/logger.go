package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

type Log struct {
	serviceName string
}

type LogEntry struct {
	ServiceName string    `json:"serviceName"`
	Level       string    `json:"level"`
	Timestamp   time.Time `json:"timestamp"`
	Message     string    `json:"msg"`
	Data        any       `json:"data,omitempty"`
}

func New(serviceName string) *Log {
	return &Log{
		serviceName: serviceName,
	}
}

func (l *Log) log(level, msg string, data any) {
	entry := LogEntry{
		ServiceName: l.serviceName,
		Level:       level,
		Timestamp:   time.Now(),
		Message:     msg,
		Data:        data,
	}

	b, err := json.Marshal(entry)
	if err != nil {
		fmt.Println("failed to marshal log:", err)
		return
	}

	fmt.Println(string(b))
}

func (l *Log) Info(msg string, data any)  { l.log("info", msg, data) }
func (l *Log) Warn(msg string, data any)  { l.log("warn", msg, data) }
func (l *Log) Error(msg string, data any) { l.log("error", msg, data) }
func (l *Log) Debug(msg string, data any) { l.log("debug", msg, data) }
