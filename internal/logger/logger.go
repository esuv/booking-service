package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	logger := log.Default()
	return &Logger{logger}
}

func (l *Logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.Printf("[Error]: %s\n", msg)
}

func (l *Logger) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.Printf("[Info]: %s\n", msg)
}
