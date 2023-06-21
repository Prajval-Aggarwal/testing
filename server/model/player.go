package model

import "github.com/google/uuid"

type Player struct {
	PlayerId   uuid.UUID `json:"playerId" gorm:"primaryKey"`
	PlayerName string    `json:"playerName" gorm:"unique"`
	Role       string    `json:"role"`
	Email      string    `json:"email"`
	Coins      int64     `json:"coins"`
	Cash       int64     `json:"cash"`
	DeviceId   string    `json:"deviceId"`
	OS         int64     `json:"os"` // o for android 1 for ios
}

type PlayerCars struct {
	PlayerId string `json:"playerId"`
	CarId    string `json:"carId"`
}
