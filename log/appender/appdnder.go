package appender

import (
	"go.uber.org/zap/zapcore"
)

type Appender interface {
	GetCore(level int)  zapcore.Core
}

