package model

import "time"

type Garage struct {
	GarageId      string    `json:"garageId"  gorm:"default:uuid_generate_v4();primaryKey,omitempty"`
	GarageName    string    `json:"garageName,omitempty"`
	Latitude      float64   `json:"latitude,omitempty"`
	Longitude     float64   `json:"longitude,omitempty"`
	Level         int64     `json:"level,omitempty"`         //level reuired to unlock the garage
	CoinsRequired int64     `json:"coinsRequired,omitempty"` //coins required to unlock the garage
	Locked        bool      `json:"locked,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
}

type GarageUpgrades struct {
	GarageId      string `json:"garageId,omitempty"`
	UpgradeLevel  int    `json:"upgradeLevel,omitempty"`  // level of the garage
	UpgradeAmount int64  `json:"upgradeAmount,omitempty"` // amount required for the grage to be upgraded
	CarLimit      int    `json:"carLimit,omitempty"`      //limit of cars a player can store in that garage
}
type GarageCarList struct {
	PlayerId string `json:"playerId,omitempty"`
	GarageId string `json:"garageId,omitempty"`
	CarId    string `json:"carId,omitempty"`
}
