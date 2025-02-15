package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tracker struct {
		TrackingHosts []string `yaml:"traking_hosts"`
	} `yaml:"tracker"`
}

var AppConfig Config

func init() {

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(basePath)

	data, err := os.ReadFile(filepath.Join(basePath, "config.yml"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		panic(err)
	}
}
