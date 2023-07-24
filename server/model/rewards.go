package model

type WinRewards struct {
	ChallengeType string  `json:"challengeType"`
	RaceLength    float64 `json:"raceLength"`
	Difficulty    string  `json:"difficulty"`
	Coins         int64   `json:"coins"`
	Cash          int64   `json:"cash"`
	RepairParts   int64   `json:"repairParts"`
	XPGained      int64   `json:"xpGained"`
}

type LostRewards struct {
	ChallengeType string  `json:"challengeType"`
	RaceLength    float64 `json:"raceLength"`
	Difficulty    string  `json:"difficulty"`
	Coins         int64   `json:"coins"`
	Cash          int64   `json:"cash"`
	RepairParts   int64   `json:"repairParts"`
	XPGained      int64   `json:"xpGained"`
}
