package logger

import (
	"encoding/json"
	"io"
	"os"
)

// JSONLogger is a json implementation of the logger interface
type JSONLogger struct {
	Logger
	writer io.Writer
}

// Info logs given payload with INFO level
func (me *JSONLogger) Info(context Context, message string) {
	current := (&Context{
		"level":   InfoLevel,
		"message": message,
	}).Append(&context)

	me.write(GREEN, current)
}

// Error logs given payload with INFO level
func (me *JSONLogger) Error(context Context, message string) {
	current := (&Context{
		"level":   ErrorLevel,
		"message": message,
	}).Append(&context)
	me.write(RED, current)
	os.Exit(0)
}

// Warning logs given payload with INFO level
func (me *JSONLogger) Warning(context Context, message string) {
	current := (&Context{
		"level":   WarningLevel,
		"message": message,
	}).Append(&context)

	me.write(YELLOW, current)
}

// Fatal logs given payload with ERROR level and stops the application
func (me *JSONLogger) Fatal(context Context, message string) {
	current := (&Context{
		"level":   ErrorLevel,
		"message": message,
	}).Append(&context)
	me.write(RED, current)
	os.Exit(-1)
}

func (me *JSONLogger) write(color string, data *Context) {
	if data == nil {
		data = &Context{}
	}

	me.writer.Write([]byte(color))

	json.NewEncoder(me.writer).Encode(data)
}

// NewJSON instantiate a new JSON format logger
func NewJSON(writer io.Writer) *JSONLogger {
	return &JSONLogger{writer: writer}
}
