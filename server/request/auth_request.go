package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AdminLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GuestLoginRequest struct {
	PlayerName string `json:"playerName"`
	DeviceId   string `json:"deviceId"`
	OS         int    `json:"os"`
	Token      string `json:"token"`
}
type LoginRequest struct {
	Email string `json:"email"`
}

type UpdateEmailRequest struct {
	Email string `json:"email"`
}
type ForgotPassRequest struct {
	Email string `json:"email" `
}

type UpdatePasswordRequest struct {
	Password string `json:"password" `
}

// Validation
func (a UpdatePasswordRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Password, validation.Required),
	)
}
func (a ForgotPassRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}

func (a AdminLoginReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
}
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
