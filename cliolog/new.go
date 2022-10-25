package cliolog

import (
	"io"
	"time"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Writer  io.Writer
	NoColor *bool
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
