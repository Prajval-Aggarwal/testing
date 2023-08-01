package model

type Car struct {
	CarId      string  `json:"carId"  gorm:"default:uuid_generate_v4();primaryKey"`
	CarName    string  `json:"carName,omitempty"`
	CurrType   string  `json:"currType,omitempty" `
	CurrAmount float64 `json:"cost,omitempty"`
	Class      string  `json:"class,omitempty"`
	Locked     bool    `json:"locked,omitempty"`
}

type CarStats struct {
	CarId       string  `json:"carId,omitempty"`
	Power       uint64  `json:"power,omitempty"`
	Grip        uint64  `json:"grip,omitempty"`
	Weight      uint64  `json:"weight,omitempty"`
	ShiftTime   float64 `json:"shiftTime,omitempty"`
	OVR         float64 `json:"or,omitempty"` //overall rating of the car
	Durability  uint64  `json:"durability,omitempty"`
	NitrousTime uint64  `json:"nitrousTime,omitempty"` //increased when nitrous is upgraded
}
