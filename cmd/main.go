package main

import (
	"flag"
	_ "go-service/docs"
	restHandler "go-service/handler/rest"
	"go-service/handler/scheduler"
	"go-service/repository"
	"go-service/service"
	"go-service/stdlib/database"
	"go-service/stdlib/logger"

	redisLock "go-service/stdlib/redis"

	"github.com/gin-gonic/gin" // Gin framework
	"github.com/go-redis/redis/v8"
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
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	// init a RedisLock instance
	redisLock := redisLock.NewRedisLock(rdb)

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
	repo := repository.Init(db, logger)

	// init service
	svc := service.Init(repo, logger, conf.Service, redisLock)
	// init router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// init Scheduler
	scheduler := scheduler.Init(svc, logger, conf.Scheduler)
	scheduler.Run()

	// init rest
	rest := restHandler.Init(svc, router, logger, conf.Rest)
	rest.Run()

}
