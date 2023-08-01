package model

type Garage struct {
	GarageId      string  `json:"garageId"  gorm:"default:uuid_generate_v4();primaryKey,omitempty"`
	GarageName    string  `json:"garageName,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longituted    float64 `json:"longitute,omitempty"`
	Level         uint64  `json:"level,omitempty"`         //level reuired to unlock the garage
	CoinsRequired uint64  `json:"coinsRequired,omitempty"` //coins required to unlock the garage
	Locked        bool    `json:"locked,omitempty"`
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
