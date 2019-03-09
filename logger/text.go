package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ContextPrintFormat for output
const ContextPrintFormat = "%s:%s;"

// TextLogger is a json implementation of the logger interface
type TextLogger struct {
	Logger
	writer io.Writer
}

// Info logs given payload with INFO level
func (me *TextLogger) Info(context Context, message string) {
	current := (&Context{
		"level":   InfoLevel,
		"message": message,
	}).Append(&context)

	me.write(GREEN, current)
}

// Error logs given payload with INFO level
func (me *TextLogger) Error(context Context, message string) {
	current := (&Context{
		"level":   ErrorLevel,
		"message": message,
	}).Append(&context)
	me.write(RED, current)
	os.Exit(0)
}

// Warning logs given payload with INFO level
func (me *TextLogger) Warning(context Context, message string) {
	current := (&Context{
		"level":   WarningLevel,
		"message": message,
	}).Append(&context)

	me.write(YELLOW, current)
	os.Exit(0)
}

// Fatal logs given payload with ERROR level and stops the application
func (me *TextLogger) Fatal(context Context, message string) {
	current := (&Context{
		"level":   ErrorLevel,
		"message": message,
	}).Append(&context)
	me.write(RED, current)
	os.Exit(-1)
}

func (me *TextLogger) write(color string, data *Context) {
	if data == nil {
		data = &Context{}
	}

	var buffer bytes.Buffer

	buffer.WriteString(color)

	for key, value := range *data {
		buffer.WriteString(" ")
		buffer.WriteString(fmt.Sprintf(ContextPrintFormat, key, value))
	}

	buffer.WriteString("\n")

	me.writer.Write(buffer.Bytes())
}

// NewText instantiate a new Text format logger
func NewText(writer io.Writer) *TextLogger {
	return &TextLogger{writer: writer}
}
