package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jaumecapdevila/gosts/logger"
)

const (
	// EntryFormat defines the format for the file entries
	EntryFormat string = "%s %s"

	// ExistentEntryError format
	ExistentEntryError string = "The entry %s already exists on file..."
)

// Handler is able to execute read/write operations over a file
type Handler struct {
	File   *os.File
	Logger logger.Logger
}

// Assert that the entry is on the file
func (h *Handler) Assert(entry string, address string) error {
	scanner := bufio.NewScanner(h.File)

	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), entry) {
			h.Logger.Info(logger.Context{}, fmt.Sprint(ExistentEntryError, entry))
			return nil
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

	return nil
}

func NewHandler(file *os.File, logger logger.Logger) *Handler {
	return &Handler{File: file, Logger: logger}
}
