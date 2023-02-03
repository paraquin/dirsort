package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/paraquin/dirsort/utils"
)

const APP_NAME = "dirsort"

type Mapping struct {
	Extensions []string `json:"extensions"`
	To         string   `json:"to"`
}

func New(mappingFile string) {
	mappingFile = utils.AbsolutePath(mappingFile)
	err := utils.CopyFile(mappingFile, configPath())
	if err != nil {
		utils.Error(err.Error())
	}
	os.Exit(0)
}

func GetMapping() ([]Mapping, error) {
	cfgPath := configPath()
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("%v\nMaybe you need to run program with --set-mapping flag", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	mapping := []Mapping{}
	if err = decoder.Decode(&mapping); err != nil {
		return nil, err
	}
	return mapping, nil
}

func configPath() string {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		utils.Error(err.Error())
	}
	return path.Join(cfgPath, APP_NAME, "mapping.json")
}
