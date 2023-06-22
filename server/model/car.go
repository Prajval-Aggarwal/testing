package model

import "github.com/google/uuid"

type Car struct {
	CarId      uuid.UUID `json:"carId"`
	CarName    string    `json:"carName"`
	CurrLevel  int       `json:"currLevel"` //current level of the car
	Level      int64     `json:"level"`     // level required to unlock the car
	CurrType   string    `json:"currType"`
	CurrAmount float64   `json:"currAmount"`
	MaxLevel   int64     `json:"maxLevel"`
	Class      string    `json:"class"`
	Locked     bool      `json:"locked"`
}

type CarStats struct {
	CarId      uuid.UUID `json:"carId"`
	Power      int64     `json:"power"`
	Grip       int64     `json:"grip"`
	ShiftTime  float64   `json:"shiftTime"`
	Weight     int64     `json:"weight"`
	OR         float64   `json:"or"` //overall rating of the car
	Durability int64     `json:"Durability"`
	RepairCost int64     `json:"RepairCost"`
}
