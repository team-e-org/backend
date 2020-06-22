package authz

import (
	"net/http"
)

const (
	authToken = "X-Auth-Token"
)

func GetUserIDIfAvailable(r *http.Request, al AuthLayerInterface) (int, error) {
	token := r.Header.Get(authToken)
	if len(token) == 0 {
		return 0, nil
	}

	return GetUserIdByToken(al, token)
}

func GetUserIdByToken(al AuthLayerInterface, token string) (int,
	error) {

	tokenData, err := al.GetTokenData(token)
	if err != nil {
		return 0, err
	}

	return tokenData.UserData.ID, nil
}
