package file

import "os"

// OSReader implementation
type OSReader struct{}

func (ior *OSReader) Read(file string, mode int) (*os.File, error) {
	f, err := os.OpenFile(file, mode, 0755)
	return f, err
}

// NewOSReader creates a new file reader implementation
func NewOSReader() *OSReader {
	return &OSReader{}
}
