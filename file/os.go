package file

import "os"

// OSReader implementation
type OSReader struct{}

func (ior *OSReader) Read(file string, mode int) (*os.File, error) {
	return os.OpenFile(file, mode, 0600)
}

// NewOSReader creates a new file reader implementation
func NewOSReader() *OSReader {
	return &OSReader{}
}
