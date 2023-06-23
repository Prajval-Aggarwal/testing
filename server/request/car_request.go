package request

import validation "github.com/go-ozzo/ozzo-validation"

type CarRequest struct {
	CarId string `json:"carId"`
}

func (a CarRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CarId, validation.Required),
	)
}

type CarUpgradesRequest struct {
	CarId       string `json:"carId"`
	PaymentMode string `json:"paymentMode"`
}

func (a CarUpgradesRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CarId, validation.Required),
		validation.Field(&a.PaymentMode, validation.Required),
	)
}
