package request

import (
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
	ArenaId  string `json:"arenaId"`
	PlayerId string `json:"playerId"`
	WinTime  string `json:"winTime2"`
	RaceId   string `json:"raceId"`
}

type AddArenaRequest struct {
	ArenaName  string  `json:"arenaName,omitempty"`
	ArenaLevel string  `json:"arenaLevel,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

type DeletArenaReq struct {
	ArenaId string `json:"arenaId"`
}

type UpdateArenaReq struct {
	ArenaId    string  `json:"arenaId"`
	ArenaName  string  `json:"arenaName,omitempty"`
	ArenaLevel string  `json:"arenaLevel,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

func (a UpdateArenaReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ArenaId, validation.Required),
		validation.Field(&a.ArenaLevel, validation.In("easy", "medium", "hard")),
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
		validation.Field(&a.ArenaLevel, validation.Required, validation.In("easy", "medium", "hard")),
	)
}

// validation on structs
func (a ChallengeReq) Validate() error {
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
