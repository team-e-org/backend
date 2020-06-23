package db

import (
	"app/config"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func ConnectToRedis(config config.RedisConfig) (redis.Conn, error) {
	host := config.Host
	port := config.Port
	uri := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.Dial("tcp", uri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
