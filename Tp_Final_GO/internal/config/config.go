package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Creacion de DbConfig ...
type DbConfig struct {
	Driver string `yaml:"driver"`
}

// Structura de Config ...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

// Load config ...
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
