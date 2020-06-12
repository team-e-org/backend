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

func TestInMemoryStorage_GetTokenData(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		searchKey string
		want      string
		wantError bool
	}{
		{"key", "abc", "k", "", true},
		{"key", "abc", "key", "abc", false},
	}

	for _, tt := range tests {
		storage := NewInMemoryTokenStorage()
		_ = storage.SetTokenData(tt.key, tt.value)
		got, err := storage.GetTokenData(tt.searchKey)

		if !tt.wantError && err != nil {
			t.Fatalf("want no err, but has error %v", err)
		}

		if tt.wantError && err == nil {
			t.Fatalf("want err, but has no err")
		}

		if tt.want != got {
			t.Fatalf("token data mismatch. want: %s, got: %s", tt.want, got)
		}
	}
}
