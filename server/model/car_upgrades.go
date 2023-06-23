package model

type Engine struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Power     int64  `json:"power"`
	CashPrice int    `json:"cashPrice"`
	CoinPrice int    `json:"coinPrice"`
}
type Turbo struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Price     int    `json:"price"`
	Power     int64  `json:"power"`
	CashPrice int    `json:"cashPrice"`
	CoinPrice int    `json:"coinPrice"`
}
type Intake struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Price     int    `json:"price"`
	Power     int64  `json:"power"`
	CashPrice int    `json:"cashPrice"`
	CoinPrice int    `json:"coinPrice"`
}
type Nitrous struct {
	CarClass    string `json:"carClass"`
	Level       int    `json:"level"`
	Price       int    `json:"price"`
	NitrousTime int    `json:"nitrousTime"` //increased when nitrous is upgraded
	CashPrice   int    `json:"cashPrice"`
	CoinPrice   int    `json:"coinPrice"`
}
type Body struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Price     int    `json:"price"`
	Grip      int64  `json:"grip"`
	Weight    int64  `json:"weight"`
	CashPrice int    `json:"cashPrice"`
	CoinPrice int    `json:"coinPrice"`
}
type Tires struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Price     int    `json:"price"`
	Grip      int64  `json:"grip"`
	CashPrice int    `json:"cashPrice"`
	CoinPrice int    `json:"coinPrice"`
}
type Transmission struct {
	CarClass  string  `json:"carClass"`
	Level     int     `json:"level"`
	Price     int     `json:"price"`
	ShiftTime float64 `json:"shiftTime"`
	CashPrice int     `json:"cashPrice"`
	CoinPrice int     `json:"coinPrice"`
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
