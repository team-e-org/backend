package handlers

import (
	"app/authz"
	"app/db"
	"app/logs"
	"app/models"
	"app/ptr"
	"app/view"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func ServePinsInBoard(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			Unauthorized(w, r)
			return
		}

		vars := mux.Vars(r)
		boardID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter id: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}

		page, err := strconv.Atoi(r.FormValue("page"))
		if err != nil {
			logs.Error("Request: %s, parse path parameter page: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}

		pins, err := data.Pins.GetPinsByBoardID(boardID, page)
		if err != nil {
			logs.Error("Request: %s, while gettign pins in board: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		pins = removePrivatePin(pins, userID)

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

func ServePin(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			Unauthorized(w, r)
			return
		}

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

		if pin.IsPrivate && pin.UserID != userID {
			logs.Error("Request: %s, pin not found in database: %v", requestSummary(r), pinID)
			NotFound(w, r)
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

func CreatePin(data db.DataStorage, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			Unauthorized(w, r)
			return
		}

		maxSize := int64(1024000)
		err = r.ParseMultipartForm(maxSize)
		if err != nil {
			logs.Error("Request: %s, parsing multipart: %v", requestSummary(r), err)
			logs.Error("Image too large. Max Size: %v", maxSize)
			BadRequest(w, r)
			return
		}

		vars := mux.Vars(r)
		boardID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter board id: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}

		var b bool
		if r.FormValue("isPrivate") != "" {
			b, err = strconv.ParseBool(r.FormValue("isPrivate"))
			if err != nil {
				logs.Error("Request: %s, parse parameter isPrivate: %v", requestSummary(r), err)
				BadRequest(w, r)
				return
			}
		}

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			logs.Error("Request: %s, getting uploaded image file: %v", requestSummary(r), err)
			BadRequest(w, r)
			return
		}
		defer file.Close()

		url, err := data.AWSS3.UploadImage(file, fileHeader)
		if err != nil {
			logs.Error("Request: %s, uploading image: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		now := time.Now()
		pin := &models.Pin{
			UserID:      userID,
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
			URL:         ptr.NewString(r.FormValue("url")),
			IsPrivate:   b,
			ImageURL:    url,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		pin, err = data.Pins.CreatePin(pin, boardID)
		if err != nil {
			logs.Error("Request: %s, creating pin: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		err = data.BoardsPins.CreateBoardPin(boardID, pin.ID)
		if err != nil {
			logs.Error("Request: %s, creating board_pin: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		response := view.NewPin(pin)
		bytes, err := json.Marshal(response)
		if err != nil {
			logs.Error("Request: %s, serializing pin response: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		w.Header().Set(contentType, jsonContent)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}

func removePrivatePin(pins []*models.Pin, userID int) []*models.Pin {
	for i, pin := range pins {
		if pin.IsPrivate && pin.UserID != userID {
			pins = append(pins[:i], pins[i+1:]...)
		}
	}

	return pins
}
