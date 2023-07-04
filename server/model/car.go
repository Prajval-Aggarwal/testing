package model

type Car struct {
	CarId      string  `json:"carId"  gorm:"default:uuid_generate_v4();primaryKey"`
	CarName    string  `json:"carName,omitempty"`
	Level      int64   `json:"level,omitempty"` // level required to unlock the car
	CurrType   string  `json:"currType,omitempty" `
	CurrAmount float64 `json:"currAmount,omitempty"`
	MaxLevel   int64   `json:"maxLevel,omitempty"`
	Class      string  `json:"class,omitempty"`
	Locked     bool    `json:"locked,omitempty"`
}

type CarStats struct {
	CarId       string  `json:"carId,omitempty"`
	Power       int64   `json:"power,omitempty"`
	Grip        int64   `json:"grip,omitempty"`
	Weight      int64   `json:"weight,omitempty"`
	ShiftTime   float64 `json:"shiftTime,omitempty"`
	OVR         float64 `json:"or,omitempty"` //overall rating of the car
	Durability  int64   `json:"Durability,omitempty"`
	NitrousTime int     `json:"nitrousTime,omitempty"` //increased when nitrous is upgraded
}
