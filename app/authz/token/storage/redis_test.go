package storage

import "testing"

func TestGetTokenDataError(t *testing.T) {
	redis, err := NewRedisTokenStorageMock()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	token := "test token"
	_, err = redis.GetTokenData(token)
	if err == nil {
		t.Fatalf("An error should occur")
	}
}

func TestSetTokenData(t *testing.T) {
	redis, err := NewRedisTokenStorageMock()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	token := "test token"
	tokenData := "test data"
	err = redis.SetTokenData(token, tokenData)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func TestGetTokenData(t *testing.T) {
	redis, err := NewRedisTokenStorageMock()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	token := "test token"
	tokenData := "test data"
	err = redis.SetTokenData(token, tokenData)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	d, err := redis.GetTokenData(token)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if tokenData != d {
		t.Fatalf("Token data should match")
	}
}

func TestDeleteTokenData(t *testing.T) {
	redis, err := NewRedisTokenStorageMock()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	token := "test token"
	tokenData := "test data"

	err = redis.SetTokenData(token, tokenData)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	err = redis.DeleteToken(token)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	_, err = redis.GetTokenData(token)
	if err == nil {
		t.Fatalf("An error should occur")
	}
}
