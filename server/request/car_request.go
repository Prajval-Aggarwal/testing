package request

import validation "github.com/go-ozzo/ozzo-validation"

type BuyCarRequest struct {
	CarId string `json:"carId"`
}

func (a BuyCarRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CarId, validation.Required),
	)
}
