package model

import "time"

type Arena struct {
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
