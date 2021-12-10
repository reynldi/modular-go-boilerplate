package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type application struct {
	AppName string `yaml:"app_name"`
	Environment string `yaml:"env"`
	Port int `yaml:"port"`
}

type AppConfig struct {
	Postgres    postgres    `yaml:"postgres"`
	Application application `yaml:"application"`
}

var globalConfig = AppConfig{}

func New(configPath string) (*AppConfig, error) {

	config := &globalConfig

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// init yaml file
	d := yaml.NewDecoder(file)

	d.Decode(&config)

	return config, nil
}

func GetConfig() *AppConfig {
	return &globalConfig
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", path)
	}

	return nil
}
