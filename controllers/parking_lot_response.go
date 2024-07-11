package controllers

import (
	"net/http"

	"github.com/Nezent/GoPark/config"
	"github.com/Nezent/GoPark/models"
	"github.com/gin-gonic/gin"
)


func CreateParkingLot(ctx *gin.Context) {
	request := models.CreateParkingLotRequest{}
	ctx.Bind(&request)

	parkingLot := models.ParkingLot{
		Name: request.Name,
		Slots: int32(request.DesiredSlots),
	}

	config.DB.Create(&parkingLot)

	slots := []models.SlotResponse{}

	for i := 1; i <= request.DesiredSlots; i++ {
		slot := models.Slots{
			ParkingLotID: parkingLot.ID,
			SlotID: int32(i),
		}
		config.DB.Create(&slot)

		slots = append(slots, models.SlotResponse{
			ID: slot.ID,
			SlotNumber: slot.SlotID,
			IsAvailable: true,
			IsMaintenance: false,
		})
	}

	response := models.ParkingLotResponse{
		ID: parkingLot.ID,
		Name: parkingLot.Name,
		DesiredSlots: request.DesiredSlots,
		Slots: slots,
	}
	ctx.JSON(http.StatusCreated,response)
}
