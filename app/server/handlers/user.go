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
	"strconv"

	"github.com/gorilla/mux"
)

func UserBoards(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])
		currentUserID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
			return
		}

		boards, err := usecase.UserBoards(data, authLayer, userID, currentUserID)
		if err != nil {
			ResponseError(w, r, err)
			return
		}

		bytes, err := json.Marshal(view.NewBoards(boards))
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
