package main

type Logger struct {
	name  string
	level string
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) WithName(name string) *Logger {
	l.name = name
	return l
}

func (l *Logger) WithLevel(level string) *Logger {
	l.level = level
	return l
}

func main() {
	logger1 := NewLogger()
	logger2 := NewLogger().WithLevel("INFO")
	logger3 := NewLogger().WithName("storage")

	_ = logger1
	_ = logger2
	_ = logger3
}
