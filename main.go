package main

import (
	"github.com/Nezent/GoPark/config"
	"github.com/Nezent/GoPark/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.ConnectDB()
	routes.ParkingRoute(router)
	router.Run(":8080")
}