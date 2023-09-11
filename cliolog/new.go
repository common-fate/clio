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
	// no-op time encoder, by default.
	ec.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {}

	// if fileWriteSyncer is present then write logs to file as well as showing to console.
	if o.FileWriteSyncer != nil {
		fec := zap.NewProductionEncoderConfig()
		fec.EncodeTime = zapcore.TimeEncoder(func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
		})
		fileEncoder := zapcore.NewJSONEncoder(fec)

		core := zapcore.NewTee(zapcore.NewCore(fileEncoder, zapcore.AddSync(*o.FileWriteSyncer), level), zapcore.NewCore(NewConsoleEncoder(&ec, o.NoColor), zapcore.AddSync(o.Writer), level))

		return zap.New(core)
	}

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
	// Name of your log file
	Filename string
	// Maximum size of the log file in MB before it gets rotated.
	// If set to 0 the default max size is used (1 MB)
	MaxSize int
	// Maximum number of old log files to retain.
	// If set to 0 the default is 30
	MaxBackups int
	// Maximum number of days to retain old log files based on the timestamp encoded in their filename.
	// If set to 0 the default is 90 days
	MaxAge int

	// ShouldCompress if set to true will compress the rotated log files using gzip.
	ShouldCompress bool
}

// WithFileLogger will write logs to a file using lumberjack package in addition to printing it in console.
func WithFileLogger(cfg FileLoggerConfig) func(*Options) {
	loggerCfg := lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    1,
		MaxBackups: 30,
		MaxAge:     90,
		Compress:   cfg.ShouldCompress,
	}

	if cfg.MaxAge != 0 {
		loggerCfg.MaxAge = cfg.MaxAge
	}

	if cfg.MaxSize != 0 {
		loggerCfg.MaxSize = cfg.MaxSize
	}

	if cfg.MaxBackups != 0 {
		loggerCfg.MaxBackups = cfg.MaxBackups
	}

	ws := zapcore.AddSync(&loggerCfg)

	return func(o *Options) {
		o.FileWriteSyncer = &ws
	}
}
