package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port       int    `envconfig:"PORT" default:"8888"`
	DBDriver   string `envconfig:"DB_DRIVER" default:"postgres"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBName     string `envconfig:"DB_NAME" default:"mydb"`
	DBUser     string `envconfig:"DB_USER" default:"myuser"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"mypassword"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
}

var DBConn *sqlx.DB

func (c Config) getDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

var config Config

func LoadConfig() error {
	err := envconfig.Process("", &config)
	if err != nil {
		return fmt.Errorf("could not envconfig.Process: %w", err)
	}
	DBConn, err = sqlx.Connect(config.DBDriver, config.getDSN())
	if err != nil {
		return fmt.Errorf("could no connect to db: %w", err)
	}
	if err := DBConn.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}
	log.Println("config loaded successfully")
	return nil
}
