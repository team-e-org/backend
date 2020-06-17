package handlers

import (
	"app/authz"
	"app/db"
	"app/logs"
	"app/models"
	"app/view"
	"encoding/json"
	"net/http"
)

func CreateBoard(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			Unauthorized(w, r)
			return
		}

		requestBoard := &view.Board{}
		if err := json.NewDecoder(r.Body).Decode(requestBoard); err != nil {
			logs.Error("Request: %s, unable to parse content: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}

		storedBoard := &models.Board{
			UserID:      userID,
			Name:        requestBoard.Name,
			Description: requestBoard.Description,
			IsPrivate:   requestBoard.IsPrivate,
		}
		storedBoard, err = data.Boards.CreateBoard(storedBoard)
		if err != nil {
			logs.Error("Request: %s, creating board: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		bytes, err := json.Marshal(view.NewBoard(storedBoard))
		if err != nil {
			logs.Error("Request: %s, serializing book: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		w.Header().Set(contentType, jsonContent)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}
