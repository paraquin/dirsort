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
		return
	}
	files := getRegularFiles(m.currentDir)
	for _, file := range files {
		for dst, extensions := range m.mapping {
			for _, extension := range extensions {
				if utils.Ext(file.Name()) == extension {
					m.move(file, dst)
				}
			}
		}
	}
}

func (m *Mapper) askUser(questionFormat string, a ...any) bool {
	question := fmt.Sprintf(questionFormat, a...)
	fmt.Printf("%s [Y/n] ", question)

	answer := "Y" // Y is a default selection
	fmt.Scanln(&answer)

	if answer != "Y" && answer != "y" && answer != "yes" {
		return false
	}

	return true
}

func (m *Mapper) informUser(filename, movedTo string) {
	if m.verbose {
		fmt.Printf("%s moved to %s\n", filename, movedTo)
	}
}

func (m *Mapper) move(file os.DirEntry, dst string) {
	dstAbsolute := utils.AbsolutePath(dst)
	oldPath := filepath.Join(m.currentDir, file.Name())
	newPath := filepath.Join(dstAbsolute, file.Name())

	if m.interactive && !m.askUser("Move %s to %s?", file.Name(), dstAbsolute) {
		return
	}

	if utils.FileExists(newPath) {
		if !m.interactive {
			return
		}
		if !m.askUser("%s already exists. Rewrite?", newPath) {
			return
		}
	}

	err := utils.EnsureDirs(newPath)
	if err != nil {
		utils.Error(err)
		return
	}
	err = os.Rename(oldPath, newPath)
	if err != nil {
		// Move file because os.Rename can't work with different drives.
		err = utils.MoveFile(oldPath, newPath)
		if err != nil {
			utils.Error(err)
		}
	}
	m.informUser(file.Name(), dstAbsolute)
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
