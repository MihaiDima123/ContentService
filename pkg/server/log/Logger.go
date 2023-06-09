package log

import "contentservice/pkg/interfaces/logging"

var logger logging.Logger

func InitLogger() {
	logger = newLogger()
}

func Info(m string) {
	logger.Info(m)
}

func Debug(m string) {
	logger.Debug(m)
}

func Warn(m string) {
	logger.Warn(m)
}

func Error(m string) {
	logger.Error(m)
}

func SetLevel(level uint8) {
	logger.SetLevel(level)
}
