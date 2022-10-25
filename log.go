// Package clio contains helpers for printing CLI output messages.
package clio

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	// ErrorWriter defaults to stderr
	ErrorWriter = color.Error
	// OutputWriter defaults to stdout
	OutputWriter = color.Output
)

var stderr = log.New(ErrorWriter, "", 0)
var stdout = log.New(OutputWriter, "", 0)

// Log prints to stdout.
func Log(format string, a ...interface{}) {
	stdout.Printf(format, a...)
}

// Info prints to stderr with an [i] indicator.
func Info(format string, a ...interface{}) {
	format = "[i] " + format
	stderr.Printf(color.WhiteString(format, a...))
}

// Success prints to stderr with a [✔] indicator.
func Success(format string, a ...interface{}) {
	format = "[✔] " + format
	stderr.Printf(color.GreenString(format, a...))
}

// Error prints to stderr with a [✘] indicator.
func Error(format string, a ...interface{}) {
	format = "[✘] " + format
	stderr.Printf(color.RedString(format, a...))
}

// Warn prints to stderr with a [!] indicator.
func Warn(format string, a ...interface{}) {
	format = "[!] " + format
	stderr.Printf(color.YellowString(format, a...))
}

// Warn prints to stderr with a [DEBUG] indicator
// if the GRANTED_LOG environment variable is set to 'debug'.
func Debug(format string, a ...interface{}) {
	if isDebug() {
		format = "[DEBUG] " + format
		stderr.Printf(color.HiBlackString(format, a...))
	}
}

func isDebug() bool {
	return strings.ToLower(os.Getenv("GRANTED_LOG")) == "debug" || strings.ToLower(os.Getenv("CF_LOG")) == "debug"
}
