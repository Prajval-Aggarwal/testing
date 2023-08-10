package request

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ChallengeReq struct {
	ArenaId string `json:"arenaId"`
}
type ReplaceReq struct {
	ArenaId string
	CarId   string
}

type AddCarArenaRequest struct {
	ArenaId string `json:"arenaId"`
	CarId   string `json:"carId"`
}
type GetArenaReq struct {
	ArenaId string `json:"arenaId"`
}
type EndChallengeReq struct {
	ArenaId  string    `json:"arenaId"`
	PlayerId string    `json:"playerId"`
	WinTime  time.Time `json:"winTime2"`
	RaceId   string    `json:"raceId"`
}

type AddArenaRequest struct {
	ArenaName  string  `json:"arenaName,omitempty"`
	ArenaLevel uint64  `json:"arenaLevel,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

type DeletArenaReq struct {
	ArenaId string `json:"arenaId"`
}

type UpdateArenaReq struct {
	ArenaId    string  `json:"arenaId"`
	ArenaName  string  `json:"arenaName,omitempty"`
	ArenaLevel uint64  `json:"arenaLevel,omitempty"`
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

func (a ReplaceReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ArenaId, validation.Required),
		validation.Field(&a.CarId, validation.Required),
	)
}

func (a AddCarArenaRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ArenaId, validation.Required),
		validation.Field(&a.CarId, validation.Required),
	)
}

func (a EndChallengeReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.PlayerId, validation.Required),
		validation.Field(&a.WinTime, validation.Required),
		validation.Field(&a.RaceId, validation.Required),
	)
}

func (a GetArenaReq) Validate() error {
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
