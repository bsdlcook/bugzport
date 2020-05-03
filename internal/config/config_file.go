package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfigName = "bugz.yml"
	defaultConfigDir  = "bugzport"
)

type Config struct {
	Dir  string `yaml:"dir"`
	Jail string `yaml:"jail"`
	Tree string `yaml:"tree"`
}

func ParseConfig(name string) (Config, error) {
	config, err := parseConfigFile(name)
	if err != nil {
		return Config{}, nil
	}

	return config, nil
}

func dirPath() string {
	dir, _ := homedir.Expand("~/.config/" + defaultConfigDir)
	return dir
}

func filePath() string {
	return path.Join(dirPath(), defaultConfigName)
}

func readConfigFile(name string) ([]byte, error) {
	file, err := os.Open(name)
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

func parseConfigFile(name string) (Config, error) {
	data, err := readConfigFile(name)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return Config{}, err
	}

	return config, err
}
