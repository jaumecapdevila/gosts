package file

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/jaumecapdevila/gosts/logger"
)

const (
	// EntryFormat defines the format for the file entries
	EntryFormat string = "%s %s\n"
)

// Operator is able to execute read/write operations over a file
type Operator struct {
	File   *os.File
	Logger logger.Logger
}

// Find entry on file
func (h *Operator) Find(entry *Entry) (int, error) {
	scanner := bufio.NewScanner(h.File)

	line := 1

	for scanner.Scan() {
		lineContent := scanner.Text()

		if strings.Contains(lineContent, entry.Address) || strings.Contains(lineContent, entry.Domain) {
			return line, nil
		}

		line++
	}

	var err error

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	if _, err = h.File.Seek(0, 0); err != nil {
		return 0, err
	}

	return 0, NewEntryNotFoundError(entry.Domain)
}

// Create entry on file
func (h *Operator) Create(entry *Entry) error {
	scanner := bufio.NewScanner(h.File)

	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), entry.Domain) {
			return NewExistentEntryError(entry.Domain, line)
		}

		line++
	}

	var err error

	if err = scanner.Err(); err != nil {
		return err
	}

	_, err = h.File.WriteString(fmt.Sprintf(EntryFormat, entry, entry.Address))

	if err != nil {
		return err
	}

	if _, err = h.File.Seek(0, 0); err != nil {
		return err
	}

	return nil
}

// Remove file entry
func (h *Operator) Remove(entry *Entry) error {
	scanner := bufio.NewScanner(h.File)

	var buffer bytes.Buffer

	for scanner.Scan() {

		content := scanner.Text()

		if !strings.Contains(content, entry.Domain) {
			buffer.WriteString(content)
			buffer.WriteString("\n")
		}
	}

	var err error

	if err = scanner.Err(); err != nil {
		return err
	}

	if err = h.File.Truncate(0); err != nil {
		return err
	}

	if _, err = h.File.Seek(0, 0); err != nil {
		return err
	}

	if _, err = h.File.Write(buffer.Bytes()); err != nil {
		return err
	}

	return nil
}

// NewOperator returns a new file operator
func NewOperator(file *os.File, logger logger.Logger) *Operator {
	return &Operator{File: file, Logger: logger}
}
