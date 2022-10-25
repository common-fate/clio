package clio

import (
	"strings"
)

// NewLine prints a newline to clio.ErrorWriter
func NewLine() {
	stderrlog.Println()
}

// NewLines prints n newline(s) to clio.ErrorWriter
func NewLines(n int) {
	stderrlog.Printf(strings.Repeat("\n", n))
}
