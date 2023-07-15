package model

type Arena struct {
	ArenaId    string  `json:"arenaId"`
	ArenaName  string  `json:"arenaName"`
	ArenaLevel int     `json:"arenaLevel"`
	Perks      int     `json:"perks"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}

// minimum requirements for an arena
type ArenaReq struct {
	ArenaLevel  int64 `json:"arenaLevel"`
	PlayerLevel int64 `json:"playerLevel"`
	MinCarReq   int64 `json:"minCarReq"`
}
