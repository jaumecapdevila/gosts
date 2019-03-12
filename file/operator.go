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

	// EntryFoundMessage defines the message displayed if the entry is found
	EntryFoundMessage string = "Entry '%s' found on line '%d'"
)

// Operator is able to execute read/write operations over a file
type Operator struct {
	File   *os.File
	Logger logger.Logger
}

// Find specified entry
func (h *Operator) Find(entry Entry) (int, error) {
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

	_, err = h.File.WriteString(fmt.Sprintf(EntryFormat, entry.Domain, entry.Address))

	if err != nil {
		return 0, err
	}

	if _, err = h.File.Seek(0, 0); err != nil {
		return 0, err
	}

	return 0, nil
}

// Assert specified entry
func (h *Operator) Assert(entry string, address string) error {
	scanner := bufio.NewScanner(h.File)

	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), entry) {
			return NewExistentEntryError(entry, line)
		}

		line++
	}

	var err error

	if err = scanner.Err(); err != nil {
		return err
	}

	_, err = h.File.WriteString(fmt.Sprintf(EntryFormat, entry, address))

	if err != nil {
		return err
	}

	if _, err = h.File.Seek(0, 0); err != nil {
		return err
	}

	return nil
}

// Remove file entry
func (h *Operator) Remove(entry string) error {
	scanner := bufio.NewScanner(h.File)

	var buffer bytes.Buffer

	for scanner.Scan() {

		content := scanner.Text()

		if !strings.Contains(content, entry) {
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
