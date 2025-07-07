package main

import (
	"Authentication_Authorization_API/initializers"
	"Authentication_Authorization_API/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
