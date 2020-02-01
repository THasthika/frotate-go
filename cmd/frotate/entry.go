package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tharindu96/frotate-go"
)

type cliArgs struct {
	filepath  string
	directory string
	prefix    string
	limit     uint
}

func main() {
	cliargs := getCliArgs()
	if err := frotate.AddFile(cliargs.filepath, cliargs.prefix, cliargs.directory, cliargs.limit); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getCliArgs() *cliArgs {

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	filepath := ""

	// addCommand
	addPrefixPtr := addCommand.String("prefix", "backup", "Prefix of the rotation file")
	addDirectoryPtr := addCommand.String("directory", ".", "Directory to store the rotation files")
	addLimitPtr := addCommand.Uint("limit", 10, "The maximum number of files allowed")

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s [command]\n", os.Args[0])
		fmt.Printf("    command - add [filename] [options]\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Printf("usage: %s add [filename]\n", os.Args[0])
			os.Exit(1)
		}
		filepath = os.Args[2]
		addCommand.Parse(os.Args[3:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addCommand.Parsed() {
		return &cliArgs{
			filepath:  filepath,
			prefix:    *addPrefixPtr,
			directory: *addDirectoryPtr,
			limit:     *addLimitPtr,
		}
	}

	return nil
}
