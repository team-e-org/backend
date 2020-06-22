package handlers

import (
	"app/authz"
	"app/db"
	"app/helpers"
	"app/logs"
	"app/usecase"
	"app/view"
	"encoding/json"
	"net/http"
)

func CreateBoard(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
		}

		requestBoard := &view.Board{}
		if err := json.NewDecoder(r.Body).Decode(requestBoard); err != nil {
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
		}

		storedBoard, err := usecase.CreateBoard(data, requestBoard, userID)
		if err != nil {
			ResponseError(w, r, err)
		}

		bytes, err := json.Marshal(view.NewBoard(storedBoard))
		if err != nil {
			err := helpers.NewInternalServerError(err)
			ResponseError(w, r, err)
		}

		w.Header().Set(contentType, jsonContent)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}
