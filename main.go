package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/utils"
)

var mapping []config.Mapping

func init() {
	mappingFile := flag.String("set-mapping", "", "move mapping JSON file to user's config directory")
	flag.Parse()

	if *mappingFile != "" {
		log.Println(*mappingFile)
		config.New(*mappingFile)
	}
	var err error
	mapping, err = config.GetMapping()
	if err != nil {
		utils.Error(err.Error())
	}
}

func main() {
	fmt.Println(mapping)

	dir, _ := os.Getwd()
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	fileEntries, err := os.ReadDir(dir)
	if err != nil {
		utils.Error(err.Error())
	}

	fileEntries = excludeDirectories(fileEntries)
	for _, file := range fileEntries {
		println(file.Name())
	}
}

func excludeDirectories(dirEntries []fs.DirEntry) []fs.DirEntry {
	result := make([]fs.DirEntry, 0, len(dirEntries))
	for _, de := range dirEntries {
		if !de.IsDir() {
			result = append(result, de)
		}
	}
	return result
}
