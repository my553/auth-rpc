package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"3600"`
	GRPC        GRPC          `yaml:"grpc"`
}

type GRPC struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"5s"`
}

func MustLoad() *Config {
	cfgPath := fetchConfigPath()

	if cfgPath == "" {
		panic("config cfgPath is required")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic("config file does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic(fmt.Sprintf("Error while reading config by path %s: %s", cfgPath, err.Error()))
	}

	return &cfg
}

func fetchConfigPath() string {
	var cfgPath string

	flag.StringVar(&cfgPath, "config", "", "path to config file")
	flag.Parse()

	if cfgPath == "" {
		cfgPath = os.Getenv("CONFIG_PATH")
	}

	return cfgPath
}
