package clio

import (
	"os"

	"github.com/mattn/go-isatty"
)

// NoColor defines if the output is colorized or not. It's dynamically set to
// false or true based on the stdout's file descriptor referring to a terminal
// or not. It's also set to true if the NO_COLOR environment variable is
// set (regardless of its value). This is a global option and affects all
// colors. For more control over each color block use the methods
// DisableColor() individually.
var NoColor = noColorExists() || os.Getenv("TERM") == "dumb" ||
	(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

// noColorExists returns true if the environment variable NO_COLOR exists.
func noColorExists() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}
