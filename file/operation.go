package file

import "os"

const (
	// ASSERT operation flags
	ASSERT int = os.O_APPEND | os.O_RDWR

	// REMOVE operation flags
	REMOVE int = os.O_RDWR
)
