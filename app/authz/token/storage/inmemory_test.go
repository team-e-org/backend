package storage

import (
	"app/authz/token"
	"sync"
	"testing"
)

func TestSetTokenDataConccurency(t *testing.T) {
	var wg sync.WaitGroup
	storage := NewInMemoryTokenStorage()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			storage.SetTokenData(token.NewToken(), "tokenData")
		}(&wg)
	}

	wg.Wait()
}
