package mapper

import (
	"fmt"
	"os"
	"path"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/utils"
)

type Mapper struct {
	mapping     []config.Mapping
	currentDir  string
	interactive bool
	verbose     bool
}

func New(mapping []config.Mapping, interactive bool, verbose bool) Mapper {
	return Mapper{
		mapping:     mapping,
		interactive: interactive,
		verbose:     verbose,
	}
}

func (m *Mapper) Sort(dir string) {
	m.currentDir = dir
	files := regularFiles(dir)
	for _, file := range files {
		for _, filetype := range m.mapping {
			for _, extension := range filetype.Extensions {
				if utils.Ext(file) == extension {
					m.handleMove(file, filetype.To)
				}
			}
		}
	}
}

func (m *Mapper) handleMove(file os.DirEntry, to string) {
	answer := "Y"
	if m.interactive {
		fmt.Printf("move %q to %q? [Y/n] ", file.Name(), to)
		fmt.Scanln(&answer)
	}
	if answer == "Y" || answer == "y" || answer == "yes" {
		m.move(file, to)
	}

}

func (m *Mapper) move(file os.DirEntry, to string) {
	oldPath := path.Join(m.currentDir, file.Name())
	newPath := path.Join(m.currentDir, to, file.Name())
	utils.EnsureDirs(newPath)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if m.verbose {
		fmt.Printf("%q moved to %q\n", file.Name(), to)
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
