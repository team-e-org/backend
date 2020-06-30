package main

import (
	"app/config"
	"app/db"
	"app/infrastructure"
	"app/logs"
	"app/server"
)

func main() {
	c, err := config.ReadConfig()
	if err != nil {
		logs.Error("Invalid config: %s", err)
		panic(err)
	}

	sqlDB, err := db.ConnectToMySql(c.DB)
	if sqlDB != nil {
		defer sqlDB.Close()
	}
	if err != nil {
		logs.Error("DB connection failure: %s", err)
		panic(err)
	}

	redis, err := db.ConnectToRedis(c.Redis)
	if redis != nil {
		defer redis.Close()
	}
	if err != nil {
		logs.Error("Redis connection failure: %s", err)
		panic(err)
	}

	s3 := infrastructure.NewAWSS3(c.AWS.S3)

	dynamo := db.ConnectToDynamo(c.AWS.Dynamo)

	if err = server.Start(c, sqlDB, redis, dynamo, s3); err != nil {
		logs.Error("Failed to start server: %s", err)
		panic(err)
	}
}
