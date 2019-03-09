package file

import "os"

// Reader defines contract to read a file
type Reader interface {
	Read(file string, mode int) (*os.File, error)
}
