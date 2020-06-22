package server

import (
	"app/authz"
	"app/authz/middleware"
	"app/config"
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

func Start(port int, dbConn *sql.DB, awsConf *config.AWS) error {
	router := mux.NewRouter()
	data := db.NewDataStorage(dbConn, awsConf)
	authLayer := authz.NewAuthLayer(*data)
	attachHandlers(router, data, authLayer)
	attachReqAuth(router, data, authLayer)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return s.ListenAndServe()
}

func attachHandlers(mux *mux.Router, data *db.DataStorage, al authz.AuthLayerInterface) {
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/users/sign-in", handlers.SignIn(*data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/users/sign-up", handlers.SignUp(*data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/users/{id}", handlers.ServeUser(*data, al)).Methods(http.MethodGet)
	mux.HandleFunc("/pins/{id}", handlers.ServePin(*data, al)).Methods(http.MethodGet)
	mux.HandleFunc("/boards/{id}/pins", handlers.ServePinsInBoard(*data, al)).Methods(http.MethodGet).Queries("page", "{page}")
	mux.HandleFunc("/users/{id}/boards", handlers.UserBoards(*data, al)).Methods(http.MethodGet)
}

func attachReqAuth(mux *mux.Router, data *db.DataStorage, al authz.AuthLayerInterface) {
	muxAuth := mux.PathPrefix("").Subrouter()
	muxAuth.Use(middleware.RequireAuthorization(al))

	mux.HandleFunc("/boards", handlers.CreateBoard(*data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/boards/{id}/pins", handlers.CreatePin(*data, al)).Methods(http.MethodPost)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
