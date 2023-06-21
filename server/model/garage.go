package model

type Garage struct {
	GarageId   string `json:"garageId"`
	GarageName string `json:"garageName"`
	Latitude   string `json:"latitude"`
	Longituted string `json:"longitute"`
	Level      string `json:"level"` //level reuired to unlock the garage

}

type GarageUpgrades struct {
	GarageId      string `json:"garageId"`
	UpgradeLevel  int    `json:"upgradeLevel"`  // level of the garage
	UpgradeAmount int64  `json:"upgradeAmount"` // amount required for the grage to be upgraded
	CarLimit      int    `json:"carLimit"`      //limit of cars a player can store in that garage
}
type GarageCarList struct {
	GarageId string `json:"garageId"`
	CarId    string `json:"carId"`
}
