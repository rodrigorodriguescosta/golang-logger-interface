package std

import (
	"fmt"
	"log"
	"logger/logger"
	"os"
)

type stdLogger struct {
	logger *log.Logger
}

func New(prefix string, output string) logger.Logger {
	var out *os.File
	if output == "" {
		out = os.Stdout
	} else {
		var err error
		out, err = os.Create(output)
		if err != nil {
			panic("falha ao criar arquivo de log")
		}
	}

	return &stdLogger{
		logger: log.New(out, prefix, log.LstdFlags|log.Lshortfile),
	}
}

func (l *stdLogger) GetLoggerName() string         { return "Log padrao" }
func (l *stdLogger) Debug(msg string, args ...any) { l.logger.Printf("[DEBUG] "+msg, args...) }
func (l *stdLogger) Info(msg string, args ...any)  { l.logger.Printf("[INFO] "+msg, args...) }
func (l *stdLogger) Warn(msg string, args ...any)  { l.logger.Printf("[WARN] "+msg, args...) }
func (l *stdLogger) Error(msg string, args ...any) { l.logger.Printf("[ERROR] "+msg, args...) }
func (l *stdLogger) Fatal(msg string, args ...any) { l.logger.Fatalf("[FATAL] "+msg, args...) }
func (l *stdLogger) Fatalf(format string, args ...any) {
	l.logger.Fatalf(format, args...)
}
func (l *stdLogger) With(args ...any) logger.Logger {
	// Não há suporte nativo para campos contextuais no logger padrão.
	// Simulamos criando uma nova instância com um prefixo.
	newPrefix := fmt.Sprintf("%v ", args...)
	return &stdLogger{
		logger: log.New(os.Stdout, newPrefix, log.LstdFlags|log.Lshortfile),
	}
}
