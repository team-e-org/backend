package handlers

import (
	"app/authz"
	"app/db"
	"app/logs"
	"app/view"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ServePin(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		vars := mux.Vars(r)
		pinID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter id: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}

		pin, err := data.Pins.GetPin(pinID)
		if err == sql.ErrNoRows {
			logs.Error("Request: %s, pin not found in database: %v", requestSummary(r), pinID)
			NotFound(w, r)
			return
		}
		if err != nil {
			logs.Error("Request: %s, getting pin from database: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		bytes, err := json.Marshal(view.NewPin(pin))

		if err != nil {
			logs.Error("Request: %s, serializing pin: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		w.Header().Set(contentType, jsonContent)

		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}
