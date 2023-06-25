package model

type Garage struct {
	GarageId      string  `json:"garageId"  gorm:"default:uuid_generate_v4();primaryKey"`
	GarageName    string  `json:"garageName"`
	Latitude      float64 `json:"latitude"`
	Longituted    float64 `json:"longitute"`
	Level         int64   `json:"level"`         //level reuired to unlock the garage
	CoinsRequired int     `json:"coinsRequired"` //coins required to unlock the garage
	Locked        bool    `json:"locked"`
}

type GarageUpgrades struct {
	GarageId      string `json:"garageId"`
	UpgradeLevel  int    `json:"upgradeLevel"`  // level of the garage
	UpgradeAmount int64  `json:"upgradeAmount"` // amount required for the grage to be upgraded
	CarLimit      int    `json:"carLimit"`      //limit of cars a player can store in that garage
}
type GarageCarList struct {
	PlayerId string `json:"playerId"`
	GarageId string `json:"garageId"`
	CarId    string `json:"carId"`
}
