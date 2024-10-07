package main

import (
	"flag"
	_ "go-service/docs"
	restHandler "go-service/handler/rest"
	"go-service/repository"
	"go-service/service"
	"go-service/stdlib/database"
	"go-service/stdlib/logger"

	"github.com/gin-gonic/gin" // Gin framework
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
	logger := logger.Init(*debug)

	// parsing conf from .env and OS env
	conf := parseConfig(logger)

	// init db
	logger.Info(conf.Database)
	db := database.InitDB(logger, conf.Database)

	// init repo
	repo := repository.Init(db)

	// init service
	svc := service.Init(repo, logger, conf.Service)
	// init router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// ini rest
	rest := restHandler.Init(svc, router, logger)

	rest.Run()
}
