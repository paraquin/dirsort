package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/mapper"
	"github.com/paraquin/dirsort/utils"

	flag "github.com/spf13/pflag"
)

var isInteractive bool
var isVerbose bool
var noWaiting bool

func init() {
	mappingFile := flag.StringP("set-mapping", "s", "", "move mapping YAML file to user's config directory")
	flag.BoolVarP(&isInteractive, "interactive", "i", false, "prompt before every move")
	flag.BoolVarP(&isVerbose, "verbose", "v", false, "explain what is being done")
	flag.BoolVarP(&noWaiting, "no-waiting", "W", false, "does not require to press a button to exit")
	help := flag.BoolP("help", "h", false, "display this help and exit")
	flag.Parse()

	if *help {
		printHelp()
	}

	if *mappingFile != "" {
		config.New(*mappingFile)
	}
}

// Prints help message
func printHelp() {
	executableName := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %s [OPTIONS]... DIR\n", executableName)
	fmt.Println(flag.CommandLine.FlagUsages())
}

func main() {
	if !noWaiting {
		defer wait()
	}

	if len(flag.Args()) != 1 {
		printHelp()
		return
	}

	path := flag.Arg(0)
	if utils.Ext(path) == "yaml" || utils.Ext(path) == "yml" {
		err := config.New(path)
		if err != nil {
			utils.Error(err)
			return
		} else {
			fmt.Println("Configuration is successfully updated.")
		}
		return
	}

	mapping, err := config.GetMapping()
	if err != nil {
		utils.Error(err)
		return
	}

	m := mapper.New(mapping, isInteractive, isVerbose)
	m.Sort(path)
}

func wait() {
	fmt.Print("\nPress ENTER to exit...")
	fmt.Scanln()
}
