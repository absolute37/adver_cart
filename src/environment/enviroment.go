package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Config default env
type Config struct {
	MySqlHost         string `env:"MY_SQL_API_HOST"`
	MySqlUsername     string `env:"MY_SQL_API_USERNAME"`
	MySqlPassword     string `env:"MY_SQL_API_PASSWORD"`
	MySqlPort         string `env:"MY_SQL_API_PORT"`
	MySqlDatabaseName string `env:"MY_SQL_API_DATABASE_NAME"`

	TZ string `env:"TZ"`
}

// Init read file env
func Init() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "develop"
	}

	if err := godotenv.Load(".env" + "." + env); err != nil {
		log.Fatal("error loading .env file")
	}
}

// Load config env
func Load() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}
