package config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Environment   string `env:"ENVIRONMENT"`
	ServerAddress string `env:"SERVER_ADDRESS"`
	DBHost        string `env:"DB_HOST"`
	DBPort        int    `env:"DB_PORT"`
	DBUser        string `env:"DB_USER"`
	DBPassword    string `env:"DB_PASSWORD"`
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

func (c Config) DBConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d?sslmode=disable", c.DBUser, c.DBPassword, c.DBHost, c.DBPort)
}
