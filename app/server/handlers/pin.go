package handlers

import (
	"app/authz"
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"app/ptr"
	"app/usecase"
	"app/view"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func ServePinsInBoard(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
			return
		}

		vars := mux.Vars(r)
		boardID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter id: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		page, err := strconv.Atoi(r.FormValue("page"))
		if err != nil {
			logs.Error("Request: %s, parse path parameter page: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		pins, err := usecase.GetPinsByBoardID(data, userID, boardID, page)
		if err != nil {
			logs.Error("Request: %s, %v", requestSummary(r), err)
			ResponseError(w, r, err)
			return
		}

		bytes, err := json.Marshal(view.NewPins(pins))
		if err != nil {
			logs.Error("Request: %s, serializing pins: %v", requestSummary(r), err)
			err := helpers.NewInternalServerError(err)
			ResponseError(w, r, err)
			return
		}

		w.Header().Set(contentType, jsonContent)
		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}

func ServePin(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
			return
		}

		vars := mux.Vars(r)
		pinID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter id: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		pin, err := usecase.ServePin(data, pinID, userID)
		if err != nil {
			logs.Error("Request: %s, %v", requestSummary(r), err)
			ResponseError(w, r, err)
			return
		}

		bytes, err := json.Marshal(view.NewPin(pin))

		if err != nil {
			logs.Error("Request: %s, serializing pin: %v", requestSummary(r), err)
			err := helpers.NewInternalServerError(err)
			ResponseError(w, r, err)
			return
		}

		w.Header().Set(contentType, jsonContent)

		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}

func CreatePin(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		userID, err := getUserIDIfAvailable(r, authLayer)
		if err != nil {
			logs.Error("Request: %s, checking if user identifiable: %v", requestSummary(r), err)
			err := helpers.NewUnauthorized(err)
			ResponseError(w, r, err)
			return
		}

		maxSize := int64(1024000)
		err = r.ParseMultipartForm(maxSize)
		if err != nil {
			logs.Error("Request: %s, parsing multipart: %v", requestSummary(r), err)
			logs.Error("Image too large. Max Size: %v", maxSize)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		vars := mux.Vars(r)
		boardID, err := strconv.Atoi(vars["id"])
		if err != nil {
			logs.Error("Request: %s, parse path parameter board id: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		var b bool
		if r.FormValue("isPrivate") != "" {
			b, err = strconv.ParseBool(r.FormValue("isPrivate"))
			if err != nil {
				logs.Error("Request: %s, parse parameter isPrivate: %v", requestSummary(r), err)
				err := helpers.NewBadRequest(err)
				ResponseError(w, r, err)
				return
			}
		}

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			logs.Error("Request: %s, getting uploaded image file: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}
		defer file.Close()

		url, err := data.AWSS3().UploadImage(file, fileHeader)
		if err != nil {
			logs.Error("Request: %s, uploading image: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		now := time.Now()
		pin := &models.Pin{
			UserID:      ptr.NewInt(userID),
			Title:       r.FormValue("title"),
			Description: ptr.NewString(r.FormValue("description")),
			URL:         ptr.NewString(r.FormValue("url")),
			IsPrivate:   b,
			ImageURL:    url,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		pin, err = usecase.CreatePin(data, pin, boardID)

		response := view.NewPin(pin)
		bytes, err := json.Marshal(response)
		if err != nil {
			logs.Error("Request: %s, serializing pin response: %v", requestSummary(r), err)
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
