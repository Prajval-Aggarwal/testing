package model

type Engine struct {
	CarClass  string  `json:"carClass,omitempty"`
	Level     int     `json:"level,omitempty"`
	Power     int64   `json:"power,omitempty"`
	OVR       float64 `json:"or,omitempty"`
	CashPrice int64   `json:"cashPrice,omitempty"`
	CoinPrice int64   `json:"coinPrice,omitempty"`
}
type Turbo struct {
	CarClass  string  `json:"carClass,omitempty"`
	Level     int     `json:"level,omitempty"`
	Power     int64   `json:"power,omitempty"`
	OVR       float64 `json:"or,omitempty"`
	CashPrice int64   `json:"cashPrice,omitempty"`
	CoinPrice int64   `json:"coinPrice,omitempty"`
}
type Intake struct {
	CarClass  string  `json:"carClass,omitempty"`
	Level     int     `json:"level,omitempty"`
	Power     int64   `json:"power,omitempty"`
	OVR       float64 `json:"or,omitempty"`
	CashPrice int64   `json:"cashPrice,omitempty"`
	CoinPrice int64   `json:"coinPrice,omitempty"`
}
type Nitrous struct {
	CarClass    string  `json:"carClass,omitempty"`
	Level       int     `json:"level,omitempty"`
	NitrousTime int64   `json:"nitrousTime,omitempty"` //increased when nitrous is upgraded
	OVR         float64 `json:"or,omitempty"`
	CashPrice   int64   `json:"cashPrice,omitempty"`
	CoinPrice   int64   `json:"coinPrice,omitempty"`
}
type Tires struct {
	CarClass  string  `json:"carClass,omitempty"`
	Level     int     `json:"level,omitempty"`
	Grip      int64   `json:"grip,omitempty"`
	OVR       float64 `json:"or,omitempty"`
	CashPrice int64   `json:"cashPrice,omitempty"`
	CoinPrice int64   `json:"coinPrice,omitempty"`
}
type Body struct {
	CarClass  string  `json:"carClass,omitempty"`
	Level     int     `json:"level,omitempty"`
	Grip      int64   `json:"grip,omitempty"`
	Weight    int64   `json:"weight,omitempty"`
	OVR       float64 `json:"or,omitempty"`
	CashPrice int64   `json:"cashPrice,omitempty"`
	CoinPrice int64   `json:"coinPrice,omitempty"`
}
type Transmission struct {
	CarClass  string  `json:"carClass,omitempty"`
	Level     int     `json:"level,omitempty"`
	ShiftTime float64 `json:"shiftTime,omitempty"`
	OVR       float64 `json:"or,omitempty"`
	CashPrice int64   `json:"cashPrice,omitempty"`
	CoinPrice int64   `json:"coinPrice,omitempty"`
}

/*○ Engine
■ Affects Power
■ Upgrade increases Power
○ Turbo
■ Affects Power
■ Upgrade increases Power
○ Intake
■ Affects Power
■ Upgrade increases Power
○ Nitrous
■ Increases Power for a certain period of time
■ Upgrade effects both:
● Amount of Power increased, and
● Period of time for which Nitrous would be enabled for
○ Period of time would be in seconds
○ Body
■ Affects Grip and Weight
■ Upgrade increases Grip and decreases Weight
○ Tires
■ Affects Grip
■ Upgrade Increases Grip
○ Transmission
■ Affects Shift-Time */
