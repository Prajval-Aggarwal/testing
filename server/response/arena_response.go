package response

// resposne format for get arena by id
type ArenaIdRes struct {
	ArenaId     string `json:"arenaId"`
	ArenaName   string `json:"arenaName"`
	Level       string `json:"arenaLevel"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	PlayerLevel int64  `json:"playerLevelReq"`
	MinCarReq   int64  `json:"minCarReq"`
}

// resposne for getting list of arenas
type ArenaResp struct {
	ArenaId   string `json:"arenaId"`
	ArenaName string `json:"arenaName"`
	Level     string `json:"arenaLevel"`
}
