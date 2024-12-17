package main

import (
	"fmt"
	logger2 "logger/logger"
	"logger/logger/logrus"
	"logger/logger/std"
	"logger/logger/zap"
)

func main() {
	loggers := []logger2.Logger{
		logrus.New("", "json", ""),
		std.New("", "json"),
		zap.New("", "json", ""),
	}
	for _, logger := range loggers {
		fmt.Printf("Usando a implementacao do logger %s \n", logger.GetLoggerName())
		logger.Info("teste Info")
		logger.Warn("teste Warn")
		logger.Debug("teste Debug")
		logger.Error("teste Error")
	}
}
