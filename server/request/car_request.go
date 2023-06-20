package request

import validation "github.com/go-ozzo/ozzo-validation"

type AddCarRequest struct {
	CarName    string  `json:"carName"`
	Level      int64   `json:"level"`
	CurrType   string  `json:"currType"`
	CurrAmount float64 `json:"currAmount"`
	MaxLevel   int64   `json:"maxLevel"`
	Class      string  `json:"class"`
	Status     string  `json:"status"`
	Power      int64   `json:"power"`
	Grip       int64   `json:"grip"`
	ShiftTime  float64 `json:"shiftTime"`
	Weight     int64   `json:"weight"`
	OR         float64 `json:"or"` //overall rating of the car
	Durability int64   `json:"durability"`
}

func (a AddCarRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CarName, validation.Required),
		validation.Field(&a.Level, validation.Required),
		validation.Field(&a.CurrType, validation.Required),
		validation.Field(&a.CurrAmount, validation.Required),
		validation.Field(&a.MaxLevel, validation.Required),
		validation.Field(&a.Class, validation.Required),
		validation.Field(&a.Status, validation.Required),
		validation.Field(&a.Power, validation.Required),
		validation.Field(&a.Grip, validation.Required),
		validation.Field(&a.ShiftTime, validation.Required),
		validation.Field(&a.Weight, validation.Required),
		validation.Field(&a.OR, validation.Required),
		validation.Field(&a.Durability, validation.Required),
	)
}
