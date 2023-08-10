package model

type PlayerLevel struct {
	Level      uint64 `json:"level"`
	XPRequired uint64 `json:"xpRequired"`
	Coins      uint64 `json:"coins"`
}
