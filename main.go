package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	fileEntries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
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
