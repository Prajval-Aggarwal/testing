package model

type Engine struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Power     int64  `json:"power"`
	OR        int64  `json:"overallRating"`
	CashPrice int64  `json:"cashPrice"`
	CoinPrice int64  `json:"coinPrice"`
}
type Turbo struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Power     int64  `json:"power"`
	OR        int64  `json:"overallRating"`
	CashPrice int64  `json:"cashPrice"`
	CoinPrice int64  `json:"coinPrice"`
}
type Intake struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Power     int64  `json:"power"`
	OR        int64  `json:"overallRating"`
	CashPrice int64  `json:"cashPrice"`
	CoinPrice int64  `json:"coinPrice"`
}
type Nitrous struct {
	CarClass    string `json:"carClass"`
	Level       int    `json:"level"`
	NitrousTime int64  `json:"nitrousTime"` //increased when nitrous is upgraded
	CashPrice   int64  `json:"cashPrice"`
	CoinPrice   int64  `json:"coinPrice"`
}
type Tires struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Grip      int64  `json:"grip"`
	OR        int64  `json:"overallRating"`
	CashPrice int64  `json:"cashPrice"`
	CoinPrice int64  `json:"coinPrice"`
}
type Body struct {
	CarClass  string `json:"carClass"`
	Level     int    `json:"level"`
	Grip      int64  `json:"grip"`
	Weight    int64  `json:"weight"`
	OR        int64  `json:"overallRating"`
	CashPrice int64  `json:"cashPrice"`
	CoinPrice int64  `json:"coinPrice"`
}
type Transmission struct {
	CarClass  string  `json:"carClass"`
	Level     int     `json:"level"`
	ShiftTime float64 `json:"shiftTime"`
	OR        int64   `json:"overallRating"`
	CashPrice int64   `json:"cashPrice"`
	CoinPrice int64   `json:"coinPrice"`
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
