package request

import validation "github.com/go-ozzo/ozzo-validation"

type GarageRequest struct {
	GarageId string `json:"carId"`
}

// add car to garage request model
type AddCarRequest struct {
	GarageId string `json:"garageId"`
	CarId    string `json:"carId"`
}

type AddGarageRequest struct {
	GarageName    string  `json:"garageName,omitempty"`
	GarageType    int64   `json:"garageType,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	Level         int64   `json:"level,omitempty"`         //level reuired to unlock the garage
	CoinsRequired int64   `json:"coinsRequired,omitempty"` //coins required to unlock the garage
}

type DeletGarageReq struct {
	GarageId string `json:"garageId"`
}

type UpdateGarageReq struct {
	GarageId      string  `json:"garageId"`
	GarageName    string  `json:"garageName,omitempty"`
	GarageType    int64   `json:"garageType"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	Level         int64   `json:"level,omitempty"`         //level reuired to unlock the garage
	CoinsRequired int64   `json:"coinsRequired,omitempty"` //coins required to unlock the garage
}

func (a UpdateGarageReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GarageId, validation.Required),
	)
}
func (a DeletGarageReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GarageId, validation.Required),
	)
}
func (a AddGarageRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GarageName, validation.Required),
		validation.Field(&a.GarageType, validation.Required),
		validation.Field(&a.Latitude, validation.Required),
		validation.Field(&a.Longitude, validation.Required),
		validation.Field(&a.Level, validation.Required),
		validation.Field(&a.CoinsRequired, validation.Required),
	)
}

func (a GarageRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GarageId, validation.Required),
	)
}

func (a AddCarRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GarageId, validation.Required),
		validation.Field(&a.CarId, validation.Required),
	)
}
