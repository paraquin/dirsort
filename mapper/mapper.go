package mapper

import (
	"fmt"
	"os"
	"path"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/utils"
)

type Mapper struct {
	mapping    []config.Mapping
	currentDir string
}

func New(mapping []config.Mapping) Mapper {
	return Mapper{
		mapping: mapping,
	}
}

func (m *Mapper) Sort(dir string) {
	m.currentDir = dir
	files := regularFiles(dir)
	for _, file := range files {
		for _, filetype := range m.mapping {
			for _, extension := range filetype.Extensions {
				fileExtension := path.Ext(file.Name())[1:]
				if fileExtension == extension {
					m.Move(file, filetype.To)
				}
			}
		}
	}
}

func (m *Mapper) Move(file os.DirEntry, to string) {
	oldPath := path.Join(m.currentDir, file.Name())
	newPath := path.Join(m.currentDir, to, file.Name())
	utils.EnsureDirs(newPath)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func regularFiles(dir string) []os.DirEntry {
	dirEntries, _ := os.ReadDir(dir)
	files := make([]os.DirEntry, 0, len(dirEntries))
	for _, entry := range dirEntries {
		if entry.Type().IsRegular() {
			files = append(files, entry)
		}
	}
	return files
}
