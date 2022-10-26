package clio

import "go.uber.org/zap/zapcore"

// IsDebug returns true if clio.Level is set to debug or lower.
func IsDebug() bool {
	return Level.Enabled(zapcore.DebugLevel)
}
