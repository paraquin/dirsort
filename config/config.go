package config

import (
	"fmt"
	"os"
	"path"

	"github.com/paraquin/dirsort/utils"
	"gopkg.in/yaml.v3"
)

const AppName = "dirsort"
const configFilename = "config.yaml"

type Mapping map[string][]string

var configPath string

func init() {
	var err error
	configPath, err = getConfigPath()
	if err != nil {
		utils.Error(err)
	}
}

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
	m := struct {
		Mapping `yaml:"mapping"`
	}{}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m.Mapping, nil
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
