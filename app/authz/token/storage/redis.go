package storage

import "github.com/go-redis/redis"

type redisTokenStorage struct {
	redis *redis.Client
}

func NewRedisTokenStorage(redis *redis.Client) TokenStorage {
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
