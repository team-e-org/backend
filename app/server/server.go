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
	data := db.NewDataStorage(dbConn)
	s3 := db.NewAwsS3(awsConf.S3)
	authLayer := authz.NewAuthLayer(*data)
	attachHandlers(router, data, authLayer)
	attachReqAuth(router, data, authLayer, s3)

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
	mux.HandleFunc("/pins/{id}", handlers.ServePin(*data, al)).Methods(http.MethodGet)
	mux.HandleFunc("/boards/{id}/pins", handlers.ServePinsInBoard(*data, al)).Methods(http.MethodGet)
}

func attachReqAuth(mux *mux.Router, data *db.DataStorage, al authz.AuthLayerInterface, s3 *db.AwsS3) {
	muxAuth := mux.PathPrefix("").Subrouter()
	muxAuth.Use(middleware.RequireAuthorization(al))

	mux.HandleFunc("/boards", handlers.CreateBoard(*data, al)).Methods(http.MethodPost)
	mux.HandleFunc("/boards/{id}/pins", handlers.CreatePin(*data, al, s3)).Methods(http.MethodPost)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
