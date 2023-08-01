package model

import "time"

type Player struct {
	PlayerId    string `json:"playerId,omitempty"`
	PlayerName  string `json:"playerName" gorm:"unique,omitempty"`
	Level       uint64 `json:"level,omitempty"`
	Role        string `json:"role,omitempty"`
	XP          int64  `json:"xp,omitempty"`
	Email       string `json:"email,omitempty"`
	Coins       uint64 `json:"coins,omitempty"`
	Cash        uint64 `json:"cash,omitempty"`
	RepairParts uint64 `json:"repairParts,omitempty"`
	DeviceId    string `json:"deviceId,omitempty"`
	OS          uint64 `json:"os,omitempty"` // o for android 1 for ios
}

type OwnedCars struct {
	PlayerId string `json:"playerId,omitempty"`
	CarId    string `json:"carId,omitempty"`
	Selected bool   `json:"selected,omitempty"`
}

type OwnedGarage struct {
	PlayerId    string `json:"playerId,omitempty"`
	GarageId    string `json:"garageId,omitempty"`
	GarageLevel uint64 `json:"garageLevel,omitempty"`
	CarLimit    uint64 `json:"carLimit,omitempty"`
}

type PlayerRaceHistory struct {
	PlayerId         string  `json:"playerId,omitempty"`
	DistanceTraveled float64 `json:"distanceTraveled"`
	ShdWon           uint64  `json:"showDownWon"`
	TotalShdPlayed   uint64  `json:"totalShdPlayed"`
	TdWon            uint64  `json:"takeDownWon"`
	TotalTdPlayed    uint64  `json:"totalTdPlayed"`
}

type OwnedBattleArenas struct {
	PlayerId  string    `json:"playerId,omitempty"`
	ArenaId   string    `json:"arenaId,omitempty"`
	WinTime   time.Time `json:"winTime,omitempty"`
	CreatedAt time.Time
	CarId     string `json:"carId,omitempty"`
	Status    string `json:"status,omitempty"`
}
type PlayerCarsStats struct {
	PlayerId    string  `json:"playerId,omitempty"`
	CarId       string  `json:"carId,omitempty"`
	Power       uint64  `json:"power,omitempty"`
	Grip        uint64  `json:"grip,omitempty"`
	ShiftTime   float64 `json:"shiftTime,omitempty"`
	Weight      uint64  `json:"weight,omitempty"`
	OVR         float64 `json:"or,omitempty"` //overall rating of the car
	Durability  uint64  `json:"Durability,omitempty"`
	NitrousTime float64 `json:"nitrousTime,omitempty"` //increased when nitrous is upgraded

}

type PlayerCarUpgrades struct {
	PlayerId     string `json:"playerId,omitempty"`
	CarId        string `json:"carId,omitempty"`
	Engine       uint64 `json:"engine,omitempty"`       // Affects Power
	Turbo        uint64 `json:"turbo,omitempty"`        // Affects Power
	Intake       uint64 `json:"intake,omitempty"`       // Affects Power
	Nitrous      uint64 `json:"nitrous,omitempty"`      // Affect Nitrous time
	Body         uint64 `json:"body,omitempty"`         //Affects Grip and Weight
	Tires        uint64 `json:"tires,omitempty"`        //Affects Grip
	Transmission uint64 `json:"transmission,omitempty"` //Affects Shift-Time
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
