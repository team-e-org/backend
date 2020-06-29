// +build integration

package db

import (
	"app/config"
	"testing"
)

func TestRedis(t *testing.T) {
	redisConfig, err := config.ReadRedisConfig()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	rdb, err := ConnectToRedis(*redisConfig)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	err = rdb.Set("test", "test data", 0).Err()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	s, err := rdb.Get("test").Result()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if s != "test data" {
		t.Fatalf("Data do not match error")
	}
}
