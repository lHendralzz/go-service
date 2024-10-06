package main

import (
	restHandler "go-service/handler/rest"
	"go-service/repository"
	"go-service/service"

	"github.com/gin-gonic/gin" // Gin framework
)

func main() {
	// init repo
    repo := repository.Init()	

	// init service
	svc := service.Init(repo)
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	rest := restHandler.Init(svc, router)

	rest.Run()
}
