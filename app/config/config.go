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
	S3     S3
	Lambda Lambda
}

type S3 struct {
	Region    string
	BaseURL   string
	Bucket    string
	PinFolder string
}

type Lambda struct {
	Region         string
	FunctionARN    string
	InvocationType string
}

func ReadDBConfig() (*DBConfig, error) {
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

func ReadRedisConfig() (*RedisConfig, error) {
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

func ReadS3Config() S3 {
	baseURL := os.Getenv("CLOUDFRONT_URL")
	bucket := os.Getenv("S3_BUCKET")
	return S3{
		Region:    "ap-northeast-1",
		BaseURL:   baseURL,
		Bucket:    bucket,
		PinFolder: "pins",
	}
}

func ReadLambdaConfig() Lambda {
	return Lambda{
		Region:         "ap-northeast-1",
		FunctionARN:    "arn:aws:lambda:ap-northeast-1:444207867088:function:attachTag",
		InvocationType: "Event",
	}
}

func ReadAWSConfig() *AWS {
	return &AWS{
		ReadS3Config(),
		ReadLambdaConfig(),
	}
}

func ReadConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'SERVER_PORT': %w", err)
	}

	dbConfig, err := ReadDBConfig()
	if err != nil {
		return nil, err
	}

	redisConfig, err := ReadRedisConfig()
	if err != nil {
		return nil, err
	}

	awsConfig := ReadAWSConfig()

	return &Config{
		Server{
			Port: port,
		},
		*dbConfig,
		*redisConfig,
		*awsConfig,
	}, nil
}
