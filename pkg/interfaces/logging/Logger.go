package logging

var LogLevels = loggerLevel{
	Info:  1,
	Warn:  2,
	Debug: 3,
	Error: 4,
}

type Logger interface {
	Info(s string)
	Warn(s string)
	Debug(s string)
	Error(s string)
	SetLevel(int uint8)
}

type loggerLevel struct {
	Info  uint8
	Warn  uint8
	Debug uint8
	Error uint8
}
