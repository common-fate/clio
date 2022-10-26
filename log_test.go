package clio

import (
	"bytes"
	"testing"
)

func TestInfo(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	NoColor = true

	Info("my message")

	got := b.String()

	want := "[i] my message\n"
	if got != want {
		t.Errorf("Info() = %q, want %q", got, want)
	}
}

func TestError(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	NoColor = true

	Error("my message")

	got := b.String()

	want := "[✘] my message\n"
	if got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestLog(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	NoColor = true

	Log("my message")

	got := b.String()

	// should print without a prefix.
	want := "my message\n"
	if got != want {
		t.Errorf("Log() = %q, want %q", got, want)
	}
}
