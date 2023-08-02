package clio

import (
	"io"
	"os"
	"sync"

	"github.com/common-fate/clio/cliolog"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetLevelFromEnv configures the global logging level based on the provided
// environment variables.
// The env vars should be provided in priority order.
// It returns a function which restores the previous log level.
func SetLevelFromEnv(vars ...string) func() {
	for _, e := range vars {
		val := os.Getenv(e)
		lvl, err := zapcore.ParseLevel(val)
		if err == nil {
			old := Level.Level()
			Level.SetLevel(lvl)
			return func() { Level.SetLevel(old) }
		}
	}
	// if we get here, we couldn't parse any env vars.
	// Return a no-op func as we did nothing.
	return func() {}
}

// SetLevelFromString configures the global logging level based on the provided
// string. Under the hood it uses zapcore.Parse() to try and parse the log level.
// Does nothing if the log level can't be parsed.
// It returns a function which restores the previous log level.
func SetLevelFromString(level string) func() {
	lvl, err := zapcore.ParseLevel(level)
	if err == nil {
		old := Level.Level()
		Level.SetLevel(lvl)
		return func() { Level.SetLevel(old) }
	}

	// if we get here, we couldn't parse the level.
	// Return a no-op func as we did nothing.
	return func() {}
}

var (
	// Level is the global logging level.
	Level = zap.NewAtomicLevel()

	// globalMu locks concurrent access to the global loggers.
	globalMu sync.RWMutex

	// errorWriter defaults to stderr
	errorWriter = colorable.NewColorableStderr()

	// stderr is a zap logger which writes to stderr
	stderr = cliolog.New(
		Level,
		cliolog.WithWriter(errorWriter),
		cliolog.WithNoColor(&NoColor),
	).Sugar()
)

// SetWriter rebuilds the global zap logger with a specific writer.
// All Info, Error, Warn, Debug, etc messages are sent here.
// clio.Log messages are sent to stdout.
func SetWriter(w io.Writer) {
	globalMu.Lock()
	defer globalMu.Unlock()

	stderr = cliolog.New(
		Level,
		cliolog.WithWriter(w),
		cliolog.WithNoColor(&NoColor),
	).Sugar()
}

func SetFileLogging(fcfg cliolog.FileLoggerConfig) {
	stderr = cliolog.New(
		Level,
		cliolog.WithFileLogger(fcfg),
		cliolog.WithWriter(errorWriter),
		cliolog.WithNoColor(&NoColor),
	).Sugar()
}

// G returns the global stderr logger
// as a zap logger.
func G() *zap.Logger {
	globalMu.RLock()
	s := stderr.Desugar()
	globalMu.RUnlock()
	return s
}

// S returns the global stderr logger
// as a sugared zap logger.
func S() *zap.SugaredLogger {
	globalMu.RLock()
	s := stderr
	globalMu.RUnlock()
	return s
}

// ReplaceGlobals replaces the global Logger and SugaredLogger, and returns a
// function to restore the original values. It's safe for concurrent use.
func ReplaceGlobals(logger *zap.Logger) func() {
	globalMu.Lock()
	prev := stderr.Desugar()
	stderr = logger.Sugar()
	globalMu.Unlock()
	return func() { ReplaceGlobals(prev) }
}
