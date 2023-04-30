package main

import (
	"fmt"
	"os"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/mapper"
	"github.com/paraquin/dirsort/utils"

	flag "github.com/spf13/pflag"
)

var mapping config.Mapping
var isInteractive bool
var isVerbose bool

func init() {
	mappingFile := flag.StringP("set-mapping", "s", "", "move mapping YAML file to user's config directory")
	flag.BoolVarP(&isInteractive, "interactive", "i", false, "prompt before every move")
	flag.BoolVarP(&isVerbose, "verbose", "v", false, "explain what is being done")
	help := flag.BoolP("help", "h", false, "display this help and exit")
	flag.Parse()

	if *help {
		printHelp()
	}

	if *mappingFile != "" {
		config.New(*mappingFile)
	}
	var err error
	mapping, err = config.GetMapping()
	if err != nil {
		utils.Error(err)
	}
}

// Prints help message end exit
func printHelp() {
	fmt.Printf("Usage: %s [OPTIONS]... DIR\n", os.Args[0])
	fmt.Println(flag.CommandLine.FlagUsages())
	os.Exit(0)
}

func main() {
	if len(flag.Args()) != 1 {
		printHelp()
	}

	path := flag.Arg(0)
	if utils.Ext(path) == "yaml" || utils.Ext(path) == "yml" {
		config.New(path)
	}

	m := mapper.New(mapping, isInteractive, isVerbose)
	m.Sort(path)
}
