package file

import "os"

// Reader defines contract to read a file
type Reader interface {
	Read(string) (*os.File, error)
}
