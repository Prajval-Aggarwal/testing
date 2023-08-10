package response

type RewardData struct {
	Coins       uint64 `json:"coins,omitempty"`
	Cash        uint64 `json:"cash,omitempty"`
	RepairParts uint64 `json:"repairParts,omitempty"`
	XPGained    uint64 `json:"xpGained,omitempty"`
	Status      string `json:"status,omitempty"`
	Level       uint64 `json:"level,omitempty"`
	XPRequired  uint64 `json:"xpRequired,omitempty"`
}

type RewardResponse struct {
	RewardName string     `json:"rewardName"`
	RewardData RewardData `json:"rewardData"`
}
