package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Configuration struct {
	Port              string `env:"PORT" envDefault:"3002"`
	JwtSecret         string `env:"SECRET_JWT_KEY,required"`
	MySQLUsername     string `env:"MYSQL_USERNAME" envDefault:"root"`
	MySQLPassword     string `env:"MYSQL_PASSWORD" envDefault:"password"`
	MySQLHost         string `env:"MYSQL_HOST" envDefault:"localhost"`
	MySQLPort         string `env:"MYSQL_PORT" envDefault:"3306"`
	MySQLDatabaseName string `env:"MYSQL_DATABASE_NAME" envDefault:"instagram"`
}

func NewConfig() *Configuration {
	cfg := Configuration{}

	// Parse env to configuration
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
