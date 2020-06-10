package main

import (
	"app/config"
	"app/logs"
	"app/server"
)

func main() {
	c, err := config.ReadConfig()
	if err != nil {
		logs.Error("Invalid config: %s", err)
		panic(err)
	}

	if err = server.Start(c.Server.Port); err != nil {
		logs.Error("Failed to start server: %s", err)
		panic(err)
	}
}
