package storage

import "github.com/go-redis/redis"

type redisTokenStorage struct {
	redis *redis.Client
}

func NewRedisTokenStorage(redis *redis.Client) TokenStorage {
	return &redisTokenStorage{redis: redis}
}

func (s *redisTokenStorage) GetTokenData(token string) (string, error) {
	data, err := s.redis.Get(token).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

func (s *redisTokenStorage) SetTokenData(token string, tokenData string) error {
	err := s.redis.Set(token, tokenData, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *redisTokenStorage) DeleteToken(token string) error {
	err := s.redis.Del(token).Err()
	if err != nil {
		return err
	}
	return nil
}
