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
	var operation uint
	var hostsFile string
	var entry string
	var address string

	flag.UintVar(&operation, "op", uint(file.Find), "Indicate the operation")
	flag.StringVar(&hostsFile, "file", "/etc/hosts", "The user hosts file")
	flag.StringVar(&entry, "entry", "", "Entry to add to the hostsfile")
	flag.StringVar(&address, "address", "127.0.0.1", "Address for the entry")

	flag.Parse()

	reader := file.NewOSReader()

	log := logger.New(logger.TEXT, []io.Writer{os.Stdout})

	if entry == "" {
		log.Fatal(nil, "You must specify a valid entry")
	}

	op := file.Operation(operation)

	flag, err := file.GetFlagForOperation(op)

	if err != nil {
		log.Fatal(nil, "Invalid operation")
	}

	f, err := reader.Read(hostsFile, flag)

	if err != nil {
		log.Error(nil, "Unable to read file")
		os.Exit(0)
	}

	defer f.Close()

	operator := file.NewOperator(f, log)

	switch op {
	case file.Find:
		line, err := operator.Find(file.NewEntry(entry, address))

		if err != nil {
			log.Error(nil, fmt.Sprintf("Entry '%s' not found...", entry))
			os.Exit(0)
		}

		log.Info(nil, fmt.Sprintf("Entry '%s' found on line '%d'", entry, line))
	}

	// if remove {
	// 	err = operator.Remove(entry)

	// 	if err != nil {
	// 		log.Error(nil, "Unable to remove the specified entry")
	// 		os.Exit(0)
	// 	}

	// 	log.Info(nil, "Entry removed successfully!!!")
	// } else {
	// 	err = operator.Assert(entry, address)

	// 	if err != nil {
	// 		log.Error(nil, "Unable to add the specified entry")
	// 		os.Exit(0)
	// 	}

	// 	log.Info(nil, "Entry added successfully!!!")
	// }
}
