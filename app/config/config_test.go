package config

import (
	"os"
	"reflect"
	"testing"
)

func TestNoServerPort(t *testing.T) {
	_, err := ReadConfig()
	if err == nil {
		t.Fatalf("'SERVER_PORT' is not set, but no error")
	}
}

func TestNoMySQLPort(t *testing.T) {
	os.Setenv("SERVER_PORT", "3000")
	_, err := ReadConfig()
	if err == nil {
		t.Fatalf("'MYSQL_PORT' is not set, but no error")
	}
}

func TestConfig(t *testing.T) {
	os.Setenv("MYSQL_DATABASE", "test_db")
	os.Setenv("MYSQL_USER", "test_user")
	os.Setenv("MYSQL_PASSWORD", "password")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_HOST", "db")
	os.Setenv("SERVER_PORT", "3000")

	config, err := ReadConfig()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	if config.Server.Port != 3000 {
		t.Fatalf("Server port is not 3000, got %v", config.Server.Port)
	}

	dbConfig := DBConfig{
		User:     "test_user",
		Password: "password",
		DBName:   "test_db",
		Host:     "db",
		Port:     3306,
	}

	if !reflect.DeepEqual(config.DB, dbConfig) {
		t.Fatalf("Configs do not match, got: %v, actual: %v", config.DB, dbConfig)
	}

	testConfig := Config{
		Server: Server{
			Port: 3000,
		},
		DB: dbConfig,
	}

	if !reflect.DeepEqual(*config, testConfig) {
		t.Fatalf("Configs do not match, got: %v, actual: %v", config, testConfig)
	}
}
