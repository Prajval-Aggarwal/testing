package model

import "time"

type Arena struct {
	ArenaId    string  `json:"arenaId"`
	ArenaName  string  `json:"arenaName"`
	ArenaLevel uint64  `json:"arenaLevel"`
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
	WinStreak uint64 `json:"winStreak"`
}

type CarSlots struct {
	PlayerId   string    `json:"PlayerId"`
	CardId     string    `json:"CardId"`
	ArenaId    string    `json:"arenaId" gorm:"default:uuid_generate_v4();primaryKey,omitempty"`
	ArenaName  string    `json:"arenaName"`
	ArenaLevel int64     `json:"arenaLevel"`
	Longitude  float64   `json:"longitude"`
	Latitude   float64   `json:"latitude"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

type ArenaLevels struct {
	TypeName string `json:"label,omitempty" gorm:"unique"`
	TypeId   int    `json:"value"`
}
