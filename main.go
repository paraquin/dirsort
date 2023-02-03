package main

import (
	"flag"
	"os"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/mapper"
	"github.com/paraquin/dirsort/utils"
)

var mapping []config.Mapping

func init() {
	mappingFile := flag.String("set-mapping", "", "move mapping JSON file to user's config directory")
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

	m := mapper.New(mapping)
	m.Sort(dir)
}
