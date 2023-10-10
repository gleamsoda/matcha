package config

import (
	"sync"

	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Environment   string `env:"ENVIRONMENT"`
	ServerAddress string `env:"SERVER_ADDRESS"`
}

// Get reads configuration from file or environment variables.
var Get = sync.OnceValue(func() (cfg Config) {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return
})

func (c Config) IsDevelopment() bool {
	return c.Environment == "development"
}
