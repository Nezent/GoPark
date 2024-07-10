package models

import (
	"time"

	"github.com/google/uuid"
)

type ParkingLot struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string `gorm:"type:varchar(100)"`
	Slots int32  `gorm:"type:smallint"`
	IsAvailable bool `gorm:"type:boolean"`
	IsMaintainance bool `gorm:"type:boolean"`
}

type Slots struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ParkingLotID uuid.UUID `gorm:"type:uuid"`
	SlotID	int32 `gorm:"type:smallint"`
	ParkedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UnparkedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`
}

type Vehicle struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	VehicleType string `gorm:"type:varchar(100)"`	
}