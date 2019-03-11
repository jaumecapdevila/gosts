package file

import "fmt"

// ExistentEntryError format
const existentEntryErrorMessage string = "The entry '%s' already exists on line '%d'"

// ExistentEntryError definition
type ExistentEntryError struct {
	message string
}

func (e *ExistentEntryError) Error() string {
	return e.message
}

// NewExistentEntryError constructor
func NewExistentEntryError(entry string, line int) error {
	return &ExistentEntryError{
		message: fmt.Sprintf(existentEntryErrorMessage, entry, line),
	}
}
