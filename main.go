package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jaumecapdevila/gosts/file"
	"github.com/jaumecapdevila/gosts/logger"
)

func main() {
	var hostsFile string
	var entry string
	var address string
	var remove bool

	flag.StringVar(&hostsFile, "file", "/etc/hosts", "The user hosts file")
	flag.StringVar(&entry, "entry", "", "Entry to add to the hostsfile")
	flag.StringVar(&address, "address", "127.0.0.1", "Address for the entry")
	flag.BoolVar(&remove, "d", false, "Indicate that entry must be removed instead")

	flag.Parse()

	reader := file.NewOSReader()

	log := logger.New(logger.JSON, []io.Writer{os.Stdout})

	if entry == "" {
		log.Fatal(logger.Context{}, "You must specify a valid entry")
	}

	mode := file.ADD

	if remove {
		mode = file.REMOVE
	}

	f, err := reader.Read(hostsFile, mode)

	if err != nil {
		log.Error(logger.Context{}, "Unable to read file")
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(file.FORMAT, entry, address))

	if err != nil {
		log.Error(logger.Context{}, "Unable to write to file")
	}

	os.Exit(0)

}
