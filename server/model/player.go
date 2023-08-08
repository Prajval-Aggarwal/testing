package model

type Player struct {
	PlayerId   string `json:"playerId,omitempty"`
	PlayerName string `json:"playerName" gorm:"unique"`
	Level      int    `json:"level,omitempty"`
	Role       string `json:"role,omitempty"`
	Email      string `json:"email,omitempty"`
	Coins      int64  `json:"coins,omitempty"`
	Cash       int64  `json:"cash,omitempty"`
	DeviceId   string `json:"deviceId,omitempty"`
	OS         int64  `json:"os,omitempty"` // o for android 1 for ios
}

type OwnedCars struct {
	PlayerId   string  `json:"playerId,omitempty"`
	CarId      string  `json:"carId,omitempty"`
	Selected   bool    `json:"selected,omitempty"`
	Level      int     `json:"level,omitempty"`
	RepairCost float64 `json:"repairCost,omitempty"` //will always b in coins no cash
}

type OwnedGarage struct {
	PlayerId    string `json:"playerId,omitempty"`
	GarageId    string `json:"garageId,omitempty"`
	GarageLevel int    `json:"garageLevel,omitempty"`
	CarLimit    int    `json:"carLimit,omitempty"`
}

type PlayerRaceHistory struct {
	PlayerId         string `json:"playerId,omitempty"`
	DistanceTraveled int64  `json:"distanceTraveled"`
	ShdWon           int64  `json:"showDownWon"`
	TotalShdPlayed   int64  `json:"totalShdPlayed"`
	TdWon            int64  `json:"takeDownWon"`
	TotalTdPlayed    int64  `json:"totalTdPlayed"`
}

type OwnedBattleArenas struct {
	PlayerId string `json:"playerId,omitempty"`
	ArenaId  string `json:"arenaId,omitempty"`
}
type PlayerCarsStats struct {
	PlayerId    string  `json:"playerId,omitempty"`
	CarId       string  `json:"carId,omitempty"`
	Power       int64   `json:"power,omitempty"`
	Grip        int64   `json:"grip,omitempty"`
	ShiftTime   float64 `json:"shiftTime,omitempty"`
	Weight      int64   `json:"weight,omitempty"`
	OVR         float64 `json:"or,omitempty"` //overall rating of the car
	Durability  int64   `json:"Durability,omitempty"`
	NitrousTime float64 `json:"nitrousTime,omitempty"` //increased when nitrous is upgraded

}

type PlayerCarUpgrades struct {
	PlayerId     string `json:"playerId,omitempty"`
	CarId        string `json:"carId,omitempty"`
	Engine       int    `json:"engine,omitempty"`       // Affects Power
	Turbo        int    `json:"turbo,omitempty"`        // Affects Power
	Intake       int    `json:"intake,omitempty"`       // Affects Power
	Nitrous      int    `json:"nitrous,omitempty"`      // Affect Nitrous time
	Body         int    `json:"body,omitempty"`         //Affects Grip and Weight
	Tires        int    `json:"tires,omitempty"`        //Affects Grip
	Transmission int    `json:"transmission,omitempty"` //Affects Shift-Time
}

type PlayerCarCustomization struct {
	PlayerId      string `json:"playerId,omitempty"`
	CarId         string `json:"carId,omitempty"`
	Part          string `json:"part,omitempty"`
	ColorCategory string `json:"colorCategory,omitempty"`
	ColorType     string `json:"colorType,omitempty"`
	ColorCode     string `json:"colorCode,omitempty"`
	ColorName     string `json:"colorName,omitempty"`
	Value         string `json:"value,omitempty"`
}
