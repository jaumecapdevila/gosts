package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaumecapdevila/gosts/file"
	"github.com/jaumecapdevila/gosts/logger"
)

func TestOperatorShouldAddEntry(t *testing.T) {
	var err error

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	reader := file.NewOSReader()

	f, err := os.Create(tmpFile)

	f.Close()

	if err != nil && !os.IsExist(err) {
		t.Errorf("Unable to create the test file")
		return
	}

	f, err = reader.Read(tmpFile, file.ASSERT)

	if err != nil {
		t.Errorf("Unable to read file '%s'", tmpFile)
	}

	log := logger.NewMockLogger(ctrl)

	operator := file.NewOperator(f, log)

	// Assert entries on file
	err = operator.Assert("entry", "127.0.0.1")

	if err != nil {
		t.Error("Unable to add entry to hosts file...")
	}

	os.Remove(tmpFile)
}

func TestOperatorShouldWarnAboutDuplicatedEntry(t *testing.T) {
	var err error

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	reader := file.NewOSReader()

	f, err := os.Create(tmpFile)

	f.Close()

	if err != nil && !os.IsExist(err) {
		t.Errorf("Unable to create the test file")
		return
	}

	f, err = reader.Read(tmpFile, file.ASSERT)

	defer f.Close()

	if err != nil {
		t.Errorf("Unable to read file '%s'", tmpFile)
	}

	log := logger.NewMockLogger(ctrl)

	operator := file.NewOperator(f, log)

	// Assert entries on file
	err = operator.Assert("entry", "127.0.0.1")

	if err != nil {
		t.Error("Unable to add entry to hosts file...")
	}

	// Assert entries on file
	err = operator.Assert("entry", "127.0.0.1")

	fmt.Println(err)

	if err == nil {
		t.Error("Not detected duplicated entry...")
	}

	os.Remove(tmpFile)
}
