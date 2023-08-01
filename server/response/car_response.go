package response

import "main/server/model"

type CarResponse struct {
	CarName       string                       `json:"carName"`
	Car           model.Car                    `json:"info"`
	CarStats      model.CarStats               `json:"stats"`
	Customization []model.DefaultCustomization `json:"customization"`
}

type CarStatResponse struct {
	Power      uint64  `json:"power,omitempty"`
	Grip       uint64  `json:"grip,omitempty"`
	Weight     uint64  `json:"weight,omitempty"`
	ShiftTime  float64 `json:"shiftTime,omitempty"`
	OVR        float64 `json:"or,omitempty"` //overall rating of the car
	Durability uint64  `json:"durability,omitempty"`
	TempPower  uint64  `json:"tempPower,omitempty"`
}

type CarUpgradeResponse struct {
	UpgradedPart string  `json:"upgradedPart,omitempty"`
	Power        uint64  `json:"power,omitempty"`
	Grip         uint64  `json:"grip,omitempty"`
	Shift_Time   float64 `json:"shiftTime,omitempty"`
	OVR          float64 `json:"ovr,omitempty"`
	Coins        uint64  `json:"coins,omitempty"`
	Upgradable   bool    `json:"upgradable,omitempty"`
	NextLevel    uint64  `json:"nextLevel,omitempty"`
	NextCost     uint64  `json:"nextCost,omitempty"`
}
