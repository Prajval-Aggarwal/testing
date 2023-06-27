package model

type Player struct {
	PlayerId   string `json:"playerId"`
	PlayerName string `json:"playerName" gorm:"unique"`
	Level      int    `json:"level"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Coins      int64  `json:"coins"`
	Cash       int64  `json:"cash"`
	DeviceId   string `json:"deviceId"`
	OS         int64  `json:"os"` // o for android 1 for ios
}

type OwnedCars struct {
	PlayerId string `json:"playerId"`
	CarId    string `json:"carId"`
	Selected bool   `json:"selected"`
	Level    int    `json:"level"`
}

type OwnedGarage struct {
	PlayerId    string `json:"playerId"`
	GarageId    string `json:"garageId"`
	GarageLevel int    `json:"garageLevel"`
	CarLimit    int    `json:"carLimit"`
}
type PlayerCarsStats struct {
	PlayerId    string  `json:"playerId"`
	CarId       string  `json:"carId"`
	Power       int64   `json:"power"`
	Grip        int64   `json:"grip"`
	ShiftTime   float64 `json:"shiftTime"`
	Weight      int64   `json:"weight"`
	OVR         float64 `json:"or"` //overall rating of the car
	Durability  int64   `json:"Durability"`
	NitrousTime float64 `json:"nitrousTime"` //increased when nitrous is upgraded

}

type PlayerCarUpgrades struct {
	PlayerId     string `json:"playerId"`
	CarId        string `json:"carId"`
	Engine       int    `json:"engine"`       // Affects Power
	Turbo        int    `json:"turbo"`        // Affects Power
	Intake       int    `json:"intake"`       // Affects Power
	Nitrous      int    `json:"nitrous"`      // Affect Nitrous time
	Body         int    `json:"body"`         //Affects Grip and Weight
	Tires        int    `json:"tires"`        //Affects Grip
	Transmission int    `json:"transmission"` //Affects Shift-Time
}

type PlayerCarCustomization struct {
	PlayerId      string `json:"playerId"`
	CarId         string `json:"carId"`
	Part          string `json:"part"`
	ColorCategory string `json:"colorCategory"`
	ColorType     string `json:"colorType"`
	ColorCode     string `json:"colorCode"`
	ColorName     string `json:"colorName"`
	Value         string `json:"value"`
}
