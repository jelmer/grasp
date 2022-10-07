package config

import (
	"math/rand"
	"net/url"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/jelmer/grasp/pkg/datastore/sqlstore"
)

// Config wraps the configuration structs for the various application parts
type Config struct {
	Database *sqlstore.Config
	Secret   string
}

// Parse environment into a Config struct
func Parse() *Config {
	var cfg Config

	// with config file loaded into env values, we can now parse env into our config struct
	err := envconfig.Process("Grasp", &cfg)
	if err != nil {
		log.Fatalf("Error parsing configuration from environment: %s", err)
	}

	if cfg.Database.URL != "" {
		u, err := url.Parse(cfg.Database.URL)
		if err != nil {
			log.Fatalf("Error parsing DATABASE_URL from environment: %s", err)
		}
		if u.Scheme == "postgres" {
			cfg.Database.Driver = "postgres"
		}
	}

	// alias sqlite to sqlite3
	if cfg.Database.Driver == "sqlite" {
		cfg.Database.Driver = "sqlite3"
	}

	// use absolute path to sqlite3 database
	if cfg.Database.Driver == "sqlite3" {
		cfg.Database.Name, _ = filepath.Abs(cfg.Database.Name)
	}

	// if secret key is empty, use a randomly generated one
	if cfg.Secret == "" {
		cfg.Secret = randomString(40)
	}

	return &cfg
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}

	return string(bytes)
}
