package redis

type User struct {
	UserID string `json:"user_id" redis:"user_id"`
	Email  string `json:"email" redis:"email"`
}
