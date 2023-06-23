package request

import validation "github.com/go-ozzo/ozzo-validation"

type GarageRequest struct {
	GarageId string `json:"carId"`
}

//add car to garage request model
type AddCarRequest struct {
	GarageId string `json:"garageId"`
	CarId    string `json:"carId"`
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
