package handlers

import (
	"app/authz"
	"app/db"
	"github.com/gorilla/mux"
	"net/http"
)

func UserBoards(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["id"]

		w.Header().Set(contentType, jsonContent)
		w.Write([]byte("GET /users/${id}/boards. User ID: " + userId))
	}
}
