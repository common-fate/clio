package cliolog

import (
	"io"
	"time"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Options struct {
	Writer          io.Writer
	NoColor         *bool
	FileWriteSyncer *zapcore.WriteSyncer
}

// New returns a CLI-friendly zap logger which prints to stderr by default.
func New(level zap.AtomicLevel, opts ...func(*Options)) *zap.Logger {
	o := Options{
		Writer: colorable.NewColorableStderr(),
	}

	for _, opt := range opts {
		opt(&o)
	}

	ec := zap.NewDevelopmentEncoderConfig()
	ec.EncodeLevel = SymbolLevelEncoder

	// if fileWriteSyncer is present then write logs to file as well as showing to console.
	if o.FileWriteSyncer != nil {
		fec := zap.NewProductionEncoderConfig()
		fec.EncodeTime = zapcore.TimeEncoder(func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.UTC().Format("2023-01-02T15:04:05Z0700"))
		})
		fileEncoder := zapcore.NewJSONEncoder(fec)

		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, zapcore.AddSync(*o.FileWriteSyncer), level), zapcore.NewCore(NewConsoleEncoder(&ec, o.NoColor), zapcore.AddSync(o.Writer), level))

		return zap.New(core)
	}

	// no-op time encoder, by default.
	ec.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {}

	log := zap.New(zapcore.NewCore(
		NewConsoleEncoder(&ec, o.NoColor),
		zapcore.AddSync(o.Writer),
		level,
	))

	return log
}

// WithWriter specifies an io.Writer to write logs to.
func WithWriter(w io.Writer) func(*Options) {
	return func(o *Options) {
		o.Writer = w
	}
}

// WithNoColor sets up a colorization bypass.
func WithNoColor(noColor *bool) func(*Options) {
	return func(o *Options) {
		o.NoColor = noColor
	}
}

type FileLoggerConfig struct {
	Filename       string
	MaxSize        *int
	MaxBackups     *int
	MaxAge         *int
	ShouldCompress *bool
}

func WithFileLogger(cfg FileLoggerConfig) func(*Options) {
	loggerCfg := lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    1,
		MaxBackups: 30,
		MaxAge:     90,
		Compress:   false,
	}

	if cfg.MaxSize != nil {
		loggerCfg.MaxSize = *cfg.MaxSize
	}

	if cfg.MaxAge != nil {
		loggerCfg.MaxAge = *cfg.MaxAge
	}

	if cfg.MaxBackups != nil {
		loggerCfg.MaxBackups = *cfg.MaxBackups
	}

	if cfg.ShouldCompress != nil {
		loggerCfg.Compress = *cfg.ShouldCompress
	}

	ws := zapcore.AddSync(&loggerCfg)

	return func(o *Options) {
		o.FileWriteSyncer = &ws
	}
}
