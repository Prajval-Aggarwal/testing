package model

import "github.com/google/uuid"

type Player struct {
	PlayerId   uuid.UUID `json:"playerId"`
	PlayerName string    `json:"playerName" gorm:"unique"`
	Level      int       `json:"level"`
	Role       string    `json:"role"`
	Email      string    `json:"email"`
	Coins      int64     `json:"coins"`
	Cash       int64     `json:"cash"`
	DeviceId   string    `json:"deviceId"`
	OS         int64     `json:"os"` // o for android 1 for ios
}

type OwnedCars struct {
	PlayerId string `json:"playerId"`
	CarId    string `json:"carId"`
	Selected bool   `json:"selected"`
}

type OwnedGarage struct {
	PlayerId    string `json:"playerId"`
	GarageId    string `json:"garageId"`
	GarageLevel int    `json:"garageLevel"`
	CarLimit    int    `json:"carLimit"`
}
type PlayerCarsStats struct {
	PlayerId    string    `json:"playerId"`
	CarId       uuid.UUID `json:"carId"`
	Power       int64     `json:"power"`
	Grip        int64     `json:"grip"`
	ShiftTime   float64   `json:"shiftTime"`
	Weight      int64     `json:"weight"`
	OR          float64   `json:"or"` //overall rating of the car
	Durability  int64     `json:"Durability"`
	NitrousTime int       `json:"nitrousTime"` //increased when nitrous is upgraded

}

type PlayerCarUpgrades struct {
	PlayerId     string `json:"playerId"`
	CarId        string `json:"carId"`
	Engine       int    `json:"engine"`       // Affects Power
	Turbo        int    `json:"turbo"`        // Affects Power
	Intake       int    `json:"intake"`       // Affects Power
	Nitrous      int    `json:"nitrous"`      // Affect Nitrous time
	Body         int    `json:"body"`         //Affects Grip and Weight
	Tires        int    `json:"tires"`        //Affects Grip
	Transmission int    `json:"transmission"` //Affects Shift-Time
}
