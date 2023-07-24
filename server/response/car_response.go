package response

import "main/server/model"

type CarResponse struct {
	CarName       string                       `json:"carName"`
	Car           model.Car                    `json:"info"`
	CarStats      model.CarStats               `json:"stats"`
	Customization []model.DefaultCustomization `json:"customization"`
}

type CarStatResponse struct {
	Power      int64   `json:"power,omitempty"`
	Grip       int64   `json:"grip,omitempty"`
	Weight     int64   `json:"weight,omitempty"`
	ShiftTime  float64 `json:"shiftTime,omitempty"`
	OVR        float64 `json:"or,omitempty"` //overall rating of the car
	Durability int64   `json:"durability,omitempty"`
	TempPower  int64   `json:"tempPower,omitempty"`
}

type CarUpgradeResponse struct {
	UpgradedPart string  `json:"upgradedPart,omitempty"`
	Power        int64   `json:"power,omitempty"`
	Grip         int64   `json:"grip,omitempty"`
	Shift_Time   float64 `json:"shiftTime,omitempty"`
	OVR          float64 `json:"ovr,omitempty"`
	Coins        int64   `json:"coins,omitempty"`
	Upgradable   bool    `json:"upgradable,omitempty"`
	NextLevel    int64   `json:"nextLevel,omitempty"`
	NextCost     int64   `json:"nextCost,omitempty"`
}
