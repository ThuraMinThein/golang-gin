package main

import (
	"github.com/ThuraMinThein/golang-gin/config"
	"github.com/ThuraMinThein/golang-gin/internal/database"
	"github.com/ThuraMinThein/golang-gin/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.Init()
	
	r := gin.Default()

	routes.SetupRoutes(r)
	r.Run()
}