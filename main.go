package main

import (
	"flag"
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

	log := logger.New(logger.TEXT, []io.Writer{os.Stdout})

	if entry == "" {
		log.Fatal(nil, "You must specify a valid entry")
	}

	mode := file.ASSERT

	if remove {
		mode = file.REMOVE
	}

	f, err := reader.Read(hostsFile, mode)

	if err != nil {
		log.Error(nil, "Unable to read file")
		os.Exit(0)
	}

	defer f.Close()

	operator := file.NewOperator(f, log)

	if remove {
		err = operator.Remove(entry)

		if err != nil {
			log.Error(nil, "Unable to remove the specified entry")
			os.Exit(0)
		}

		log.Info(nil, "Entry removed successfully!!!")
	} else {
		err = operator.Assert(entry, address)

		if err != nil {
			log.Error(nil, "Unable to add the specified entry")
			os.Exit(0)
		}

		log.Info(nil, "Entry added successfully!!!")
	}
}
