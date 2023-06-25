package model

type Car struct {
	CarId      string  `json:"carId"  gorm:"default:uuid_generate_v4();primaryKey"`
	CarName    string  `json:"carName"`
	CurrLevel  int     `json:"currLevel"` //current level of the car
	Level      int64   `json:"level"`     // level required to unlock the car
	CurrType   string  `json:"currType"`
	CurrAmount float64 `json:"currAmount"`
	MaxLevel   int64   `json:"maxLevel"`
	Class      string  `json:"class"`
	Locked     bool    `json:"locked"`
}

type CarStats struct {
	CarId       string  `json:"carId"`
	Power       int64   `json:"power"`
	Grip        int64   `json:"grip"`
	Weight      int64   `json:"weight"`
	ShiftTime   float64 `json:"shiftTime"`
	OR          float64 `json:"or"` //overall rating of the car
	Durability  int64   `json:"Durability"`
	NitrousTime int     `json:"nitrousTime"` //increased when nitrous is upgraded
}
