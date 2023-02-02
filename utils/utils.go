package utils

import (
	"fmt"
	"log"
	"os"
	"path"
)

// Write text to os.Stderr and exit.
func Error(text string) {
	fmt.Fprintln(os.Stderr, text)
	os.Exit(1)
}

// Copy src file content to dst
// If dst contains not existing path, create it
func CopyFile(src, dst string) error {
	log.Printf("copy %s to %s...", src, dst)
	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err = enshureDirs(dst); err != nil {
		log.Println("error in enshureDir")
		return err
	}
	err = os.WriteFile(dst, buf, 0644)
	return err
}

// Creates necessary directories for provided filepath
func enshureDirs(filepath string) error {
	filepath = path.Dir(filepath)
	err := os.MkdirAll(filepath, os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}
