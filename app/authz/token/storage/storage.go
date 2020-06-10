package storage

import "errors"

type TokenStorage interface {
	GetTokenData(string) (string, error)
	SetTokenData(string, string) error
	DeleteToken(string) error
}

var ErrInvalidToken = errors.New("Invalid Token")
var ErrTokenDataUpdateFailed = errors.New("Unable to update token data")
var ErrTokenInvalidationFailed = errors.New("Unable to delete token")
