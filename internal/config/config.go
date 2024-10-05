package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true"` //env-default:"production"
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

// MustRead reads the configuration from a given path or from the environment.
func MustRead(configPath string) *Config {
	// If no config path is passed, try to read from ENV variable
	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	// If no config path is available, terminate with an error
	if configPath == "" {
		log.Fatal("Config path is not set. Provide either -config flag or set CONFIG_PATH env variable.")
	}

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	// Parse the config file
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	return &cfg
}
