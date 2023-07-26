package model

type Rewards struct {
	RewardId    string `json:"id"`
	Coins       int64  `json:"coins"`
	Cash        int64  `json:"cash"`
	RepairParts int64  `json:"repairParts"`
	Status      string `json:"status"`
	XPGained    int64  `json:"xpGained"`
}
