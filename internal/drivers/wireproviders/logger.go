package wireproviders

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	callerKey     = "caller"
	levelKey      = "level"
	messageKey    = "message"
	nameKey       = "name"
	stacktraceKey = "stacktrace"
	timestampKey  = "timestamp"
)

func NewLoggerFromConfig(conf zap.Config) (*zap.Logger, error) {
	return conf.Build(zap.AddCallerSkip(1))
}

func NewLoggerConfig(conf LoggerConfig) (zap.Config, error) {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(conf.Level)); err != nil {
		return zap.Config{}, err
	}

	return zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        timestampKey,
			LevelKey:       levelKey,
			MessageKey:     messageKey,
			NameKey:        nameKey,
			StacktraceKey:  stacktraceKey,
			CallerKey:      callerKey,
			LineEnding:     "\n",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		},
	}, nil
}
