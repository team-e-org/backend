package authz

import (
	"app/authz/token"
	"app/authz/token/storage"
	"app/db"
	"app/models"
	"encoding/json"
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

func NewAuthLayerMock(data *db.DataStorage) AuthLayerInterface {
	tokenStorage := storage.NewInMemoryTokenStorage()
	return &AuthLayer{
		tokenStorage,
		data,
	}
}

func (a *AuthLayer) AuthenticateUser(email string, password string) (string, error) {
	user, err := a.dataStorage.Users.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	passwordCheckError := checkUserPassword(password, user.HashedPassword)
	if passwordCheckError != nil {
		return "", ErrInvalidPassword
	}

	bytes, err := json.Marshal(&TokenData{
		UserData: user,
	})
	if err != nil {
		return "", ErrInvalidToken
	}

	token := token.NewToken()
	if err = a.tokenStorage.SetTokenData(token, string(bytes)); err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthLayer) GetTokenData(token string) (*TokenData, error) {
	if len(token) == 0 {
		return nil, ErrInvalidToken
	}

	tokenDataString, err := a.tokenStorage.GetTokenData(token)
	if err == storage.ErrInvalidToken {
		return nil, ErrInvalidToken
	}

	var tokenData TokenData
	if err = json.Unmarshal([]byte(tokenDataString), &tokenData); err != nil {
		return nil, err
	}
	return &tokenData, nil
}

func (al *AuthLayer) TokenStorage() storage.TokenStorage {
	return al.tokenStorage
}

func (al *AuthLayer) DataStorage() *db.DataStorage {
	return al.dataStorage
}
