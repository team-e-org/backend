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
)

func ServeTags(data db.DataStorageInterface, authLayer authz.AuthLayerInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		pinID, err := strconv.Atoi(r.FormValue("pin_id"))
		if err != nil {
			logs.Error("Request: %s, parse path parameter pinID: %v", requestSummary(r), err)
			err := helpers.NewBadRequest(err)
			ResponseError(w, r, err)
			return
		}

		tags, err := usecase.GetTagsByPinID(data, pinID)
		if err != nil {
			logs.Error("Request: %s, %v", requestSummary(r), err)
			ResponseError(w, r, err)
			return
		}

		bytes, err := json.Marshal(view.NewTags(tags))
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
