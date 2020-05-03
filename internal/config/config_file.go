package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

const (
	configName = "bugz.yml"
	configDir  = "bugzport"
)

type Config struct {
	Dir  string `yaml:"dir"`
	Jail string `yaml:"jail"`
	Tree string `yaml:"tree"`
}

func Dir() string {
	dir, _ := homedir.Expand("~/.config/" + configDir)
	return dir
}

func File() string {
	return path.Join(Dir(), configName)
}

var ReadConfigFile = func(name string) ([]byte, error) {
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

func ParseConfigFile(name string) (Config, error) {
	data, err := ReadConfigFile(name)
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

func ParseConfig(name string) (Config, error) {
	config, err := ParseConfigFile(name)
	if err != nil {
		return Config{}, nil
	}

	return config, nil
}
