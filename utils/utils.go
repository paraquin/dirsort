package utils

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

// Writes an error message to os.Stderr.
func Error(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

func WrapError(caller string, err error) error {
	return fmt.Errorf("%s: %w", caller, err)
}

// Copy src file content to dst
// If dst contains not existing path, create it
func CopyFile(src, dst string) error {
	const op = "utils.CopyFile"

	buf, err := os.ReadFile(src)
	if err != nil {
		return WrapError(op, err)
	}

	if err = EnsureDirs(dst); err != nil {
		return WrapError(op, err)
	}

	err = os.WriteFile(dst, buf, 0644)
	if err != nil {
		return WrapError(op, err)
	}

	return nil
}

// Move src file content to dst
// If dst contains not existing path, create it
func MoveFile(src, dst string) error {
	const op = "utils.MoveFile"

	srcFile, err := os.Open(src)
	if err != nil {
		return WrapError(op, err)
	}

	buf, err := io.ReadAll(srcFile)
	if err != nil {
		return WrapError(op, err)
	}

	if err = EnsureDirs(dst); err != nil {
		return WrapError(op, err)
	}

	srcStat, err := srcFile.Stat()
	if err != nil {
		return WrapError(op, err)
	}

	newFile, err := os.Create(dst)
	if err != nil {
		return WrapError(op, err)
	}
	defer newFile.Close()

	if _, err = newFile.Write(buf); err != nil {
		return WrapError(op, err)
	}

	newFile.Chmod(srcStat.Mode())
	os.Chtimes(dst, srcStat.ModTime(), srcStat.ModTime())

	srcFile.Close() // Need close before removing
	if err = os.Remove(src); err != nil {
		return WrapError(op, err)
	}

	return nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Creates necessary directories for provided filepath
func EnsureDirs(file string) error {
	dir := filepath.Dir(file)
	err := os.MkdirAll(dir, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

// AbsolutePath returns an absolute representation of path.
// In addition handles '~' symbol as user home directory.
func AbsolutePath(p string) string {
	if p[0] == '~' {
		homeDir, _ := os.UserHomeDir()
		return filepath.Join(homeDir, p[1:])
	}
	abs, err := filepath.Abs(p)
	if err != nil {
		Error(err)
	}
	return abs
}

// Returns the file name extension withot a dot.
// If file has not extension returns empty string
func Ext(file string) string {
	ext := path.Ext(file)
	if len(ext) == len(file) { // hidden file without extension
		return ""
	}
	if len(ext) > 0 {
		return ext[1:]
	}
	return ""
}
