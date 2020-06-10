package storage

import "sync"

type inMemoryStorage struct {
	storage sync.Map
}

func NewInMemoryTokenStorage() *inMemoryStorage {
	return &inMemoryStorage{}
}

func (s *inMemoryStorage) GetTokenData(token string) (string, error) {
	data, exists := s.storage.Load(token)
	if !exists {
		return "", ErrInvalidToken
	}
	return data.(string), nil
}

func (s *inMemoryStorage) SetTokenData(token string, tokenData string) error {
	s.storage.Store(token, tokenData)
	return nil
}

func (s *inMemoryStorage) DeleteToken(token string) error {
	s.storage.Delete(token)
	return nil
}
