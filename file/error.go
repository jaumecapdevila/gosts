package file

import "fmt"

const (
	existentEntryErrorMessage    string = "The entry '%s' already exists on line '%d'"
	nonExistingEntryErrorMessage string = "Entry '%s' for address '%s' does not exist..."
)

// ExistentEntryError definition
type ExistentEntryError struct {
	message string
}

func (e *ExistentEntryError) Error() string {
	return e.message
}

// EntryNotFoundError definition
type EntryNotFoundError struct {
	message string
}

func (e *EntryNotFoundError) Error() string {
	return e.message
}

// NewExistentEntryError constructor
func NewExistentEntryError(entry string, line int) error {
	return &ExistentEntryError{
		message: fmt.Sprintf(existentEntryErrorMessage, entry, line),
	}
}

// NewEntryNotFoundError constructor
func NewEntryNotFoundError(entry *Entry) error {
	return &EntryNotFoundError{
		message: fmt.Sprintf(nonExistingEntryErrorMessage, entry.Domain, entry.Address),
	}
}
