package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiKey       string `yaml:"apiKey"`
	Instructions string `yaml:"instructions"`
}

func NewConfig(path string) (*Config, error) {
	return readFile(path)
}

func readFile(path string) (*Config, error) {
	file, err := os.Open(path)

	config := Config{}

	if err == nil {
		if err := yaml.NewDecoder(file).Decode(&config); err != nil {
			return &config, err
		}
	} else {
		println(err.Error())
	}

	if config.ApiKey == "" {
		return &config, errors.New("no apiKey")
	}

	if config.Instructions == "" {
		config.Instructions = defaultInstructions()
	}

	return &config, nil
}

func defaultInstructions() string {
	return "You are a CLI assistant that provides Linux commands. Respond only with the exact command, without any explanation or additional text. If multiple commands are needed, separate them with '&&' or ';' as appropriate."
}
