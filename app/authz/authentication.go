package authz

import (
	"app/authz/token/storage"
	"app/db"
	"app/models"
	"errors"

	"github.com/gomodule/redigo/redis"
)

var ErrInvalidPassword = errors.New("password is not collect")
var ErrInvalidToken = errors.New("token is invalid")

type TokenData struct {
	UserData *models.User
}

type AuthLayerInterface interface {
	AuthenticateUser(string, string) (string, error)
	GetTokenData(string) (*TokenData, error)
	TokenStorage() storage.TokenStorage
	DataStorage() *db.DataStorage
}

type AuthLayer struct {
	tokenStorage storage.TokenStorage
	dataStorage  *db.DataStorage
}

func NewAuthLayer(data *db.DataStorage, redis redis.Conn) AuthLayerInterface {
	tokenStorage := storage.NewRedisTokenStorage(redis)
	return &AuthLayer{
		tokenStorage,
		data,
	}
}

func (al *AuthLayer) AuthenticateUser(email string, password string) (string, error) {
	return "", nil
}

func (al *AuthLayer) GetTokenData(token string) (*TokenData, error) {
	return nil, nil
}

func (al *AuthLayer) TokenStorage() storage.TokenStorage {
	return al.tokenStorage
}

func (al *AuthLayer) DataStorage() *db.DataStorage {
	return al.dataStorage
}
