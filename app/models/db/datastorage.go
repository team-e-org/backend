package db

import (
	"database/sql"
)

type DataStorage interface{}

func NewSQLDataStorage(sqlDB *sql.DB) SQLDataStorage {
	return SQLDataStorage{DB: sqlDB}
}

type SQLDataStorage struct {
	DB *sql.DB
}
