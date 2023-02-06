package main

import (
	"os"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/mapper"
	"github.com/paraquin/dirsort/utils"

	flag "github.com/spf13/pflag"
)

var mapping []config.Mapping
var isInteractive bool
var isVerbose bool

func init() {
	mappingFile := flag.StringP("set-mapping", "s", "", "move mapping JSON file to user's config directory")
	flag.BoolVarP(&isInteractive, "interactive", "i", false, "prompt before every move")
	flag.BoolVarP(&isVerbose, "verbose", "v", false, "explain what is being done")
	flag.Parse()

	if *mappingFile != "" {
		config.New(*mappingFile)
	}
	var err error
	mapping, err = config.GetMapping()
	if err != nil {
		utils.Error(err.Error())
	}
}

func main() {
	dir, _ := os.Getwd()
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	m := mapper.New(mapping, isInteractive, isVerbose)
	m.Sort(dir)
}
