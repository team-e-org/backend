package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server Server
	DB     DBConfig
	Redis  RedisConfig
	AWS    AWS
}

type Server struct {
	Port int
}

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	TimeZone string
}

type RedisConfig struct {
	Host string
	Port int
}

type AWS struct {
	S3 S3
}

type S3 struct {
	Region          string
	Bucket          string
	PinFolder       string
	AccessKeyID     string
	SecretAccessKey string
}

func readDBConfig() (*DBConfig, error) {
	dbPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'MYSQL_PORT': %w", err)
	}

	dbConfig := &DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     dbPort,
		TimeZone: os.Getenv("TZ"),
	}

	return dbConfig, nil
}

func readRedisConfig() (*RedisConfig, error) {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'REDIS_PORT': %w", err)
	}

	redisConfig := &RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: port,
	}

	return redisConfig, nil
}

func readAWSConfig() *AWS {
	awsConfig := &AWS{
		S3{
			"ap-northeast-1",
			"pinko-bucket",
			"pins/",
			os.Getenv("AWS_S3_ACCESS_KEY_ID"),
			os.Getenv("AWS_S3_SECRET_ACCESS_KEY"),
		},
	}

	return awsConfig
}

func ReadConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'SERVER_PORT': %w", err)
	}

	dbConfig, err := readDBConfig()
	if err != nil {
		return nil, err
	}

	redisConfig, err := readRedisConfig()
	if err != nil {
		return nil, err
	}

	awsConfig := readAWSConfig()

	return &Config{
		Server{
			Port: port,
		},
		*dbConfig,
		*redisConfig,
		*awsConfig,
	}, nil
}
