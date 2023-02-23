package redis

type User struct {
	UserID string `json:"userId" redis:"user_id"`
	Email  string `json:"email" redis:"email"`
}
