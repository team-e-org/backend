package server

import (
	"app/authz"
	"app/authz/middleware"
	"app/config"
	"app/db"
	"app/repository"
	"app/server/handlers"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/guregu/dynamo"
)

type S3 repository.FileRepository
type Lambda repository.LambdaRepository

func init() {
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
}

func Start(config *config.Config, dbConn *sql.DB, redis *redis.Client, dynamo *dynamo.DB, s3 S3, lambda Lambda) error {
	router := mux.NewRouter()
	data := db.NewDataStorage(dbConn, dynamo, s3)
	authLayer := authz.NewAuthLayer(data, redis)
	attachHandlers(router, data, authLayer)
	attachReqAuth(router, data, authLayer, lambda)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: router,
	}

	return s.ListenAndServe()
}

func attachHandlers(mux *mux.Router, data db.DataStorageInterface, al authz.AuthLayerInterface) {
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/users/sign-in", handlers.SignIn(data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/users/sign-up", handlers.SignUp(data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/users/{id}", handlers.ServeUser(data, al)).Methods(http.MethodGet)
	mux.HandleFunc("/pins/{id}", handlers.ServePin(data, al)).Methods(http.MethodGet)
	mux.HandleFunc("/pins", handlers.ServePins(data, al)).Methods(http.MethodGet)
	mux.HandleFunc("/boards/{id}/pins", handlers.ServePinsInBoard(data, al)).Methods(http.MethodGet).Queries("page", "{page}")
	mux.HandleFunc("/users/{id}/boards", handlers.UserBoards(data, al)).Methods(http.MethodGet)
	mux.HandleFunc("/tags", handlers.ServeTags(data, al)).Methods(http.MethodGet).Queries("pin_id", "{pin_id}")
}

func attachReqAuth(mux *mux.Router, data db.DataStorageInterface, al authz.AuthLayerInterface, lambda Lambda) {
	muxAuth := mux.PathPrefix("").Subrouter()
	muxAuth.Use(middleware.RequireAuthorization(al))

	mux.HandleFunc("/boards", handlers.CreateBoard(data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/boards/{id}/pins", handlers.CreatePin(data, al, lambda)).Methods(http.MethodPost)
	mux.HandleFunc("/boards/{boardID}/pins/{pinID}", handlers.SavePin(data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/boards/{boardID}/pins/{pinID}", handlers.UnsavePin(data, al)).Methods(http.MethodDelete)
	mux.HandleFunc("/boards/{id}", handlers.UpdateBoard(data, al)).Methods(http.MethodPut)
	mux.HandleFunc("/pins/{id}", handlers.UpdatePin(data, al)).Methods(http.MethodPut)
	mux.HandleFunc("/users/{id}", handlers.UpdateUser(data, al)).Methods(http.MethodPut)
	mux.HandleFunc("/pins", handlers.ServeHomePins(data, al)).Methods(http.MethodPost)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
