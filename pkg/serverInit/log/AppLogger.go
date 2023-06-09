package log

import (
	"contentservice/pkg/interfaces/logging"
	"fmt"
	"log"
	"sync"
)

type Logger struct {
	log   *log.Logger
	level uint8
	mu    sync.RWMutex
}

func (l *Logger) Info(s string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level == logging.LogLevels.Info {
		log.Println(formatLog("INFO", s))
	}
}

func (l *Logger) Warn(s string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level == logging.LogLevels.Warn || l.level == logging.LogLevels.Error {
		log.Println(formatLog("WARN", s))
	}
}

func (l *Logger) Debug(s string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level == logging.LogLevels.Debug || l.level == logging.LogLevels.Warn || l.level == logging.LogLevels.Error {
		log.Println(formatLog("DEBUG", s))
	}
}

func (l *Logger) Error(s string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level == logging.LogLevels.Warn || l.level == logging.LogLevels.Error {
		log.Println(formatLog("ERROR", s))
	}
}

func formatLog(level string, message string) string {
	return fmt.Sprintf("%s: %s", level, message)
}

func (l *Logger) SetLevel(level uint8) {
	l.mu.Lock()
	defer l.mu.Unlock()

	switch level {
	case logging.LogLevels.Info:
		l.level = logging.LogLevels.Info
		break

	case logging.LogLevels.Debug:
		l.level = logging.LogLevels.Debug
		break

	case logging.LogLevels.Warn:
		l.level = logging.LogLevels.Warn
		break

	case logging.LogLevels.Error:
		l.level = logging.LogLevels.Error
		break

	default:
		l.level = logging.LogLevels.Warn
	}
}

func newLogger() logging.Logger {
	l := Logger{
		level: logging.LogLevels.Warn,
	}
	l.log = new(log.Logger)
	return &l
}
