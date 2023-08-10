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
		// Validate Latitude: must be between -90 and 90 degrees
		validation.Field(&a.Latitude, validation.Min(-90), validation.Max(90)),
		// Validate Longitude: must be between -180 and 180 degrees
		validation.Field(&a.Longitude, validation.Min(-180), validation.Max(180)),
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
		// Validate Latitude: must be between -90 and 90 degrees
		validation.Field(&a.Latitude, validation.Required, validation.Min(-90.0), validation.Max(90.0)),
		// Validate Longitude: must be between -180 and 180 degrees
		validation.Field(&a.Longitude, validation.Required, validation.Min(-180.0), validation.Max(180.0)),
	)
}
