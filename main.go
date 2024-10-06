package main

import (
	"github.com/gin-gonic/gin" // Gin framework
)

func main() {
    router := gin.Default()
/*
    // Initialize repository and service
    userRepo := repository.NewUserRepository()
    userService := services.NewUserService(userRepo)	

    // Inject service into the controller
    userController := controllers.NewUserController(userService)

    // Set up routes
    router.GET("/user/:id", userController.GetUser)
*/
    // Run the server
    router.Run(":8080")
}
