package main

import (
	"fmt"
	restHandler "go-service/handler/rest"
	"go-service/handler/scheduler"
	"go-service/stdlib/config"
	"go-service/stdlib/database"
	"go-service/stdlib/redis"
	"go-service/usecase"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func parseConfig(logger *logrus.Logger) Conf {
	err := godotenv.Load(".env") // convert env file to OS ENV
	if err != nil {
		logger.Debug(fmt.Errorf("failed open .env file: %v", err))
	}
	conf := Conf{}

	logger.Debug("parse config")
	err = config.LoadConfig(&conf)
	if err != nil {
		logger.Debug(fmt.Errorf("failed parse Config: %v", err))
		return conf
	}

	logger.Debug(fmt.Sprintf("Success Parse Config"))
	return conf
}

type Conf struct {
	Database  database.Option
	Rest      restHandler.Option
	Scheduler scheduler.Option

	Redis redis.Option
	// Business Config
	Business
}

type Business struct {
	Usecase usecase.Options
}
