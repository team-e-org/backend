package authz

import (
	"app/authz/token"
	"app/authz/token/storage"
	"app/db"
	"encoding/json"
)

type AuthLayerMock struct {
	tokenStorage storage.TokenStorage
	dataStorage  *db.DataStorage
}

func (al *AuthLayerMock) TokenStorage() storage.TokenStorage {
	return al.tokenStorage
}

func (al *AuthLayerMock) DataStorage() *db.DataStorage {
	return al.dataStorage
}

func NewAuthLayerMock(data *db.DataStorage) AuthLayerInterface {
	tokenStorage := storage.NewInMemoryTokenStorage()
	return &AuthLayerMock{
		tokenStorage,
		data,
	}
}

func (a *AuthLayerMock) AuthenticateUser(email string, password string) (string, error) {
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

func (a *AuthLayerMock) GetTokenData(token string) (*TokenData, error) {
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
