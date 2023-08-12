package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string	`yaml:"host"`
		Port string	`yaml:"port"`
	}	`yaml:"server"`
	AppInfo struct {
		Name string `yaml:"name"`
		Version string `yaml:"version"`
	}	`yaml:"app-info"`
}

var cfg Config

func LoadConfig(env string) Config {
	f, err := os.Open("./config/" + env + ".yaml")
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err.Error()))
	}
	defer f.Close()
	
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("Error loading config: %v", err.Error())
	}

	return cfg
}

func GetConfig() Config {
	return cfg
}