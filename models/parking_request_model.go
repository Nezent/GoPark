package models

import "github.com/google/uuid"

type CreateParkingLotRequest struct {
	Name         string `json:"name"`
	DesiredSlots int    `json:"desiredSlots"`
}

type SlotResponse struct {
	ID            uuid.UUID `json:"id"`
	SlotNumber    int32     `json:"slotNumber"`
	IsAvailable   bool      `json:"isAvailable"`
	IsMaintenance bool      `json:"isMaintenance"`
}

type ParkingLotResponse struct {
	ID           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	DesiredSlots int            `json:"desiredSlots"`
	Slots        []SlotResponse `json:"slots"`
}