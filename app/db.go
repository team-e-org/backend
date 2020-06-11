package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"app/logs"
)

func connectToDb(host string, port int, user string, password string, dbName string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logs.Info("Connected to DB %s:%d/%s...", host, port, dbName)
	return db, nil
}
