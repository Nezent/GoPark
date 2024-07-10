package routes

import (
	"github.com/Nezent/GoPark/controllers"
	"github.com/gin-gonic/gin"
)

func ParkingRoute(router *gin.Engine){
	router.GET("/",controllers.ParkingLotController)
}