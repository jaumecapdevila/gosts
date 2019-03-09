package logger

import "io"

// Logger interface
type Logger interface {
	Info(Context, string)
	Error(Context, string)
	Fatal(Context, string)
	Warning(Context, string)
}

// New generates and returns a new logger depending the format specified
func New(format Format, writers []io.Writer) Logger {
	switch format {
	case JSON:
		return NewJSON(io.MultiWriter(writers...))
	case TEXT:
		return NewText(io.MultiWriter(writers...))
	default:
		panic("Invalid logger format")
	}
}
