package cliolog

import (
	"io"
	"os"
	"time"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	// LevelEnvVars will try and load log level from the provided environment
	// variables. The slice defines the priority to load environment variables from.
	// The first element in the slice is the highest priority.
	LevelEnvVars []string
	Writer       io.Writer
	NoColor      *bool
}

func (o Options) LogLevel() zapcore.Level {
	for _, e := range o.LevelEnvVars {
		val := os.Getenv(e)
		lvl, err := zapcore.ParseLevel(val)
		if err == nil {
			return lvl
		}
	}

	// return info level by default
	return zapcore.InfoLevel
}

// New returns a CLI-friendly zap logger which prints to stderr by default.
func New(opts ...func(*Options)) *zap.Logger {
	o := Options{
		Writer: colorable.NewColorableStderr(),
	}

	for _, opt := range opts {
		opt(&o)
	}

	level := o.LogLevel()

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

// WithLevelEnvVars specifies environment variables to load log levels from.
// The variables are in priority order, with the first variable being the
// highest priority.
func WithLevelEnvVars(vars ...string) func(*Options) {
	return func(o *Options) {
		o.LevelEnvVars = vars
	}
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
