package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfigName = "config.yml"
	defaultConfigDir  = "bugzport"
)

type Config struct {
	Dir  string `yaml:"poud_port_dir"`
	Jail string `yaml:"poud_jail"`
	Tree string `yaml:"poud_tree"`
}

func ParseConfig() (Config, error) {
	config, err := parseConfigFile()
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func dirPaths() []string {
	homeDir, _ := homedir.Expand("~/.config/" + defaultConfigDir)
	globalDir := "/usr/local/etc/" + defaultConfigDir
	return []string{homeDir, globalDir}
}

func configPath() (string, error) {
	var paths []string
	for _, configPath := range dirPaths() {
		paths = append(paths, path.Join(configPath, defaultConfigName))
	}

	var config string
	for _, configFile := range paths {
		_, err := os.Stat(configFile)
		if !os.IsNotExist(err) {
			config = configFile
			break
		}
	}

	if config == "" {
		return "", fmt.Errorf("could not find bugzport configuration '%s' file in path(s) %s", defaultConfigName, dirPaths())
	}

	return config, nil
}

func readConfigFile() ([]byte, error) {
	configFile, err := configPath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseConfigFile() (Config, error) {
	data, err := readConfigFile()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
