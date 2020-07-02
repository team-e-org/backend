package db

import (
	"app/config"
	"testing"
)

func TestConnectToRedis(t *testing.T) {
	c := config.RedisConfig{
		Host: "test",
		Port: 0,
	}
	_, err := ConnectToRedis(c)
	if err == nil {
		t.Error("Error should occur")
	}
}
