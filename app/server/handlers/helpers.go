package handlers

import (
	"fmt"
	"net/http"

	"app/authz"
	"app/helpers"
	"app/logs"
)

const (
	contentType           = "Content-Type"
	authToken             = "X-Auth-Token"
	jsonContent           = "application/json"
	mp3Content            = "audio/mpeg"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	UNAUTHORIZED          = "UNAUTHORIZED"
	NOT_FOUND             = "NOT_FOUND"
)

func logRequest(r *http.Request) {
	logs.Info("Received request: %s", requestSummary(r))
}

func requestSummary(r *http.Request) string {
	return fmt.Sprintf("%v %v", r.Method, r.URL)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400 bad request", http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "401 unauthorized", http.StatusUnauthorized)
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "403 forbidden", http.StatusForbidden)
}

func Conflict(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "409 conflict", http.StatusConflict)
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}

func ResponseError(w http.ResponseWriter, r *http.Request, err error) {
	switch err.(type) {
	case *helpers.InternalServerError:
		logs.Error("Request: %s, while getting user's boards: %v", requestSummary(r), err)
		InternalServerError(w, r)
	case *helpers.Unauthorized:
		logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
		Unauthorized(w, r)
	case *helpers.NotFound:
		logs.Error("Request: %s, board not found for userID: %d", requestSummary(r), err)
		NotFound(w, r)
	case *helpers.BadRequest:
		logs.Error("Request: %s, unable to parse content: %v", requestSummary(r), err)
		BadRequest(w, r)
	}
}

func getUserIDIfAvailable(r *http.Request, al authz.AuthLayerInterface) (int, error) {
	token := r.Header.Get(authToken)
	if len(token) == 0 {
		return 0, nil
	}

	return getUserIdByToken(al, token)
}

func getUserIdByToken(al authz.AuthLayerInterface, token string) (int,
	error) {

	tokenData, err := al.GetTokenData(token)
	if err != nil {
		return 0, err
	}

	return tokenData.UserData.ID, nil
}
