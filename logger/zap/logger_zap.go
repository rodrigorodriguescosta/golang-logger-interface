package zap

import (
	"fmt"
	"logger/logger"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func New(level string, format string, output string) logger.Logger {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	var encoder zapcore.Encoder
	if format == "json" {
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}

	var syncer zapcore.WriteSyncer
	if output == "" {
		syncer = zapcore.AddSync(os.Stdout)
	} else {
		file, err := os.Create(output)
		if err != nil {
			panic("falha ao criar arquivo de log")
		}
		syncer = zapcore.AddSync(file)
	}

	core := zapcore.NewCore(encoder, syncer, zapLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &zapLogger{logger: logger.Sugar()}
}

func (l *zapLogger) GetLoggerName() string         { return "ZAP Logger" }
func (l *zapLogger) Debug(msg string, args ...any) { l.logger.Debugw(msg, args...) }
func (l *zapLogger) Info(msg string, args ...any)  { l.logger.Infow(msg, args...) }
func (l *zapLogger) Warn(msg string, args ...any)  { l.logger.Warnw(msg, args...) }
func (l *zapLogger) Error(msg string, args ...any) { l.logger.Errorw(msg, args...) }
func (l *zapLogger) Fatal(msg string, args ...any) { l.logger.Fatalw(msg, args...) }
func (l *zapLogger) Fatalf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...) // Formatar a mensagem
	l.logger.Fatalw(msg)                // Logar e encerrar
}
func (l *zapLogger) With(args ...any) logger.Logger {
	sugar := l.logger.With(args...)
	return &zapLogger{logger: sugar}
}
