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

	me.write(current)
}

// Error logs given payload with INFO level
func (me *JSONLogger) Error(context Context, message string) {
	current := (&Context{
		"level":   ErrorLevel,
		"message": message,
	}).Append(&context)

	me.write(current)
	os.Exit(0)
}

// Fatal logs given payload with ERROR level and stops the application
func (me *JSONLogger) Fatal(context Context, message string) {
	current := (&Context{
		"level":   ErrorLevel,
		"message": message,
	}).Append(&context)

	me.write(current)
	os.Exit(-1)
}

func (me *JSONLogger) write(data *Context) {
	if data == nil {
		data = &Context{}
	}

	json.NewEncoder(me.writer).Encode(data)
}

// NewJSON instantiate a new JSON format logger
func NewJSON(writer io.Writer) *JSONLogger {
	return &JSONLogger{writer: writer}
}
