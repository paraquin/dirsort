package mapper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/paraquin/dirsort/config"
	"github.com/paraquin/dirsort/utils"
)

type Mapper struct {
	mapping     config.Mapping
	currentDir  string
	interactive bool
	verbose     bool
}

func New(mapping config.Mapping, interactive bool, verbose bool) Mapper {
	return Mapper{
		mapping:     mapping,
		interactive: interactive,
		verbose:     verbose,
	}
}

func (m *Mapper) Sort(dir string) {
	m.currentDir = utils.AbsolutePath(dir)
	err := os.Chdir(m.currentDir)
	if err != nil {
		utils.Error(err)
	}
	files := getRegularFiles(m.currentDir)
	for _, file := range files {
		for dst, extensions := range m.mapping {
			for _, extension := range extensions {
				if utils.Ext(file.Name()) == extension {
					m.handleMove(file, dst)
				}
			}
		}
	}
}

func (m *Mapper) handleMove(file os.DirEntry, dst string) {
	if !m.promptUser(file, dst) {
		return
	}
	err := m.move(file, dst)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	m.informUser(file.Name(), dst)
}

func (m *Mapper) promptUser(file os.DirEntry, to string) bool {
	if !m.interactive {
		return true
	}

	answer := "Y" // Y is a default selection

	fmt.Printf("move %q to %q? [Y/n] ", file.Name(), utils.AbsolutePath(to))
	fmt.Scanln(&answer)

	if answer != "Y" && answer != "y" && answer != "yes" {
		return false
	}

	return true
}

func (m *Mapper) informUser(filename, movedTo string) {
	if m.verbose {
		fmt.Printf("%q moved to %q\n", filename, movedTo)
	}
}

func (m *Mapper) move(file os.DirEntry, dst string) (err error) {
	dstAbsolute := utils.AbsolutePath(dst)
	oldPath := filepath.Join(m.currentDir, file.Name())
	newPath := filepath.Join(dstAbsolute, file.Name())
	err = utils.EnsureDirs(newPath)
	if err != nil {
		return
	}
	err = os.Rename(oldPath, newPath)
	return
}

func getRegularFiles(dir string) []os.DirEntry {
	dirEntries, err := os.ReadDir(utils.AbsolutePath(dir))
	if err != nil {
		log.Println(err)
	}
	files := make([]os.DirEntry, 0, len(dirEntries))
	for _, entry := range dirEntries {
		if entry.Type().IsRegular() {
			files = append(files, entry)
		}
	}
	return files
}
