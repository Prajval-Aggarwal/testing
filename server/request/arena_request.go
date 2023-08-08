package request

import validation "github.com/go-ozzo/ozzo-validation"

type AddArenaRequest struct {
	ArenaName  string  `json:"arenaName,omitempty"`
	ArenaLevel int64   `json:"arenaLevel,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

type DeletArenaReq struct {
	ArenaId string `json:"arenaId"`
}

type UpdateArenaReq struct {
	ArenaId    string  `json:"arenaId"`
	ArenaName  string  `json:"arenaName,omitempty"`
	ArenaLevel int64   `json:"arenaLevel,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

func (a UpdateArenaReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ArenaId, validation.Required),
		validation.Field(&a.ArenaLevel),
	)
}

func (a DeletArenaReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ArenaId, validation.Required),
	)
}

func (a AddArenaRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ArenaName, validation.Required),
		validation.Field(&a.Latitude, validation.Required),
		validation.Field(&a.Longitude, validation.Required),
		validation.Field(&a.ArenaLevel, validation.Required),
	)
}
