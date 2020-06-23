package authz

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	authToken = "X-Auth-Token"
)

func GetUserIDByToken(al AuthLayerInterface, token string) (int,
	error) {

	tokenData, err := al.GetTokenData(token)
	if err != nil {
		return 0, err
	}

	return tokenData.UserData.ID, nil
}

func checkUserPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
