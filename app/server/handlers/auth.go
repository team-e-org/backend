package handlers

import (
	"app/authz"
	"app/db"
	"app/logs"
	"app/models/view"
	"encoding/json"
	"net/http"
)

func SignIn(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		signInRequest, err := view.NewSignInRequest(r.Body)
		if err != nil {
			logs.Error("Request: %s, getting user data from req body: %v", requestSummary(r), err)
			BadRequest(w, r)
		}

		token, err := authLayer.AuthenticateUser(signInRequest.Email, signInRequest.Password)
		if err != nil {
			logs.Error("Request: %s, authenticate user: %v", requestSummary(r), err)
			Unauthorized(w, r)
		}

		response := view.NewLSignInResponse(token)
		bytes, err := json.Marshal(response)
		if err != nil {
			logs.Error("Request: %s, serializing login response: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		w.Header().Set(contentType, jsonContent)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}
