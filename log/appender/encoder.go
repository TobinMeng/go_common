package appender

import (
	"github.com/TobinMeng/go_common/log/config"
	"go.uber.org/zap/zapcore"
	"time"
)

func getEncoder(formatter string, timeFmt string) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		MessageKey:    "M",
		StacktraceKey: "S",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format(timeFmt))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder

	switch formatter {
	case config.ConsoleFormatter:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case config.JsonFormatter:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	return encoder
}
