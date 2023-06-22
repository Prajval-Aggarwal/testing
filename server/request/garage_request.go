package request

import validation "github.com/go-ozzo/ozzo-validation"

type GarageRequest struct {
	GarageId string `json:"carId"`
}

func (a GarageRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.GarageId, validation.Required),
	)
}
