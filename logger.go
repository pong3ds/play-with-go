package main

// ILogger is interface for logger
type ILogger interface {
	Log(msg string)
}

// Logger is struct for logger
type Logger struct {
}

// NewLogger return new logger
func NewLogger() *Logger {
	return &Logger{}
}

// Log will log the message
func (l *Logger) Log(msg string) {
	//
}
