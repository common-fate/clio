package cliolog

import (
	"encoding/hex"

	"github.com/common-fate/clio/ansi"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

//nolint:gochecknoinits
func init() {
	_ = zap.RegisterEncoder("clio-color", func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return NewConsoleEncoder(&cfg, nil), nil
	})
}

func levelToANSIColor(l zapcore.Level) string {
	if l < zapcore.DebugLevel {
		return dim
	}
	switch l {
	case zapcore.DebugLevel:
		return ansi.LightBlack
	case zapcore.InfoLevel:
		return ansi.White
	case zapcore.WarnLevel:
		return ansi.Yellow
	default:
		return ansi.Red
	}
}

var levelToSymbol = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "i",
	zapcore.WarnLevel:   "!",
	zapcore.ErrorLevel:  "✘",
	zapcore.DPanicLevel: "DPANIC",
	zapcore.PanicLevel:  "PANIC",
	zapcore.FatalLevel:  "FATAL",
}

// SymbolLevelEncoder serializes a Level to a symbol.
// The mapping is as follows:
//
// INFO: [i]
// ERROR: [✘]
// WARN: [!]
// DEBUG: [DEBUG]
func SymbolLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	if s, ok := levelToSymbol[l]; ok {
		enc.AppendString("[" + s + "]")
	}
}

type consoleEncoder struct {
	*ltsvEncoder
	noColor *bool
}

// NewConsoleEncoder creates an encoder whose output is designed for human -
// rather than machine - consumption. It serializes the core log entry data
// (message, level, timestamp, etc.) in a plain-text format.  The context is
// encoded in LTSV.
//
// Note that although the console encoder doesn't use the keys specified in the
// encoder configuration, it will omit any element whose key is set to the empty
// string.
func NewConsoleEncoder(cfg *zapcore.EncoderConfig, noColor *bool) zapcore.Encoder {
	ltsvEncoder := newLTSVEncoder(cfg)
	ltsvEncoder.allowNewLines = true
	ltsvEncoder.allowTabs = true
	ltsvEncoder.blankKey = "value"
	ltsvEncoder.binaryEncoder = hex.Dump

	return &consoleEncoder{ltsvEncoder: ltsvEncoder, noColor: noColor}
}

// Clone implements the Encoder interface
func (c *consoleEncoder) Clone() zapcore.Encoder {
	return &consoleEncoder{
		ltsvEncoder: c.ltsvEncoder.Clone().(*ltsvEncoder),
		noColor:     c.noColor,
	}
}

// dim is the color used for context keys, time, and caller information
var dim = ansi.ColorCode("240")

// EncodeEntry implements the Encoder interface
func (c *consoleEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	final := *c.ltsvEncoder
	context := final.buf
	final.buf = bufPool.Get()

	origLen := final.buf.Len()

	if c.TimeKey != "" {
		c.colorDim(final.buf)
		final.skipNextElementSeparator = true
		c.EncodeTime(ent.Time, &final)
	}

	// color the level symbol and log message based on the level.
	c.colorLevel(final.buf, ent.Level)

	if c.LevelKey != "" {
		final.skipNextElementSeparator = true

		// zap doesn't have a 'success' level, nor does it have an easy way to define
		// custom logging levels.
		//
		// To get around this, we use a special named logger called 'clio.success'.
		// emitting logs to this logger will cause messages to appear in green with a
		// [✔] symbol as the logging level.
		if ent.LoggerName == SuccessName {
			c.applyColor(final.buf, ansi.ColorCode("green"))
			final.buf.AppendString("[✔]")
		} else {
			c.EncodeLevel(ent.Level, &final)
		}
	}

	if final.buf.Len() > origLen {
		final.buf.AppendString(" ")
	} else {
		final.buf.Reset()
	}

	// Add the message itself.
	if c.MessageKey != "" {
		final.safeAddString(ent.Message, false)
		// ensure a minimum of 2 spaces between the message and the fields,
		// to improve readability
		if len(fields) > 0 {
			final.buf.AppendString("  ")
		}
	}

	c.colorDim(final.buf)

	// Add fields.
	for _, f := range fields {
		f.AddTo(&final)
	}

	// Add context
	if context.Len() > 0 {
		final.addFieldSeparator()
		_, _ = final.buf.Write(context.Bytes())
	}

	// If there's no stacktrace key, honor that; this allows users to force
	// single-line output.
	if ent.Stack != "" && c.StacktraceKey != "" {
		final.buf.AppendByte('\n')
		final.buf.AppendString(ent.Stack)
	}
	c.colorReset(final.buf)
	final.buf.AppendByte('\n')

	return final.buf, nil
}

func (c *consoleEncoder) colorDim(buf *buffer.Buffer) {
	c.applyColor(buf, dim)
}

func (c *consoleEncoder) colorLevel(buf *buffer.Buffer, level zapcore.Level) {
	if c.shouldColorize() {
		c.applyColor(buf, levelToANSIColor(level))
	}
}

func (c *consoleEncoder) shouldColorize() bool {
	return c.noColor == nil || !*c.noColor
}

func (c *consoleEncoder) applyColor(buf *buffer.Buffer, s string) {
	if c.shouldColorize() {
		buf.AppendString(ansi.Reset)
		if s != "" {
			buf.AppendString(s)
		}
	}
}

func (c *consoleEncoder) colorReset(buf *buffer.Buffer) {
	c.applyColor(buf, "")
}
