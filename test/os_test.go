package test

import (
	"os"
	"testing"

	"github.com/jaumecapdevila/gosts/file"
)

func TestReaderShouldSeeFileContent(t *testing.T) {
	var err error

	reader := file.NewOSReader()

	f, err := os.Create(tmpFile)

	f.Close()

	if err != nil && !os.IsExist(err) {
		t.Errorf("Unable to create the test file")
		return
	}

	_, err = reader.Read(tmpFile, file.ASSERT)

	if err != nil {
		t.Errorf("Unable to read file '%s'", tmpFile)
	}

	os.Remove(tmpFile)
}
