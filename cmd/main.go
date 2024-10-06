package main

import (
	"flag"
	"fmt"
	_ "go-service/docs"
	restHandler "go-service/handler/rest"
	"go-service/repository"
	"go-service/service"
	"go-service/stdlib/config"
	"go-service/stdlib/logger"

	"github.com/gin-gonic/gin" // Gin framework
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title Swagger for go-service
// @version 1.0
// @description Swagger for backend API service
// @description Get the Bearer token on the Authentication Service
// @description JSON Link: <a href=/swagger/doc.json>docs.json</a>

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @docExpansion none
func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	// init logger
	logger := logger.Init(*debug);
	// parsing conf from .env and OS env
	conf := parseConfig(logger)

	// init repo
    repo := repository.Init()	

	// init service
	svc := service.Init(repo, conf.Service)
	// init router 
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// ini rest 
	rest := restHandler.Init(svc, router, logger)

	rest.Run()
}

func parseConfig(logger *logrus.Logger) Conf {
	err := godotenv.Load(".env") // convert env file to OS ENV
	if err != nil {
		logger.Debug(fmt.Errorf("failed open .env file: %v", err))
	}
	conf := Conf{}

	logger.Debug("parse config")
	err = config.LoadConfig(&conf)
	if err != nil{
		logger.Debug(fmt.Errorf("failed parse Config: %v", err))
		return conf
	}

	logger.Debug(fmt.Sprintf("Success Parse Config"))
	return conf
}

type Conf struct{
	// Business Config
	Business
}

type Business struct{
	Service service.Options
}