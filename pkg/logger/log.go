package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"io"
	"new-chat/config"
	"sync"
)

var (
	once         sync.Once
	AccessLogger io.Writer
	Logger       *zap.Logger
)

func NewLogger(log config.Logger) {
	Logger = InitLogger(log.DebugFile, log.MaxSize, log.MaxAge, zap.ErrorLevel)
}

func NewAccessLogger(log config.Logger) {
	AccessLogger = &lumberjack.Logger{
		Filename:  log.AccessFile, // 日志文件位置
		MaxSize:   log.MaxSize,    // 日志文件的最大大小（MB）
		MaxAge:    log.MaxAge,     // 保留旧文件的最大天数
		LocalTime: true,
	}
}

// GetAccessLogWriter 项目启动时会自动加载，不用手动二次加载
func GetAccessLogWriter() io.Writer {
	return AccessLogger
}
