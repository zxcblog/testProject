package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"new-project/pkg/config"
)

func NewLogger(log config.Logger) *zap.Logger {
	return InitLogger(log.DebugFile, log.MaxSize, log.MaxAge, zap.ErrorLevel)
}

func NewAccessLogger(log config.Logger) io.Writer {
	return &lumberjack.Logger{
		Filename:  log.AccessFile, // 日志文件位置
		MaxSize:   log.MaxSize,    // 日志文件的最大大小（MB）
		MaxAge:    log.MaxAge,     // 保留旧文件的最大天数
		LocalTime: true,
	}
}

func InitLogger(filepath string, maxSize, maxAge int, level zapcore.Level) *zap.Logger {
	writeSyncer := getLogWriter(filepath, maxSize, maxAge)
	encoder := getEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// 如果打印的日志级别不是info级别，将函数调用信息记录到日志中
	if level != zapcore.InfoLevel {
		return zap.New(core, zap.AddCaller())
	}
	return zap.New(core)
}

func getEncoder(options zapcore.EncoderConfig) zapcore.Encoder {
	options.EncodeTime = zapcore.RFC3339TimeEncoder
	options.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(options)
}

func getLogWriter(filepath string, maxSize, maxAge int) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:  filepath, // 日志文件位置
		MaxSize:   maxSize,  // 日志文件的最大大小（MB）
		MaxAge:    maxAge,   // 保留旧文件的最大天数
		LocalTime: true,
	})
}
