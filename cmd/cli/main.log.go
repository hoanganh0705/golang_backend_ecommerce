package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 2 ways to log a message with zap

func main() {
	// 1
	// sugar logger
	// sugar := zap.NewExample().Sugar()
	// sugar.Infow("Hello %s, %d", "Hoang Anh", 20)

	// logger ( use key-value pairs )
	// logger := zap.NewExample()
	// logger.Info("hello", zap.String("name", "Hoang Anh"), zap.Int("age", 20))

	// when to use which one ?

	// 2
	// logger := zap.NewExample()
	// logger.Info("Hello")

	// //development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development ")

	// //production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production ")

	// 3. Custom
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

}

// format/ customize log output
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

func getWriterSync() zapcore.WriteSyncer {
	logDir := "pkg/logger"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logDir+"/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
