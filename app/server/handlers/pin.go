package handlers

import (
	"app/authz"
	"app/db"
	"app/logs"
	"app/view"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ServePinsInBoard(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		vars := mux.Vars(r)
		boardID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter id: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}

		pins, err := data.Pins.GetPinsByBoardID(boardID)
		if err != nil {
			logs.Error("Request: %s, while gettign pins in board: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}
		if len(pins) == 0 {
			logs.Error("Request: %s, pin not found in boardID: %v", requestSummary(r), boardID)
			NotFound(w, r)
			return
		}

		bytes, err := json.Marshal(view.NewPins(pins))
		if err != nil {
			logs.Error("Request: %s, serializing pins: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		w.Header().Set(contentType, jsonContent)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}
