package db

import (
	"app/config"
	"testing"
)

func TestConnectToMySQL(t *testing.T) {
	c := config.DBConfig{
		Host:     "test",
		Port:     0,
		User:     "test",
		Password: "test",
		DBName:   "test",
		TimeZone: "test",
	}
	_, err := ConnectToMySql(c)
	if err == nil {
		t.Error("Error should occur")
	}
}
