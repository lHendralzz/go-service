package main

import (
	"flag"
	_ "go-service/docs"
	restHandler "go-service/handler/rest"
	"go-service/handler/scheduler"
	"go-service/repository"
	"go-service/stdlib/database"
	"go-service/stdlib/logger"
	"go-service/usecase"
	"strconv"

	redisLock "go-service/stdlib/redis"

	"github.com/gin-gonic/gin" // Gin framework
	"github.com/go-redis/redis/v8"
)

// @title Swagger for product-service
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

	dbredis, err := strconv.Atoi(conf.Redis.DB)
	if err != nil {
		logger.Error(err)
		return
	}
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,     // Redis server address
		Password: conf.Redis.Password, // No password
		DB:       dbredis,             // Default DB
	})

	// init a RedisLock instance
	redisLock := redisLock.NewRedisLock(rdb)

	// init db
	logger.Info(conf.Database)
	db := database.InitDB(logger, conf.Database)

	// init repo
	repo := repository.Init(db, logger)

	// init service
	uc := usecase.Init(repo, logger, conf.Usecase, redisLock)
	// init router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// init Scheduler
	scheduler := scheduler.Init(uc, logger, conf.Scheduler)
	scheduler.Run()

	// init rest
	rest := restHandler.Init(uc, router, logger, conf.Rest)
	rest.Run()

}
