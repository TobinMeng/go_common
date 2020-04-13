package appender

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_common/log/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultFileMaxSize    = 500
	DefaultFileMaxBackups = 10
	DefaultFileMaxAge     = 1000
	DefaultCompress       = false
)

type RollingFileAppender struct {
	Formatter  string
	TimeFmt    string
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func NewRollingFileAppender(formatter string, timeFmt string, fileName string, maxSize int, maxBackups int, maxAge int, compress bool) *RollingFileAppender {
	return &RollingFileAppender{
		Formatter:  formatter,
		TimeFmt:    timeFmt,
		FileName:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

func (r RollingFileAppender) GetCore(level int) zapcore.Core {
	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   r.FileName,
		MaxSize:    r.MaxSize,
		MaxBackups: r.MaxBackups,
		MaxAge:     r.MaxAge,
		Compress:   r.Compress,
	})
	encoder := getEncoder(r.Formatter, r.TimeFmt)
	lvl := zap.NewAtomicLevelAt(config.Levels[level])
	return zapcore.NewCore(encoder, ws, lvl)
}
