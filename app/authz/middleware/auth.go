package middleware

import (
	"app/authz"
	"app/logs"
	"app/server/handlers"
	"fmt"
	"net/http"
)

const authToken = "X-Auth-Token"

func RequireAuthorization(al authz.AuthLayerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(authToken)

			data, err := al.GetTokenData(token)
			if err == authz.ErrInvalidToken {
				logs.Error(fmt.Sprintf("error while looking up session token: %s, err: %v", token, err))
				handlers.Unauthorized(w, r)
				return
			}
			if err != nil {
				logs.Error(fmt.Sprintf("error while looking up session token: %s, err: %v", token, err))
				handlers.InternalServerError(w, r)
				return
			}

			if data == nil {
				logs.Info("Token data not found %s", token)
				handlers.Unauthorized(w, r)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
