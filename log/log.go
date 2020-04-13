package log

import (
	"go_common/log/appender"
	"go_common/log/config"
)

var defaultLogger Logger

func init() {
	defaultLogger = NewZapLogger(config.LevelInfo)
	consoleAppender := appender.NewConsoleAppender(config.ConsoleFormatter,config.DefaultTimeFmt)
	//_, fileName, _, _ := runtime.Caller(5)
	rollingFileAppender := appender.NewRollingFileAppender(
		config.ConsoleFormatter,
		config.DefaultTimeFmt,
		"111.log",
		appender.DefaultFileMaxSize,
		appender.DefaultFileMaxBackups,
		appender.DefaultFileMaxAge,
		appender.DefaultCompress,
	)
	defaultLogger.AddAppender(consoleAppender)
	defaultLogger.AddAppender(rollingFileAppender)
	defaultLogger.Init()
}

// Debug log
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format,args...)
}

// Info log
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format,args...)
}

// Warning log
func Warning(format string, args ...interface{}) {
	defaultLogger.Warning(format,args...)
}

// Error log
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format,args...)
}
