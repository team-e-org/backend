package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"app/logs"
)

func ConnectToMySql(host string, port int, user string, password string, dbName string, loc string) (*sql.DB, error) {
	loc = strings.Replace(loc, "/", "%2F", 1)
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&%s", user, password, host, port, dbName, loc)
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
