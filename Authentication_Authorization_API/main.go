package main

import (
	"Authentication_Authorization_API/controllers"
	"Authentication_Authorization_API/initializers"
	"Authentication_Authorization_API/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	router.Run()
}
