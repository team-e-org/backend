package main

import (
	"app/config"
	"app/db"
	"app/logs"
	"app/server"
)

func main() {
	c, err := config.ReadConfig()
	if err != nil {
		logs.Error("Invalid config: %s", err)
		panic(err)
	}

	sqlDB, err := db.ConnectToMySql(c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.DBName)
	if sqlDB != nil {
		defer sqlDB.Close()
	}
	if err != nil {
		logs.Error("DB connection failure: %s", err)
		panic(err)
	}

	if err = server.Start(c.Server.Port, sqlDB); err != nil {
		logs.Error("Failed to start server: %s", err)
		panic(err)
	}
}
