package appender

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_common/log/config"
	"os"
)

type ConsoleAppender struct {
	Formatter string
	TimeFmt   string
}

func NewConsoleAppender(formatter string, timeFmt string) *ConsoleAppender {
	if timeFmt == "" {
		timeFmt = config.DefaultTimeFmt
	}
	if formatter == "" {
		formatter = config.ConsoleFormatter
	}

	return &ConsoleAppender{Formatter: formatter, TimeFmt: timeFmt}
}

func (c *ConsoleAppender) GetCore(level int) zapcore.Core {
	encoder  := getEncoder(c.Formatter,c.TimeFmt)
	lvl := zap.NewAtomicLevelAt(config.Levels[level])
	return zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), lvl)
}
