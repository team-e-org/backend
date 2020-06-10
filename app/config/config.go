package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server Server
}

type Server struct {
	Port int
}

func ReadConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'SERVER_PORT': %w", err)
	}

	return &Config{
		Server{
			Port: port,
		},
	}, nil
}
