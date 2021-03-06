package db

import (
	"app/config"
	"app/logs"
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectToRedis(config config.RedisConfig) (*redis.Client, error) {
	host := config.Host
	port := config.Port
	uri := fmt.Sprintf("%s:%d", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr: uri,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}

	logs.Info("Connected to Redis %s:%d", config.Host, config.Port)
	return rdb, nil
}
