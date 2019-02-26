package file

import "os"

// OSReader implementation
type OSReader struct{}

func (ior *OSReader) Read(file string) (*os.File, error) {
	f, err := os.OpenFile(file, os.O_RDWR, 0755)
	return f, err
}
