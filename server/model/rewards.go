package model

type Rewards struct {
	RewardId    string `json:"id"`
	Coins       uint64 `json:"coins"`
	Cash        uint64 `json:"cash"`
	RepairParts uint64 `json:"repairParts"`
	Status      string `json:"status"`
	XPGained    int64  `json:"xpGained"`
}
