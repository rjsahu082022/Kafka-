package config

import (
	"fmt"
	"os"

	"kafka/golang-kafka/kafka-consumer/internal/log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Uri      string `yaml:"uri"`
	Server   struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Consumer struct {
		Address string   `yaml:"address"`
		Topic   []string `yaml:"topic"`
	} `yaml:"consumer"`
}

func Load(log log.Logger, configFile string) (*Config, error) {
	config := &Config{}
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		log.Errorf("error reading config file: %s", err)
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	if err = yaml.Unmarshal(bytes, &config); err != nil {
		log.Errorf("error unmarshalling config: %s", err)
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}
	return config, nil
}
