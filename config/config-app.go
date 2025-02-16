package config

import (
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tracker struct {
		TrackingHosts []string `yaml:"traking_hosts"`
	} `yaml:"tracker"`
	Notifier struct {
		NotifierType string `yaml:"notifier_type"`
		Slack        struct {
			Token      string   `yaml:"token"`
			ChannelIds []string `yaml:"channel_ids"`
		} `yaml:"slack"`
		Lark struct {
			WebhookUrl string `yaml:"webhook_url"`
		} `yaml:"lark"`
	} `yaml:"notifier"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		Timeout  int    `yaml:"timeout"`
	} `yaml:"redis"`
}

var AppConfig Config

func init() {

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filepath.Join(basePath, "config.yml"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		panic(err)
	}

	AppConfig.Redis.Host = getEnv("REDIS_HOST", AppConfig.Redis.Host)
	AppConfig.Redis.Port = getEnvAsInt("REDIS_PORT", AppConfig.Redis.Port)
	AppConfig.Redis.Password = getEnv("REDIS_PASSWORD", AppConfig.Redis.Password)
	AppConfig.Redis.DB = getEnvAsInt("REDIS_DB", AppConfig.Redis.DB)
	AppConfig.Redis.Timeout = getEnvAsInt("REDIS_TIMEOUT", AppConfig.Redis.Timeout)
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		parsedValue, err := strconv.Atoi(value)
		if err == nil {
			return parsedValue
		}
	}
	return defaultValue
}
