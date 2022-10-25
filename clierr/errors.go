package clierr

import (
	"github.com/common-fate/clio"
)

type Printer interface {
	Print()
}

type PrintCLIErrorer interface {
	PrintCLIError()
}

// Err is a CLI error.
// Calling PrintCLIError() on it will print error messages to the console.
type Err struct {
	// Err is a string to avoid puntuation linting
	Err string
	// setting ExcludeDefaultError to true will cause the PrintCLIError method to only print the contents of Messages
	// by default PrintCLIError first prints the contents of Err, followed by the contents of Messages
	ExcludeDefaultError bool
	// Messages are items which implement the clio.Printer interface.
	// when PrintCLIError is called, each of the messages Print() method is called in order of appearence in the slice
	Messages []Printer
}

// New creates a new CLI error. You can append additional log messages to the error by adding fields to the 'msgs' argument.
//
// Example:
//
//	clierr.New("something bad happened", clierr.Error("some extra context here"))
func New(err string, msgs ...Printer) *Err {
	return &Err{
		Err:      err,
		Messages: msgs,
	}
}

type msgtype uint8

const (
	log msgtype = iota
	info
	success
	// errort is error, but that is a reserved word in Go
	errort
	warn
	debug

	logf
	infof
	successf
	errorf
	warnf
	debugf
)

type Msg struct {
	Format  string
	Args    []any
	msgtype msgtype
}

func (m Msg) Print() {
	switch m.msgtype {
	case info:
		clio.Info(m.Args...)
	case success:
		clio.Success(m.Args...)
	case errort:
		clio.Error(m.Args...)
	case warn:
		clio.Warn(m.Args...)
	case debug:
		clio.Debug(m.Args...)
	case infof:
		clio.Infof(m.Format, m.Args...)
	case successf:
		clio.Successf(m.Format, m.Args...)
	case errorf:
		clio.Errorf(m.Format, m.Args...)
	case warnf:
		clio.Warnf(m.Format, m.Args...)
	case debugf:
		clio.Debugf(m.Format, m.Args...)
	default:
		clio.Logf(m.Format, m.Args...)
	}
}

// Logf adds a formatted log message. Warning: this will be printed to stdout, rather than stderr.
func Logf(format string, a ...any) Msg {
	return Msg{Format: format, Args: a, msgtype: logf}
}

// Infof adds a formatted info message.
func Infof(format string, a ...any) Msg {
	return Msg{Format: format, Args: a, msgtype: infof}
}

// Successf adds a formatted success message.
func Successf(format string, a ...any) Msg {
	return Msg{Format: format, Args: a, msgtype: successf}
}

// Errorf adds a formatted error message.
func Errorf(format string, a ...any) Msg {
	return Msg{Format: format, Args: a, msgtype: errorf}
}

// Warnf adds a formatted warning message.
func Warnf(format string, a ...any) Msg {
	return Msg{Format: format, Args: a, msgtype: warnf}
}

// Debugf adds a formatted debug message. This message will only be displayed if CF_LOG or GRANTED_LOG is set to 'debug'.
func Debugf(format string, a ...any) Msg {
	return Msg{Format: format, Args: a, msgtype: debugf}
}

// Logf adds a log message. Warning: this will be printed to stdout, rather than stderr.
func Log(args ...any) Msg {
	return Msg{Args: args, msgtype: logf}
}

// Infof adds an info message.
func Info(args ...any) Msg {
	return Msg{Args: args, msgtype: info}
}

// Successf adds a success message.
func Success(args ...any) Msg {
	return Msg{Args: args, msgtype: success}
}

// Errorf adds a error message.
func Error(args ...any) Msg {
	return Msg{Args: args, msgtype: errort}
}

// Warnf adds a warning message.
func Warn(args ...any) Msg {
	return Msg{Args: args, msgtype: warn}
}

// Debugf adds a debug message. This message will only be displayed if CF_LOG or GRANTED_LOG is set to 'debug'.
func Debug(args ...any) Msg {
	return Msg{Args: args, msgtype: debug}
}

// Error implements the error interface. It uses the default message of the
// wrapped error.
func (e *Err) Error() string {
	return e.Err
}

// PrintCLIError prints the error message and then any messages in order from the slice
// The indended use is to surface errors with useful messages then os.Exit without having to place os.Exit within methods other than the cli main function
//
//	err := CLIError{Err: "new error", Messages: []Printer{Log("hello world")}}
//
//	err.PrintCLIError()
//	// produces
//	[✘] new error
//	hello world
func (e *Err) PrintCLIError() {
	if !e.ExcludeDefaultError {
		clio.Error(e.Err)
	}

	for i := range e.Messages {
		e.Messages[i].Print()
	}
}
