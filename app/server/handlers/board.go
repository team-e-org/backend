package handlers

import (
	"app/authz"
	"app/db"
	"app/helpers"
	"app/logs"
	"app/usecase"
	"app/view"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBoard(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
			return
		}

		requestBoard := &view.Board{}
		if err := json.NewDecoder(r.Body).Decode(requestBoard); err != nil {
			logs.Error("Request: %s, unable to parse content: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		if userID != requestBoard.UserID {
			err := fmt.Errorf("UserIDs do not match error")
			logs.Error("Request: %s, an error occurred: %v", requestSummary(r), err)
			err = helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		board := view.NewBoardModel(requestBoard)
		storedBoard, err := usecase.CreateBoard(data, board)
		if err != nil {
			logs.Error("Request: %s, %v", requestSummary(r), err)
			ResponseError(w, r, err)
			return
		}

		bytes, err := json.Marshal(view.NewBoard(storedBoard))
		if err != nil {
			logs.Error("Request: %s, serializing board: %v", requestSummary(r), err)
			err := helpers.NewInternalServerError(err)
			ResponseError(w, r, err)
			return
		}

		w.Header().Set(contentType, jsonContent)
		w.WriteHeader(http.StatusCreated)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}

func UpdateBoard(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		vars := mux.Vars(r)
		boardID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter id: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
		}

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
			return
		}

		requestBoard := &view.Board{}
		if err := json.NewDecoder(r.Body).Decode(requestBoard); err != nil {
			logs.Error("Request: %s, unable to parse content: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		if userID != requestBoard.UserID {
			logs.Error("Request: %s, forbidden: %v", requestSummary(r), err)
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
			return
		}

		if boardID != requestBoard.ID {
			err := fmt.Errorf("BoardIDs do not match error")
			logs.Error("Request: %s, an error occurred: %v", requestSummary(r), err)
			err = helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		newBoard := view.NewBoardModel(requestBoard)

		storedBoard, err := usecase.UpdateBoard(data, newBoard)
		if err != nil {
			logs.Error("Request: %s, %v", requestSummary(r), err)
			ResponseError(w, r, err)
			return
		}

		bytes, err := json.Marshal(view.NewBoard(storedBoard))
		if err != nil {
			logs.Error("Request: %s, serializing board: %v", requestSummary(r), err)
			err := helpers.NewInternalServerError(err)
			ResponseError(w, r, err)
			return
		}

		w.Header().Set(contentType, jsonContent)
		w.WriteHeader(http.StatusCreated)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}

func SavePin(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		boardID, pinID, err := getBoardIdAndPinId(r)
		if err != nil {
			logs.Error("Request: %s, an error occurred: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		err = usecase.SavePin(data, boardID, pinID)
		if err != nil {
			logs.Error("Request: %s, failed to save pin: %v", requestSummary(r), err)
			ResponseError(w, r, err)
			return
		}

		w.Header().Set(contentType, jsonContent)
		w.WriteHeader(http.StatusCreated)
		if _, err = w.Write([]byte("{}")); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}

func UnsavePin(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		boardID, pinID, err := getBoardIdAndPinId(r)
		if err != nil {
			logs.Error("Request: %s, an error occurred: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		err = usecase.UnsavePin(data, boardID, pinID)
		if err != nil {
			logs.Error("Request: %s, failed to unsave pin: %v", requestSummary(r), err)
			ResponseError(w, r, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func getBoardIdAndPinId(r *http.Request) (int, int, error) {
	vars := mux.Vars(r)
	boardID, err := strconv.Atoi(vars["boardID"])
	if err != nil {
		return 0, 0, fmt.Errorf("boardID is invalid: %v", err)
	}
	pinID, err := strconv.Atoi(vars["pinID"])
	if err != nil {
		return 0, 0, fmt.Errorf("pinID is invalid: %v", err)
	}

	return boardID, pinID, nil
}
