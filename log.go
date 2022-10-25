// Package clio contains helpers for printing CLI output messages.
package clio

import (
	"github.com/common-fate/clio/cliolog"
)

// Log prints to stdout.
func Logf(template string, args ...any) {
	stdoutlog.Printf(template, args...)
}

// Info prints to stderr with an [i] indicator.
func Infof(template string, args ...interface{}) {
	S().Infof(template, args...)
}

// Success prints to stderr with a [✔] indicator.
func Successf(template string, args ...interface{}) {
	S().Named(cliolog.SuccessName).Infof(template, args...)
}

// Error prints to stderr with a [✘] indicator.
func Errorf(template string, args ...any) {
	S().Errorf(template, args...)
}

// Warn prints to stderr with a [!] indicator.
func Warnf(template string, args ...any) {
	S().Warnf(template, args...)
}

// Debug prints to stderr with a [DEBUG] indicator.
// Messages will be shown if the GRANTED_LOG or CF_LOG environment variable is set to 'debug'.
func Debugf(template string, args ...any) {
	S().Debugf(template, args...)
}

// Logln prints to stdout using fmt.Sprintln.
func Logln(args ...any) {
	stdoutlog.Println(args...)
}

// Infoln prints to stderr with an [i] indicator using fmt.Sprintln.
func Infoln(args ...any) {
	S().Infoln(args...)
}

// Successln prints to stderr with a [✔] indicator using fmt.Sprintln.
func Successln(args ...any) {
	S().Named(cliolog.SuccessName).Infoln(args...)
}

// Errorln prints to stderr with a [✘] indicator using fmt.Sprintln.
func Errorln(args ...any) {
	S().Errorln(args...)
}

// Warnln prints to stderr with a [!] indicator using fmt.Sprintln.
func Warnln(args ...any) {
	S().Warnln(args...)
}

// Debugln prints to stderr with a [DEBUG] indicator using fmt.Sprintln.
// Messages will be shown if the GRANTED_LOG or CF_LOG environment variable is set to 'debug'.
func Debugln(args ...any) {
	S().Debugln(args...)
}

// Log uses fmt.Sprint to construct and log a message.
func Log(args ...any) {
	stdoutlog.Print(args...)
}

// Info prints to stderr with an [i] indicator using fmt.Sprint.
func Info(args ...any) {
	S().Info(args...)
}

// Success prints to stderr with a [✔] indicator using fmt.Sprint.
func Success(args ...any) {
	S().Named(cliolog.SuccessName).Info(args...)
}

// Error prints to stderr with a [✘] indicator using fmt.Sprint.
func Error(args ...any) {
	S().Error(args...)
}

// Warn prints to stderr with a [!] indicator using fmt.Sprint.
func Warn(args ...any) {
	S().Warn(args...)
}

// Debug prints to stderr with a [DEBUG] indicator using fmt.Sprint.
// Messages will be shown if the GRANTED_LOG or CF_LOG environment variable is set to 'debug'.
func Debug(args ...any) {
	S().Debug(args...)
}

// Infow prints to stderr with an [i] indicator with additional key-value pairs.
func Infow(msg string, keysAndValues ...any) {
	S().Infow(msg, keysAndValues...)
}

// Success prints to stderr with a [✔] indicator with additional key-value pairs.
func Successw(msg string, keysAndValues ...any) {
	S().Named(cliolog.SuccessName).Infow(msg, keysAndValues...)
}

// Error prints to stderr with a [✘] indicator with additional key-value pairs.
func Errorw(msg string, keysAndValues ...any) {
	S().Errorw(msg, keysAndValues...)
}

// Warn prints to stderr with a [!] indicator with additional key-value pairs.
func Warnw(msg string, keysAndValues ...any) {
	S().Warnw(msg, keysAndValues...)
}

// Debug prints to stderr with a [DEBUG] indicator with additional key-value pairs.
// Messages will be shown if the GRANTED_LOG or CF_LOG environment variable is set to 'debug'.
func Debugw(msg string, keysAndValues ...any) {
	S().Debugw(msg, keysAndValues...)
}
