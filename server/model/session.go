package model

type Session struct {
	SessionId string `json:"sessionId" gorm:"default:uuid_generate_v4()"`
	Token     string `json:"token"`
}

type ResetSession struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
