package config

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type TimeUnit string

const (
	DefaultTimeFmt   = "2006-01-02 15:04:05.000"
	ConsoleFormatter = "console"
	JsonFormatter    = "json"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
)

var Levels = map[int]zapcore.Level{
	LevelDebug: zapcore.DebugLevel,
	LevelInfo:  zapcore.InfoLevel,
	LevelWarn:  zapcore.WarnLevel,
	LevelError: zapcore.ErrorLevel,
}

// 时间单位配置字段
const (
	Minute = "minute"
	Hour   = "hour"
	Day    = "day"
	Month  = "month"
	Year   = "year"
)

// Format 返回时间单位的格式字符串（c风格），默认返回day的格式字符串
func (t TimeUnit) Format() string {
	switch t {
	case Minute:
		return ".%Y%m%d%H%M"
	case Hour:
		return ".%Y%m%d%H"
	case Day:
		return ".%Y%m%d"
	case Month:
		return ".%Y%m"
	case Year:
		return ".%Y"
	default:
		return ".%Y%m%d"
	}
}

// RotationGap 返回时间单位对应的时间值，默认返回一天的时间
func (t TimeUnit) RotationGap() time.Duration {
	switch t {
	case Minute:
		return time.Minute
	case Hour:
		return time.Hour
	case Day:
		return time.Hour * 24
	case Month:
		return time.Hour * 24 * 30
	case Year:
		return time.Hour * 24 * 365
	default:
		return time.Hour * 24
	}
}
