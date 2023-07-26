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
	RaceType string    `json:"raceType"`
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
		validation.Field(&a.RaceType, validation.Required),
	)
}

func (a GetArenaReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ArenaId, validation.Required),
	)
}
