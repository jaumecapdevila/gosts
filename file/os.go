package file

import "os"

// OSReader implementation
type OSReader struct{}

func (ior *OSReader) Read(file string) (*os.File, error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0755)
	return f, err
}

// NewOSReader creates a new file reader implementation
func NewOSReader() *OSReader {
	return &OSReader{}
}
