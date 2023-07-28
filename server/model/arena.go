package model

import "time"

type Arena struct {
	ArenaId    string  `json:"arenaId"`
	ArenaName  string  `json:"arenaName"`
	ArenaLevel string  `json:"arenaLevel"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}

// minimum requirements for an arena
type ArenaRaceRecord struct {
	RaceId   string    `json:"raceId" gorm:"default:uuid_generate_v4();primaryKey"`
	PlayerId string    `json:"playerId"`
	ArenaId  string    `json:"arenaId"`
	Time     time.Time `json:"time"`
	Result   string    `json:"result"`
}

type ArenaSeries struct {
	ArenaId   string `json:"ArenaId"`
	PlayerId  string `json:"playerId"`
	WinStreak int64  `json:"winStreak"`
}

type CarSlots struct {
	PlayerId string `json:"PlayerId"`
	ArenaId  string `json:"ArenaId"`
	CardId   string `json:"CardId"`
}
