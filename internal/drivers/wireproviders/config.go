package wireproviders

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger LoggerConfig `yaml:"logger"`
	Server ServerConfig `yaml:"server"`
	TLS    TLSConfig    `yaml:"tls"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type ServerConfig struct {
	PublisherPort  uint16 `yaml:"publisherPort"`
	SubscriberPort uint16 `yaml:"subscriberPort"`
}

type TLSConfig struct {
	X509 string `yaml:"x509"`
	Key  string `yaml:"key"`
}

func NewConfig(options Options) (Config, error) {
	yamlBytes, err := os.ReadFile(options.File)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read configuration file: %s", err)
	}

	var conf Config
	if err := yaml.NewDecoder(bytes.NewReader(yamlBytes)).Decode(&conf); err != nil {
		return Config{}, fmt.Errorf("failed to decode yaml configuration file: %s", err)
	}

	return conf, nil
}
