package _zap

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOption func(*zap.Config, *[]zap.Option) error

// -1 debug
//  0 info
//  1 warn
func WithLoggingLevel(level int) LoggerOption {
	return func(cfg *zap.Config, _ *[]zap.Option) error {
		cfg.Level = zap.NewAtomicLevelAt(zapcore.Level(int8(level)))
		return nil
	}
}

func WithEncoded(encode string) LoggerOption {
	return func(cfg *zap.Config, _ *[]zap.Option) error {
		switch encode {
		case "json", "console":
		default:
			return fmt.Errorf("unexpected encode %s", encode)
		}
		cfg.Encoding = encode
		return nil
	}
}

func WithColor() LoggerOption {
	return func(cfg *zap.Config, _ *[]zap.Option) error {
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return nil
	}
}

func WithAddStacktrace(level int) LoggerOption {
	return func(_ *zap.Config, opts *[]zap.Option) error {
		*opts = append(*opts, zap.AddStacktrace((toZapLevel(level))))
		return nil
	}
}

func WithAddCaller() LoggerOption {
	return func(_ *zap.Config, opts *[]zap.Option) error {
		*opts = append(*opts, zap.AddCaller())
		return nil
	}
}

func NewLogger(options ...LoggerOption) (*zap.Logger, error) {
	cfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(int8(-1))),
		Development:      true,
		Encoding:         "console", // or json
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	var buildOpts = []zap.Option{}
	for _, opt := range options {
		if err := opt(cfg, &buildOpts); err != nil {
			return nil, err
		}
	}

	logger, err := cfg.Build(buildOpts...)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func toZapLevel(level int) zap.AtomicLevel {
	return zap.NewAtomicLevelAt(zapcore.Level(int8(level)))
}
