// Package clio contains helpers for printing CLI output messages.
package clio

import "strings"

// NewLine prints a newline to clio.OutputWriter
func NewLine() {
	stdout.Println()
}

// NewLine prints n newline(s) to clio.OutputWriter
func NewLines(n int) {
	stdout.Printf(strings.Repeat("\n", n))
}
