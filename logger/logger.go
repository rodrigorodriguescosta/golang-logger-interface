package logger

type Logger interface {
	Fatal(msg string, args ...any)
	Fatalf(format string, args ...any) // Novo método
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
	GetLoggerName() string
}
