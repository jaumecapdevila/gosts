package file

import (
	"fmt"
	"os"
)

// Operation of the file
type Operation uint

const (
	// Find entry on file
	Find Operation = 1

	// Create entry on file
	Create Operation = 2

	// Remove entry of the file
	Remove Operation = 3
)

// GetFlagForOperation returns the proper flag to open the file
func GetFlagForOperation(operation Operation) (int, error) {
	switch operation {
	case Find:
		return os.O_RDONLY, nil
	case Create:
		return os.O_APPEND | os.O_RDWR, nil
	case Remove:
		return os.O_RDWR, nil
	default:
		return 0, fmt.Errorf("The operation '%d' is not valid", operation)
	}
}
