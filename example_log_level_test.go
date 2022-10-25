package clio_test

import (
	"os"

	"github.com/common-fate/clio"
	"go.uber.org/zap/zapcore"
)

// You can use `clio.SetLevelFromEnv` to set the log level from environment variables.
// If the environment variables aren't found, this package defaults to the 'info' level.
func Example_levelFromEnv() {
	clio.SetLevelFromEnv("CF_LOG")
	// running CF_LOG=debug <your Go binary> will print debug logs.
}

// You can use `clio.Level.SetLevel()` to set the log level dynamically.
func Example_dynamicLevel() {
	clio.SetErrorWriter(os.Stdout) // print to stdout just to show logs in the example.

	clio.Debug("this isn't printed")
	clio.Level.SetLevel(zapcore.DebugLevel)
	clio.Debug("debug logs now printed!")
	// Output: [DEBUG] debug logs now printed!
}

// You can use `clio.SetLevelFromString()` to set the log level dynamically from a provided string.
func Example_levelFromString() {
	clio.SetErrorWriter(os.Stdout) // print to stdout just to show logs in the example.

	clio.Debug("this isn't printed")
	clio.SetLevelFromString("debug")
	clio.Debug("debug logs now printed!")
	// Output: [DEBUG] debug logs now printed!
}
