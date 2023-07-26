package model

type RaceTypes struct {
	RaceId     string  `json:"raceId" gorm:"default:uuid_generate_v4()"`
	RaceName   string  `json:"raceName"`
	RaceLength float64 `json:"raceLength"`
	RaceSeries int64   `json:"raceSeries"`
	RaceLevel  string  `json:"raceLevel"`
}
