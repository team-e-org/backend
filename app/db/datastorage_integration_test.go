// +build integration

package db

import (
	"app/config"
	"database/sql"
	"log"
	"os"
	"testing"
)

var sqlDB *sql.DB

func dbHandlingWrapper(m *testing.M) int {
	c, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	sqlDB, err := ConnectToMySql(c.DB)
	if err != nil {
		log.Panicf("Can not connect to DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(dbHandlingWrapper(m))
}

func TestDataStorage(t *testing.T) {

}
