package server

import (
	"app/models/db"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(port int, dbConn *sql.DB) error {
	router := mux.NewRouter()
	data := db.NewSQLDataStorage(dbConn)
	attachHandlers(router, data)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return s.ListenAndServe()
}

func attachHandlers(mux *mux.Router, data db.DataStorage) {
	mux.HandleFunc("/", Hello)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
