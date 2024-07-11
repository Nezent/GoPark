package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParkingLotController(ctx *gin.Context){
	ctx.String(http.StatusOK,"Hello")
}
