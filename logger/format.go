package logger

const (
	// JSON logging format
	JSON Format = "json"

	// TEXT logging format
	TEXT Format = "text"
)

// Format is format to log messages
type Format string
