package config

import (
	"fmt"
	"os"

	"kafka/golang-kafka/kafka-producer/internal/log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Producer struct {
		Address string   `yaml:"address"`
		Topic   []string `yaml:"topic"`
	} `yaml:"producer"`
}

func Load(log log.Logger, configFile string) (*Config, error) {
	config := &Config{}
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		log.Error("error reading config file: %s", err)
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	if err = yaml.Unmarshal(bytes, &config); err != nil {
		log.Errorf("error unmarshalling config: %s", err)
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return config, nil
}
