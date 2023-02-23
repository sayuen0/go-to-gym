package models

type Session struct {
	SessionID string `json:"sessionId" redis:"session_id"`
	UserID    string `json:"userId" redis:"user_id"`
}
