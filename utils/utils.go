package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Writes text to os.Stderr and exit.
func Error(text string) {
	fmt.Fprintln(os.Stderr, text)
	os.Exit(1)
}

// Copy src file content to dst
// If dst contains not existing path, create it
func CopyFile(src, dst string) error {
	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err = EnsureDirs(dst); err != nil {
		return err
	}
	err = os.WriteFile(dst, buf, 0644)
	return err
}

// Creates necessary directories for provided filepath
func EnsureDirs(filepath string) error {
	filepath = path.Dir(filepath)
	err := os.MkdirAll(filepath, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

// AbsolutePath returns a absolute representation of path.
// In addition handles '~' symbol as user home directory.
func AbsolutePath(p string) string {
	if p[0] == '~' {
		homeDir, _ := os.UserHomeDir()
		return path.Join(homeDir, p[1:])
	}
	abs, err := filepath.Abs(p)
	if err != nil {
		Error(err.Error())
	}
	return abs
}

// Returns the file name extension withot a dot.
// If file has not extension returns empty string
func Ext(file string) string {
	ext := path.Ext(file)
	if len(ext) > 0 {
		return ext[1:]
	}
	return ""
}
