package server

import (
	"app/authz"
	"app/db"
	"app/server/handlers"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
}

func Start(port int, dbConn *sql.DB) error {
	router := mux.NewRouter()
	data := db.NewDataStorage(dbConn)
	authLayer := authz.NewAuthLayer(*data)
	attachHandlers(router, data, authLayer)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return s.ListenAndServe()
}

func attachHandlers(mux *mux.Router, data *db.DataStorage, al authz.AuthLayerInterface) {
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/users/sign-in", handlers.SignIn(*data, al)).Methods(http.MethodPost)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
