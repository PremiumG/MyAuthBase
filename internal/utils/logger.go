package utils

import (
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var L *slog.Logger

func init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if AppConfig.Debug {
		opts.Level = slog.LevelDebug
	}

	var handler slog.Handler
	if AppConfig.Debug {
		// Development: readable text
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		// Production: structured JSON
		logFileName := AppConfig.LogDir + "/log.log"
		logWriter := &lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    10, // MB
			MaxBackups: 5,  // files
			MaxAge:     30, // days
			Compress:   true,
		}
		handler = slog.NewJSONHandler(logWriter, opts)
	}

	L = slog.New(handler)
}
