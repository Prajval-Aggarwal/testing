package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type GuestLoginRequest struct {
	PlayerName string `json:"playerName"`
	DeviceId   string `json:"deviceId"`
	OS         int64  `json:"os"`
	Token      string `json:"token"`
}
type LoginRequest struct {
	Email string `json:"email"`
}

type UpdateEmailRequest struct {
	Email string `json:"email"`
}

// Validation
func (a GuestLoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.PlayerName, validation.Required),
		validation.Field(&a.DeviceId, validation.Required),
		validation.Field(&a.OS, validation.Required),
	)
}
func (a LoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}

func (a UpdateEmailRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}
