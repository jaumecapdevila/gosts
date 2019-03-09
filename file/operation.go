package file

import "os"

const (
	// ADD operation flags
	ADD int = os.O_APPEND | os.O_WRONLY

	// REMOVE operation flags
	REMOVE int = os.O_RDWR

	// UPDATE operation flags
	UPDATE int = os.O_RDWR

	// CHECK operation flags
	CHECK int = os.O_RDONLY
)
