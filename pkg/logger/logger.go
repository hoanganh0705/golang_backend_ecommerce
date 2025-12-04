package logger

import (
	"GolangBackendEcommerce/pkg/settings"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config settings.LogSetting) *LoggerZap {
	logLevel := config.Log_level
	// debug -> info -> warn -> error -> dpanic -> panic -> fatal
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}
	// format log entries as JSON with specific formatting rules.
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.Log_file_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), level)

	return &LoggerZap{
		zap.New(core, zap.AddCaller()),
	}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1764122298.7574997 -> 2025-11-26T08:58:18.757+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//ts -> time
	encodeConfig.TimeKey = "time"

	//from info to INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:27"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
