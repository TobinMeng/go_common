package log

import (
	"go_common/log/appender"
)


// Logger 日志器
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})

	SetLevel(level int)
	GetLevel() int

	Init()
	AddAppender(appender appender.Appender)
}
