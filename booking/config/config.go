package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServiceName string `env:"NAME"`

	Port string `env:"PORT" env-default:"8080"`

	Brokers []string `env:"BROKERS"`
}

var (
	cfg  *Config
	once sync.Once
)

func ParseConfig() *Config {
	once.Do(func() {
		if err := cleanenv.ReadConfig("/.env", cfg); err != nil {
			panic(err)
		}
	})

	return cfg
}
