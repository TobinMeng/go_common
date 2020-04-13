package appender

import (
	"github.com/TobinMeng/go_common/log/config"
	rotates "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type DailyRollingFileAppender struct {
	Formatter string
	TimeFmt   string
	FileName  string
	MaxAge    int
	TimeUnit  config.TimeUnit
}

func NewDailyRollingFileAppender(formatter string, timeFmt string, maxAge int, timeUnit config.TimeUnit) *DailyRollingFileAppender {
	return &DailyRollingFileAppender{
		Formatter: formatter,
		TimeFmt:   timeFmt,
		MaxAge:    maxAge,
		TimeUnit:  timeUnit,
	}
}

func (d *DailyRollingFileAppender) GetCore(level int) zapcore.Core {
	hook, _ := rotates.New(
		d.FileName+d.TimeUnit.Format(),
		rotates.WithMaxAge(time.Duration(int64(24*time.Hour)*int64(d.MaxAge))),
		rotates.WithRotationTime(d.TimeUnit.RotationGap()),
	)
	ws := zapcore.AddSync(hook)
	encoder := getEncoder(d.Formatter, d.TimeFmt)
	lvl := zap.NewAtomicLevelAt(config.Levels[level])
	return zapcore.NewCore(encoder, ws, lvl)
}
