package handlers

import (
	"app/authz"
	"app/db"
	"app/logs"
	"app/models"
	"app/view"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UserBoards(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])

		if err != nil {
			logs.Error("Request: %s, parse path parameter id: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}

		boards, err := data.Boards.GetBoardsByUserID(userID)
		if err != nil {
			logs.Error("Request: %s, while getting user's boards: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		tokenUser, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			Unauthorized(w, r)
			return
		}

		boards = removePrivateBoards(boards, tokenUser)

		if len(boards) == 0 {
			logs.Error("Request: %s, board not found for userID: %d", requestSummary(r), userID)
			NotFound(w, r)
			return
		}

		bytes, err := json.Marshal(view.NewBoards(boards))
		if err != nil {
			logs.Error("Request: %s, serializing boards: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		w.Header().Set(contentType, jsonContent)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}

func removePrivateBoards(boards []*models.Board, userID int) []*models.Board {
	for i, board := range boards {
		if board.IsPrivate && board.UserID != userID {
			boards = append(boards[:i], boards[i+1:]...)
		}
	}

	return boards
}
