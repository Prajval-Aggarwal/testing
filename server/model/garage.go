package model

import "time"

type Garage struct {
	GarageId      string    `json:"garageId"  gorm:"default:uuid_generate_v4();primaryKey"`
	GarageName    string    `json:"garageName,omitempty"`
	GarageType    int64     `json:"garageType,omitempty"`
	Latitude      float64   `json:"latitude,omitempty"`
	Longitude     float64   `json:"longitude,omitempty"`
	Level         uint64    `json:"level,omitempty"`         //level reuired to unlock the garage
	CoinsRequired uint64    `json:"coinsRequired,omitempty"` //coins required to unlock the garage
	Locked        bool      `json:"locked,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
}

type GarageUpgrades struct {
	GarageId      string `json:"garageId,omitempty"`
	UpgradeLevel  uint64 `json:"upgradeLevel,omitempty"`  // level of the garage
	UpgradeAmount uint64 `json:"upgradeAmount,omitempty"` // amount required for the grage to be upgraded
	CarLimit      uint64 `json:"carLimit,omitempty"`      //limit of cars a player can store in that garage
}
type GarageCarList struct {
	PlayerId string `json:"playerId,omitempty"`
	GarageId string `json:"garageId,omitempty"`
	CarId    string `json:"carId,omitempty"`
}

type GarageTypes struct {
	TypeName string `json:"label,omitempty" gorm:"unique"`
	TypeId   int    `json:"value"`
}
