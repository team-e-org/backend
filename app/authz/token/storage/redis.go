package storage

import (
	"github.com/gomodule/redigo/redis"
)

type redisTokenStorage struct {
	redis redis.Conn
}

func NewRedisTokenStorage(redis redis.Conn) TokenStorage {
	return &redisTokenStorage{redis: redis}
}

func (s *redisTokenStorage) GetTokenData(token string) (string, error) {
	return "", nil
}

func (s *redisTokenStorage) SetTokenData(token string, tokenData string) error {
	return nil
}

func (s *redisTokenStorage) DeleteToken(token string) error {
	return nil
}
