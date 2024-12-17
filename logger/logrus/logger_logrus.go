package logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"logger/logger"
	"os"
)

type logrusLogger struct {
	logger *logrus.Entry
}

func New(level string, format string, output string) logger.Logger {
	log := logrus.New()

	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	if format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

	if output == "" {
		log.SetOutput(os.Stdout)
	} else {
		file, err := os.Create(output)
		if err != nil {
			panic("falha ao criar arquivo de log")
		}
		log.SetOutput(file)
	}

	return &logrusLogger{logger: log.WithFields(logrus.Fields{})}
}

func (l *logrusLogger) GetLoggerName() string         { return "Logrus" }
func (l *logrusLogger) Debug(msg string, args ...any) { l.logger.Debugf(msg, args...) }
func (l *logrusLogger) Info(msg string, args ...any)  { l.logger.Infof(msg, args...) }
func (l *logrusLogger) Warn(msg string, args ...any)  { l.logger.Warnf(msg, args...) }
func (l *logrusLogger) Error(msg string, args ...any) { l.logger.Errorf(msg, args...) }
func (l *logrusLogger) Fatal(msg string, args ...any) { l.logger.Fatalf(msg, args...) }
func (l *logrusLogger) Fatalf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.logger.Fatal(msg)
}

func (l *logrusLogger) With(args ...any) logger.Logger {
	fields := logrus.Fields{}
	for i := 0; i < len(args)-1; i += 2 {
		if key, ok := args[i].(string); ok {
			fields[key] = args[i+1]
		}
	}
	return &logrusLogger{logger: l.logger.WithFields(fields)}
}
