package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/paraquin/dirsort/utils"
)

const AppName = "dirsort"
const configFilename = "mapping.json"

type Mapping map[string][]string

type mappingItem struct {
	Extensions []string `json:"extensions"`
	To         string   `json:"to"`
}

var configPath string

func init() {
	var err error
	configPath, err = getConfigPath()
	if err != nil {
		utils.Error(err)
	}
}

// reada
func New(mappingFilePath string) {
	mappingFilePath = utils.AbsolutePath(mappingFilePath)
	err := utils.CopyFile(mappingFilePath, configPath)
	if err != nil {
		utils.Error(err)
	}
	os.Exit(0)
}

func GetMapping() (Mapping, error) {
	data, err := readConfigFile()
	if err != nil {
		return nil, err
	}
	mappingItems := []mappingItem{}
	err = json.Unmarshal(data, &mappingItems)
	if err != nil {
		return nil, err
	}
	return mappingItemsToMapping(mappingItems), nil
}

func mappingItemsToMapping(items []mappingItem) map[string][]string {
	result := make(map[string][]string)
	for _, item := range items {
		result[item.To] = item.Extensions
	}
	return result
}

func readConfigFile() ([]byte, error) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("%v\nMaybe you need to run program with --set-mapping flag", err)
	}
	return configData, nil
}

// getConfigPath returns path to user config dir
func getConfigPath() (string, error) {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(cfgPath, AppName, configFilename), nil
}
