package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
)

type Config struct {
	HttpPort int `mapstructure:"http_port"`
}

var cfg Config

func LoadConfig() {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	viper.AddConfigPath(currentDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}

	log.Printf("Configuration loaded successfully:%+v\n", cfg)
}

func GetConfig() *Config {
	return &cfg
}
