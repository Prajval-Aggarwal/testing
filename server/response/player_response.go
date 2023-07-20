package response

type PlayerResposne struct {
	PlayerId         string `json:"playerId"`
	PlayerName       string `json:"playerName" gorm:"unique"`
	Level            int    `json:"level"`
	XP               int64  `json:"xp"`
	Role             string `json:"role"`
	Email            string `json:"email"`
	Coins            int64  `json:"coins"`
	Cash             int64  `json:"cash"`
	CarsOwned        int64  `json:"carsOwned"`
	GaragesOwned     int64  `json:"garagesOwned"`
	DistanceTraveled int64  `json:"distanceTraveled"`
	ShdWon           int64  `json:"showDownWon"`
	ShdWinRatio      int64  `json:"showDownWinRatio"`
	TdWon            int64  `json:"takeDownWon"`
	TdWinRatio       int64  `json:"takeDownWinRatio"`
}
