package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server Server
	DB     DBConfig
}

type Server struct {
	Port int
}

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

func ReadConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'SERVER_PORT': %w", err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'MYSQL_PORT': %w", err)
	}

	return &Config{
		Server{
			Port: port,
		},
		DBConfig{
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			DBName:   os.Getenv("MYSQL_DATABASE"),
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     dbPort,
		},
	}, nil
}
