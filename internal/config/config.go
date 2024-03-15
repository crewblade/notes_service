package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	ConnectionString string     `yaml:"connection_string"`
	GRPC             GRPCConfig `yaml:"grpc"`
}
type GRPCConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := os.Getenv("CONFIG_PATH")
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config with path " + path + "\n" + err.Error())
	}
}
