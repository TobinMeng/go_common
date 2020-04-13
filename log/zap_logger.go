package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_common/log/appender"
)

type ZapLogger struct {
	appenders []appender.Appender
	logger    *zap.Logger
	Level     int
}

func NewZapLogger(level int) *ZapLogger {
	return &ZapLogger{Level: level}
}

func (z *ZapLogger) Init() {
	var cores = make([]zapcore.Core, 0)
	for _, appender := range z.appenders {
		core :=appender.GetCore(z.Level)
		cores = append(cores,core)
	}
	z.logger = zap.New(zapcore.NewTee(cores...), zap.AddCallerSkip(2),zap.AddCaller())
}

func (z *ZapLogger) Debug(format string, args ...interface{}) {
	z.logger.Debug(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) Info(format string, args ...interface{}) {
	z.logger.Info(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) Warning(format string, args ...interface{}) {
	z.logger.Warn(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) Error(format string, args ...interface{}) {
	z.logger.Error(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) SetLevel(level int) {
	panic("implement me")
}

func (z *ZapLogger) GetLevel() int {
	panic("implement me")
}

func (z *ZapLogger) AddAppender(appender appender.Appender) {
	z.appenders = append(z.appenders, appender)
}
