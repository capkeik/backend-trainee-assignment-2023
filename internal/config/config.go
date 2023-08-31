package config

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

type Config struct {
	HTTPAddr         string `envconfig:"HTTP_ADDR"`
	PgURL            string `envconfig:"PG_URL"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		_, err = json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Configuration loaded successfully")
	})

	return &config
}
