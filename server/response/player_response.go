package response

type PlayerResposne struct {
	PlayerId         string `json:"playerId"`
	PlayerName       string `json:"playerName" gorm:"unique"`
	Level            int    `json:"level"`
	XP               uint64  `json:"xp"`
	Role             string `json:"role"`
	Email            string `json:"email"`
	Coins            uint64  `json:"coins"`
	Cash             uint64  `json:"cash"`
	RepairParts      uint64  `json:"repairRewards"`
	CarsOwned        uint64  `json:"carsOwned"`
	GaragesOwned     uint64  `json:"garagesOwned"`
	DistanceTraveled uint64  `json:"distanceTraveled"`
	ShdWon           uint64  `json:"showDownWon"`
	ShdWinRatio      uint64  `json:"showDownWinRatio"`
	TdWon            uint64  `json:"takeDownWon"`
	TdWinRatio       uint64  `json:"takeDownWinRatio"`
}
